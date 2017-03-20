package main

import (
	"reflect"
	"fmt"
)

func main() {
	type myStruct struct {
		Field1 int `alias:"f1" desc:"field number 1"`
		Field2 string `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}
	mys := myStruct{2, "Hello", 2.4}
	InspectStructType(&mys)
}

func InspectStructType(i interface{}){
	mysRValue := reflect.ValueOf(i)
	if mysRValue.Kind() != reflect.Ptr{
		return
	}
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct{
		return
	}
	mysRValue.Field(0).SetInt(15)
	mysRType := mysRValue.Type() //reflect.TypeOf(i)

	for i:=0;i<mysRType.NumField();i++ {
		fieltRType := mysRType.Field(i)
		fieldRValue := mysRValue.Field(i)
		fmt.Printf("Field Name: '%s', field type: '%s', field value: '%v' \n", fieltRType.Name, fieltRType.Type, fieldRValue.Interface())
		fmt.Println("Struct tags, alias: ", fieltRType.Tag.Get("alias"), " description: ", fieltRType.Tag.Get("desc"))
	}
}