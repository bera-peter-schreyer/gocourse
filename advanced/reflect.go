package main

import (
	"fmt"
	"reflect"
)


type Greeter struct {}

func (g Greeter) Greet(str string) string {
	return "Hello, " + str
}

func reflection() {
 g := Greeter{}
 t := reflect.TypeOf(g)
 v := reflect.ValueOf(g)
 fmt.Println("Type:", t)
 for i := range t.NumMethod() {
	method := t.Method(i)
	fmt.Printf("Method %d: %s\n", i, method.Name)
 }

 m := v.MethodByName("Greet")
 results := m.Call([]reflect.Value{reflect.ValueOf("World")})
 fmt.Println("Results:", string(results[0].String()))
}

// // ====== WORKING WITH STRUCTS AND FIELDS ======
// // Capitalize the field to make it editable by reflect
// type person struct {
// 	Name string
// 	age  int
// }

// func main() {
// 	p := person{
// 		Name: "John",
// 		age:  30,
// 	}
// 	v := reflect.ValueOf(p)
// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}
// 	v1 := reflect.ValueOf(&p).Elem()
// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() {
// 		nameField.SetString("Jane")
// 	}
//  This will not work because age is unexported
// 	ageField := v1.FieldByName("age")
// 	if ageField.CanSet() {
// 		ageField.SetInt(25)
// 	}
// 	fmt.Printf("Updated person: %+v\n", p)
// }

// 	y := "hello"
// 	v = reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
// 	fmt.Println(v.Type())
// 	fmt.Println(v2.Type())
// 	fmt.Println("Original value:", v.String())
// 	v.SetString("world")
// 	fmt.Println("Updated value:", v.String())

// 	var itf interface{} = "howdy"
// 	v3 := reflect.ValueOf(itf)

// 	if v3.Kind() == reflect.String {
// 		fmt.Println("itf is a string", v3.String())
// 	} else {
// 		fmt.Println("itf is not a string")
// 	}
// }