package main
import "fmt"

func main() {
	str := "Hello"
	for index, value := range str {
		fmt.Printf("Index: %d, Value %c\n", index, value)
	}
}