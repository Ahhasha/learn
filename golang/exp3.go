package main
import "fmt"

func main() {
	users := [3]string{"Ilya", "Masha", "Shiisha"}
	for i := 0; i < len(users); i++{
		fmt.Println(users[i])
	}
	}