package model

import "strconv"

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(Type string, Length int, Value string) Result {
	return Result{Type, Length, Value}
}

func changeType(values string) Result {
	t := values[0:2]
	l := values[2:4]
	v := values[4:]
	lInt, err := strconv.Atoi(l) //convierte string en int
	if err == nil {
		return NewResult(t, lInt, v)
	} else {
		return NewResult("Error", 00, "Error")
	}
}
