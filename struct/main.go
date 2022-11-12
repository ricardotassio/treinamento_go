package main

import "fmt"

type Cliente struct {
	Name     string
	Age      int
	Activate bool
}

func main() {
	Ricardo := Cliente{
		Name:     "Ricardo",
		Age:      37,
		Activate: true,
	}
	fmt.Printf("Meu nome é %s, tenho %d e estou ativado? %t", Ricardo.Name, Ricardo.Age, Ricardo.Activate)

}
