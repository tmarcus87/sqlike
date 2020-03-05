package statement

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/tmarcus87/sqlike/logger"
	"reflect"
	"strings"
)

var (
	ErrorNoSteps          = errors.New("no steps")
	ErrorMustBeASlice     = errors.New("must be a slice")
	ErrorMustBeAPtr       = errors.New("must be a pointer")
	ErrorMustBeANonNilPtr = errors.New("must be a non-nil pointer")
	ErrorMustBeAStructPtr = errors.New("must be a pointer to struct")
)

type StatementAcceptor interface {
	// Parent 親のStatementAcceptorを返します。親がない場合はnilです
	Parent() StatementAcceptor

	// Accept Statementを受け取ってクエリを組み立てます
	Accept(stmt *StatementImpl) error
}

func NewStatementBuilder(s StatementAcceptor) Statement {
	return &StatementImpl{sa: s}
}

type Statement interface {
	StatementAndBindings() (string, []interface{}, error)
	FetchMap() ([]map[string]string, error)
	FetchInto(p interface{}) error
	FetchOneInto(p interface{}) (bool, error)
	Execute() Result
}

type StatementImpl struct {
	sa    StatementAcceptor
	built bool

	Statement string
	Bindings  []interface{}
	State     map[string]interface{}

	queryer Queryer
}

func (s *StatementImpl) FetchMap() ([]map[string]string, error) {
	if err := s.buildStatement(); err != nil {
		return nil, fmt.Errorf("failed to build sql : %w", err)
	}

	rows, err := s.queryer.Query(s.Statement, s.Bindings...)
	if err != nil {
		return nil, fmt.Errorf("failed to query : %w", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Warn(err.Error())
		}
	}()

	names, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	res := make([]map[string]string, 0)
	for rows.Next() {
		pmap := make(map[string]*string)
		vptrs := make([]interface{}, 0)
		for _, name := range names {
			var v string
			vptrs = append(vptrs, &v)
			pmap[name] = &v
		}

		if err := rows.Scan(vptrs...); err != nil {
			return nil, err
		}

		vmap := make(map[string]string)
		for k, v := range pmap {
			vmap[k] = *v
		}

		res = append(res, vmap)
	}
	return res, nil
}

func (s *StatementImpl) FetchInto(p interface{}) error {
	if err := s.buildStatement(); err != nil {
		return fmt.Errorf("failed to build sql : %w", err)
	}

	var isPtrElement bool

	sliceValue := reflect.ValueOf(p)
	if sliceValue.Kind() != reflect.Ptr {
		return ErrorMustBeAPtr
	}

	sliceValue = sliceValue.Elem()
	if sliceValue.Kind() != reflect.Slice {
		return ErrorMustBeASlice
	}

	// Ptrの場合は値に戻す
	elementType := sliceValue.Type().Elem()
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
		isPtrElement = true
	}

	rows, err := s.queryer.Query(s.Statement, s.Bindings...)
	if err != nil {
		return fmt.Errorf("failed to query : %w", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Warn(err.Error())
		}
	}()

	names, err := rows.Columns()
	if err != nil {
		return err
	}

	for rows.Next() {
		// elementはpointer
		element := reflect.New(elementType).Interface()

		vptrs, err := s.toFieldPtr(element, names)
		if err != nil {
			return err
		}
		if err := rows.Scan(vptrs...); err != nil {
			return err
		}

		// 元のSliceの値が値の場合はpointerから値に戻す
		elementValue := reflect.ValueOf(element)
		if !isPtrElement {
			elementValue = elementValue.Elem()
		}

		sliceValue.Set(reflect.Append(sliceValue, elementValue))
	}

	return nil
}

