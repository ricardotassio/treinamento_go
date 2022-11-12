package main

import "fmt"

func main() {
	salarios := map[string]int{"Ricardo": 1000, "João": 200, "Maria": 100}
	fmt.Println(salarios["Ricardo"])

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}
	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}

}
