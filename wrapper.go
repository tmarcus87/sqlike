package sqlike

import "github.com/tmarcus87/sqlike/model"

func Record(value interface{}) *model.Record {
	return &model.Record{
		Value: value,
	}
}

func RecordWithSkip(value interface{}, cols ...model.Column) *model.Record {
	return &model.Record{
		Value: value,
		Skip:  cols,
	}
}

func RecordWithOnly(value interface{}, cols ...model.Column) *model.Record {
	return &model.Record{
		Value: value,
		Only:  cols,
	}
}

func Count(field model.ColumnField) model.ColumnField {
	return model.Count(field)
}

func Distinct(field model.ColumnField) model.ColumnField {
	return model.Distinct(field)
}
