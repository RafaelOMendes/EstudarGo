package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func (c *ContaCorrente) Sacar(saque float64) (float64, string) {
	if saque <= c.saldo && saque > 0 {
		c.saldo -= saque
		return c.saldo, ""
	}
	return c.saldo, "Não foi possível realizar o saque"
}

func (c *ContaCorrente) Depositar(deposito float64) (float64, string) {
	if deposito < 0 {
		return c.saldo, "Não foi possivel realizar o deposito"
	}
	c.saldo += deposito
	return c.saldo, ""
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) bool {
	_, err := c.Sacar(valor)
	if err != "" {
		return false
	}
	_, err = destino.Depositar(valor)

	if err != "" {
		return false
	}

	return true
}

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	status := contaDoGustavo.Transferir(200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	// saque, err := contaDaSilvia.Sacar(500)
	// if err != "" {
	// 	log.Fatal(err)
	// }
	// fmt.Println(saque)

	// deposito, err := contaDaSilvia.Depositar(100)
	// if err != "" {
	// 	log.Fatal(err)
	// }
	// fmt.Println(deposito)
}
