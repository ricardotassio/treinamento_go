package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(51, 2))
}

func sum(a int, b int) int {
	return a + b
}
func sum2(a, b int) int {
	return a + b
}
func sum3(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}
