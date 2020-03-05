package statement

import (
	"fmt"
	"github.com/tmarcus87/sqlike/model"
	"strings"
)

const (
	StateUpdateStmtSet = "UPDATE_STMT_SET"
)

type UpdateStep struct {
	parent StatementAcceptor
	table  model.Table
}

func (s *UpdateStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *UpdateStep) Accept(stmt *StatementImpl) error {
	stmt.Statement += fmt.Sprintf("UPDATE `%s` ", s.table.SQLikeTableName())
	return nil
}

type UpdateSetStep struct {
	parent StatementAcceptor
	column model.Column
	value  interface{}
}

func (s *UpdateSetStep) Parent() StatementAcceptor {
	return s.parent
}

func (s *UpdateSetStep) Accept(stmt *StatementImpl) error {
	if _, ok := stmt.State[StateUpdateStmtSet]; ok {
		stmt.Statement = strings.TrimSuffix(stmt.Statement, " ")
		stmt.Statement += ", "
	} else {
		stmt.Statement += "SET "
	}
	stmt.Statement += fmt.Sprintf("`%s` = ? ", s.column.SQLikeColumnName())
	stmt.Bindings = append(stmt.Bindings, s.value)
	stmt.State[StateUpdateStmtSet] = true
	return nil
}

type UpdateSetRecordStep struct {
	parent StatementAcceptor
	record *model.Record
}

func (s UpdateSetRecordStep) Parent() StatementAcceptor {
	return s.parent
}

func (s UpdateSetRecordStep) Accept(stmt *StatementImpl) error {

	stmt.Statement += "SET "

	setColumns := make([]string, 0)
	setBindings := make([]interface{}, 0)

	fvm, err := getColumnName2FieldValueMap(s.record.Value)
	if err != nil {
		return err
	}

	if len(s.record.Only) > 0 {
		// 指定されたカラムのみ変更する
		for _, onlyColumn := range s.record.Only {
			fv, ok := fvm[onlyColumn.SQLikeColumnName()]
			if !ok {
				return fmt.Errorf("struct field for '%s' is not found", onlyColumn.SQLikeColumnName())
			}
			setColumns = append(setColumns, onlyColumn.SQLikeColumnName())
			setBindings = append(setBindings, fv.Interface())
		}
	} else if len(s.record.Skip) > 0 {
		// 指定されたカラム以外を変更する

		skipColumnNames := make(map[string]struct{}, 0)
		for _, skipColumn := range s.record.Skip {
			skipColumnNames[skipColumn.SQLikeColumnName()] = struct{}{}
		}

		for colName, fv := range fvm {
			if _, ok := skipColumnNames[colName]; !ok {
				setColumns = append(setColumns, colName)
				setBindings = append(setBindings, fv.Interface())
			}
		}

	} else {
		// すべてのカラムを変更する
		fields, err := getOrderedColumnName(s.record.Value)
		if err != nil {
			return err
		}

		for _, field := range fields {
			fv := fvm[field]
			setColumns = append(setColumns, field)
			setBindings = append(setBindings, fv.Interface())
		}
	}

	setStmt := make([]string, 0)
	for _, column := range setColumns {
		setStmt = append(setStmt, fmt.Sprintf("`%s` = ?", column))
	}

	stmt.Statement += strings.Join(setStmt, ", ") + " "
	stmt.Bindings = append(stmt.Bindings, setBindings...)

	return nil
}
