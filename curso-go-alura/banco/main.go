package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	podeSacar := saque <= c.saldo && saque > 0
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	}
	return "Não foi possível realizar o saque"
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(500))
	fmt.Println(contaDaSilvia.saldo)
}
