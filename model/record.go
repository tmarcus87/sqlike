package model

type Record struct {
	Skip  []Column
	Only  []Column
	Value interface{}
}
