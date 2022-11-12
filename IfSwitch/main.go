package main

func main() {
	a := 10
	b := 2
	c := 3

	switch a {
	case 1:
		println("Este é o valor de A1")
	case 2:
		println("Este é o valor de A2")
	case 3:
		println("Este é o valor de A3")
	default:
		println("O valor de A não foi encontrado!")
	}
	if a < b && b < c {
		println("Correto!")
	} else {
		println("Errado!")
	}
}
