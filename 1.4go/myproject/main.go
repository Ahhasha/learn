package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Обычный текст")
	color.Red("Красный текст")
	color.Green("Зеленый текст")
	color.Blue("Синий текст")
}
