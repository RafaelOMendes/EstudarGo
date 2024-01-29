package main

import (
	"EstudarGo/curso-go-alura/banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

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
