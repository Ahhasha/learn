package main

import (
	"fmt"
	"reflect"
)

func main() {

	people1 := map[string]int{"Tom": 1, "Bob": 2, "Sam": 3}
	people2 := map[string]int{"Tom": 1, "Bob": 2, "Sam": 3}
	people3 := map[string]int{"Tom": 1, "Bob": 2, "Sam": 4}
	fmt.Println("people1 == people2:", reflect.DeepEqual(people1, people2))
	fmt.Println("people1 == people3:", reflect.DeepEqual(people1, people3))
}
