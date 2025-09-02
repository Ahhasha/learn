package main

import (
	"fmt"
	"reflect"
)

func main() {

	slice1 := []int{1, 2, 3, 4}
	slice2 := []string{"Tom", "Bob", "Sam"}
	slice3 := []int{1, 2, 3}
	slice4 := []int{1, 2, 3, 4}

	fmt.Println("s1 == s4", reflect.DeepEqual(slice1, slice4))
	fmt.Println("s2 == s3", reflect.DeepEqual(slice2, slice3))
}
