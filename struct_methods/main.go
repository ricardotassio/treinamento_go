package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}
type Cliente struct {
	Name     string
	Age      int
	Activate bool
	Endereco
}

func (c Cliente) Deactivate() {
	c.Activate = false
	fmt.Println("Ricardo foi desativado!")
}

func main() {
	Ricardo := Cliente{
		Name:     "Ricardo",
		Age:      37,
		Activate: true,
	}
	Ricardo.Cidade = "Feira de Santana"
	Ricardo.Deactivate()
	fmt.Printf("Meu nome é %s, tenho %d e estou ativado? %t, Cidade: %s", Ricardo.Name, Ricardo.Age, Ricardo.Activate, Ricardo.Cidade)

}
