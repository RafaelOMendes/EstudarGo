package main

import (
	"EstudarGo/curso-go-alura/banco/clientes"
	"EstudarGo/curso-go-alura/banco/contas"
	"fmt"
)

func main() {
	clienteDoBruno := clientes.Titular{Nome: "Bruno", CPF: "12345678900", Profissao: "Programador"}

	contaDoBruno := contas.ContaCorrente{Titular: clienteDoBruno, NumAgencia: 123, NumConta: 123456}

	contaDoBruno.Depositar(-100)
	fmt.Println(contaDoBruno.GetSaldo())
}
