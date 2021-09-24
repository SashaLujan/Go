package model

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(Type string, Length int, Value string) Result {
	return Result{Type, Length, Value}
}

const minLen = 5

func GetStructure(values string) (*Result, error) {
	if len(values) >= minLen {
		t := values[0:2]
		l := values[2:4]
		v := values[4:]
		lInt, err := strconv.Atoi(l) //convierte string en int
		if err == nil && lInt == len(v) {
			aux := NewResult(t, lInt, v)
			return &aux, nil
		} else {
			return nil, errors.New("Cadena invalida: " + values)
		}
	} else {
		return nil, errors.New("Cadena invalida: " + values)
	}

}
