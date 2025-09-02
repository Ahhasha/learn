package main
import "fmt"

func main() {
	str := "Hello"
	for index, value := range str {
		fmt.Println("index:", index, "value:", value)
	}
}