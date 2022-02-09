package main

import (
	"encoding/json"
	"fmt"
)

type Human interface {
	speak()
}

type Person struct {
	First string
	last  string
}

func (person Person) speak() {
	fmt.Printf("Type: %T -- Name %v -- Last Name %v \n", person, person.First, person.last)
}

func doSomething(h Human) {
	h.speak()
}

func main() {

	p := Person{First: "James", last: "Bond"}

	doSomething(p)

	encodedP, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(encodedP), "------")
	//var person Person

	//decodedP, err := json.Unmarshal(encodedP, *person)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(decodedP)

	//p.speak()
	//fmt.Printf("Our type is: %tFirst ---- Value %v\n", p, p)
	//fmt.Println(p.First, p.last)
	//p.spe
	//var myType MyType = "Hola"
	//myType.testing()
	//
	//fmt.Println(&myType)
	//changeMe(&myType)
	//var choice uint8
	//words := []string{"car", "hello", "something"}
	//
	//fmt.Printf("%v \n", words)
	//
	//for index, word := range words {
	//	fmt.Printf("Index: %d --- Word: %s \n", index, word)
	//}
	//
	//intMap := map[int]string{
	//	1: "h",
	//	2: "e",
	//	3: "l",
	//}
	//
	//for key, value := range intMap {
	//	if value == "h" {
	//		delete(intMap, key)
	//	}
	//	fmt.Printf("Index: %d --- Word: %s \n", key, value)
	//}
	//fmt.Printf("%v", intMap)
	//fmt.Println("Enter which example you want to run (1-5):")
	//a := 42
	//fmt.Printf("%d \t %b \t %#x \n", a, a, a)
	//fmt.Printf("Binary:  %b \t Hex: %#x \t %d", 23, a, a)
	//_, _ = fmt.Scanf("%d", &choice)
	//playground.RunExample(choice)
}

//func (anyone MyType) testing() {
//	fmt.Println(anyone)
//}
//
//func changeMe(x *MyType) {
//	*x = "adios!"
//	x.testing()
//	fmt.Println(x)
//}
