package main

import (
	"testing"
	"tpGoTudai2021/model"
)

func TestParser(t *testing.T) {
	values := "TX040001"
	//values := "TB020"
	result, err := model.GetStructure(values)

	if result == nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