func (s *StatementImpl) FetchOneInto(p interface{}) (bool, error) {
	if err := s.buildStatement(); err != nil {
		return false, fmt.Errorf("failed to build sql : %w", err)
	}

	// 入力型をチェック
	v := reflect.ValueOf(p)
	if v.Kind() != reflect.Ptr {
		return false, ErrorMustBeAPtr
	}
	if v.IsNil() {
		return false, ErrorMustBeANonNilPtr
	}

	// Ptr => Struct変換
	ve := v.Elem()
	if ve.Kind() != reflect.Struct {
		return false, ErrorMustBeAStructPtr
	}

	rows, err := s.queryer.Query(s.Statement, s.Bindings...)
	if err != nil {
		return false, fmt.Errorf("failed to query : %w", err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Warn(err.Error())
		}
	}()

	names, err := rows.Columns()
	if err != nil {
		return false, err
	}

	if !rows.Next() {
		return false, nil
	}

	vptrs, err := s.toFieldPtr(p, names)
	if err != nil {
		return false, err
	}

	if err := rows.Scan(vptrs...); err != nil {
		return false, err
	}
	return true, nil
}

func (s *StatementImpl) toFieldPtr(p interface{}, names []string) ([]interface{}, error) {
	t := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(p)).Interface())

	name2index := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name2index[strings.ToLower(f.Name)] = i
		if tag, ok := f.Tag.Lookup("sqlike"); ok {
			name2index[strings.ToLower(tag)] = i
		}
	}

	val := reflect.ValueOf(p).Elem()

	vptrs := make([]interface{}, t.NumField())
	for i, name := range names {
		fi, ok := name2index[name]
		if !ok {
			return nil, fmt.Errorf("failed to find field for '%s'", name)
		}

		valueField := val.Field(fi)
		vptrs[i] = valueField.Addr().Interface()
	}

	return vptrs, nil
}

func (s *StatementImpl) Execute() Result {
	if err := s.buildStatement(); err != nil {
		return &BasicResult{err: fmt.Errorf("failed to build sql : %w", err)}
	}

	result, err := s.queryer.Execute(s.Statement, s.Bindings...)

	return &BasicResult{native: result, err: err}
}

func (s *StatementImpl) StatementAndBindings() (string, []interface{}, error) {
	if err := s.buildStatement(); err != nil {
		return "", nil, fmt.Errorf("failed to build sql : %w", err)
	}

	return s.Statement, s.Bindings, nil
}

func (s *StatementImpl) buildStatement() error {
	if s.built {
		return nil
	}

	steps := getSteps(s.sa)

	// RootStepがQueryerでなければバグのためpanic
	rootStep := steps[0]
	q, ok := rootStep.(Queryer)
	if !ok {
		return fmt.Errorf("RootStep(%T) is not a Queryer", rootStep)
	}

	s.State = make(map[string]interface{})
	s.queryer = q
	for _, step := range steps {
		if err := step.Accept(s); err != nil {
			return err
		}
	}

	s.Statement = strings.TrimSuffix(s.Statement, " ")
	s.built = true
	return nil
}

func getSteps(lastStep StatementAcceptor) []StatementAcceptor {
	revSteps := make([]StatementAcceptor, 0)

	current := lastStep
	for current != nil {
		revSteps = append(revSteps, current)
		current = current.Parent()
	}

	steps := make([]StatementAcceptor, 0)
	for i := len(revSteps) - 1; i >= 0; i-- {
		steps = append(steps, revSteps[i])
	}
	return steps
}

func getQueryer(lastStep StatementAcceptor) (Queryer, error) {
	steps := getSteps(lastStep)

	if len(steps) == 0 {
		return nil, ErrorNoSteps
	}

	q, ok := steps[0].(Queryer)
	if !ok {
		return nil, fmt.Errorf("RootStep(%T) is not a Queryer", q)
	}

	return q, nil
}

type Result interface {
	Error() error
	AffectedRows() (int64, error)
	LastInsertId() (int64, error)
}

type BasicResult struct {
	native sql.Result
	err    error
}

func (b BasicResult) Error() error {
	return b.err
}

func (b BasicResult) AffectedRows() (int64, error) {
	if b.err != nil {
		return 0, b.err
	}
	return b.native.RowsAffected()
}

func (b BasicResult) LastInsertId() (int64, error) {
	if b.err != nil {
		return 0, b.err
	}
	return b.native.LastInsertId()
}
