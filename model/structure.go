package model

type Result struct {
	Type   string
	Value  string
	Length int
}

func NewResult(types string, value string, length int) Result {
	return Result{types, value, length}
}
