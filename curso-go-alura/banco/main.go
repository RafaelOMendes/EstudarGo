package main

import (
	"EstudarGo/curso-go-alura/banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (float64, string)
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.GetSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(200)
	PagarBoleto(&contaDaLuiza, 150)

	fmt.Println(contaDaLuiza.GetSaldo())
}
