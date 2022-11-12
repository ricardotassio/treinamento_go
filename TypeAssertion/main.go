package main

import "fmt"

func main() {
	var myVar interface{} = "Ricardo Tassio"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)

}
