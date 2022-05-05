package example

import (
	"fmt"
	"reflect"
)

func ExString() {
	str := "2233"

	typeOfStr := reflect.TypeOf(str)

	newTypeOfStr := reflect.New(typeOfStr)

	newTypeOfStr = newTypeOfStr.Elem()
	newTypeOfStr.SetString("123")
	fmt.Println(str)
	fmt.Println(newTypeOfStr.String())
}
