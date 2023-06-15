package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	lastName  string
	Age       int
}

func main() {
	people := []Person{
		{"pat", "Patterson", 17},
		{"Tracey", "Bobbert", 23},
		{"Fred", "Fredson", 32},
	}
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("Anonymus Function: ", j)
		}(i)
	}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
