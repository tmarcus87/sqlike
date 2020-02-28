package sqlike

import (
	"database/sql"
	"fmt"
	"github.com/tmarcus87/sqlike/logger"
	"reflect"
	"strings"
)

type StatementAcceptor interface {
	// DialectStatement RDBにより異なるSQL方言を取得します
	DialectStatement(st StatementType) string

	// Parent 親のStatementAcceptorを返します。親がない場合はnilです
	Parent() StatementAcceptor

	// Accept Statementを受け取ってクエリを組み立てます
	Accept(stmt *StatementImpl)
}

type Statement interface {
	StatementAndBindings() (string, []interface{})
	FetchMap() ([]map[string]string, error)
	FetchInto(p interface{}) error
	FetchOneInto(p interface{}) (bool, error)
	Execute() (sql.Result, error)
}

type StatementImpl struct {
	Statement string
	Bindings  []interface{}

	queryer Queryer
}

func (s *StatementImpl) FetchMap() ([]map[string]string, error) {

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
	var isPtrElement bool

	sliceValue := reflect.ValueOf(p)
	if sliceValue.Kind() != reflect.Ptr {
		return ErrorMustBeAPtr
	}
	// Ptrの場合は値に戻す
	sliceValue = sliceValue.Elem()
	if sliceValue.Kind() != reflect.Slice {
		return ErrorMustBeASlice
	}

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

func (s *StatementImpl) Execute() (sql.Result, error) {
	return s.queryer.Execute(s.Statement, s.Bindings...)
}

func (s *StatementImpl) StatementAndBindings() (string, []interface{}) {
	return s.Statement, s.Bindings
}

func buildStatement(lastStep StatementAcceptor) *StatementImpl {
	revSteps := make([]StatementAcceptor, 0)

	current := lastStep
	for current != nil {
		revSteps = append(revSteps, current)
		current = current.Parent()
	}

	// RootStepがQueryerでなければバグのためpanic
	rootStep := revSteps[len(revSteps)-1]
	q, ok := rootStep.(Queryer)
	if !ok {
		panic("RootStep is not a Queryer")
	}

	stmt := StatementImpl{queryer: q}
	for i := len(revSteps) - 1; i >= 0; i-- {
		revSteps[i].Accept(&stmt)
	}

	stmt.Statement = strings.TrimSuffix(stmt.Statement, " ")

	return &stmt
}
