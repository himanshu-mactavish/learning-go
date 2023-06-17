package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head Node
	tail Node
}

type Mapi map[string]string

type Test struct {
	mape map[string]string
}

type newType struct {
	firstName string "firstName of Person"
	lastName  string `LastName of person`
}

func main() {
	var myLL = new(LinkedList)
	myLL.head.data = 10
	myLL.head.next = &myLL.tail
	fmt.Println(myLL)
	myStructMap := new(Test)
	myCustomMap := make(Mapi)
	fmt.Println(myStructMap)
	fmt.Println(myCustomMap)
	tt := new(newType)
	tt.firstName = "Himanshu"
	tt.lastName = "MacTavish"
	fmt.Println(reflect.TypeOf(tt))

}
