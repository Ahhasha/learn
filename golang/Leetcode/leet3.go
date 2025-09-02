package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	for i, ch := range s {
		if i+1 < len(s) && romanMap[ch] < romanMap[rune(s[i+1])] {
			total -= romanMap[ch]
		} else {
			total += romanMap[ch]
		}
	}

	return total
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))

}
