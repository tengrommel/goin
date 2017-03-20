package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x1 float32 = 5.7
	inspectIfTypeFloat(x1)
	type myFloat float64
	var  x3  myFloat = 5.7
	v := reflect.ValueOf(x3)
	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
}

func inspectIfTypeFloat(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println("type: ", v.Type()) // same as reflect.TypeOf(x1)
	fmt.Println("Is type is float64?", v.Kind() == reflect.Float64)
	fmt.Println("Float Value:", v.Float())
	x2 := v.Float()
	v2 := reflect.ValueOf(x2)
	fmt.Println("Is v type float64?", v.Kind() == reflect.Float64)
	fmt.Println("Type of v2", v2.Type())

	interfaceValue := v.Interface()

	switch t := interfaceValue.(type) {
	case float32:
		fmt.Println("Original float32 value", t)
	}

}
