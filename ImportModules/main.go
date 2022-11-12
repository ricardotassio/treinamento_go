package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ricardotassio/estudo-go/matematica"
)

func main() {
	soma := matematica.Sum(10, 20)
	fmt.Println("O resultado da soma é", soma)

	// Para que uma função, variável ou struct possa estar exportada ela precisará iniciar com latra maiúscula
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())

	fmt.Println(uuid.New())
}
