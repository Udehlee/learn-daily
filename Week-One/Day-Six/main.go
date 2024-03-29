package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)

	i, j := 0, len(strX)-1

	for i < j {
		if strX[i] != strX[j] {
			return false
		}
		i++
		j--
	}

	return true
}
func main() {

	x := 123
	fmt.Println(isPalindrome(x))

}
