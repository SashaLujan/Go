package main

import (
	"fmt"
	"tpGoTudai2021/model"
)

func main() {

	r := model.NewResult("NN", 03, "001")
	fmt.Println(r)
	fmt.Print(model.GetStructure("TX03ABC"))
	fmt.Println("")
	fmt.Print(model.GetStructure("TX02ABC"))
}
