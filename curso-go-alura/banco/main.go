package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {
	contaDoRafa := ContaCorrente{
		titular:    "Rafa",
		numAgencia: 589,
		numConta:   123456,
		saldo:      123.5,
	}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoRafa)
	fmt.Println(contaDaBruna)
}
