package sqlike

import "errors"

var (
	ErrorReadonlySession  = errors.New("readonly session can execute only SELECT")
	ErrorMustBeASlice     = errors.New("must be a slice")
	ErrorMustBeAPtr       = errors.New("must be a pointer")
	ErrorMustBeANonNilPtr = errors.New("must be a non-nil pointer")
	ErrorMustBeAStructPtr = errors.New("must be a pointer to struct")
)
