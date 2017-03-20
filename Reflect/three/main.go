package main

import (
	//"fmt"
	"reflect"
	"fmt"
)

func main() {
	var x1 float32 = 5.8
	v := reflect.ValueOf(&x1) // *float32 ==> x1
	//v.SetFloat(2.2)
	fmt.Println("v settable?", v.CanSet())
	vpElem := v.Elem()
	fmt.Println("vpElem settable?", vpElem.CanSet())
	vpElem.SetFloat(2.2)
	fmt.Println(x1)
}
