package main

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	original := x
	reversed := 0

	for x > 0 {
		last := x % 10
		reversed = reversed*10 + last
		x = x / 10
	}
	return original == reversed
}

func main() {

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))

}
