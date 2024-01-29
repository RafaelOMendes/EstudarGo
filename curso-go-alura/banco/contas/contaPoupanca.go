package contas

import "EstudarGo/curso-go-alura/banco/clientes"

type ContaPoupanca struct {
	Titular                        clientes.Titular
	NumAgencia, NumConta, Operacao int
	saldo                          float64
}

func (c *ContaPoupanca) Sacar(saque float64) (float64, string) {
	if saque <= c.saldo && saque > 0 {
		c.saldo -= saque
		return c.saldo, ""
	}
	return c.saldo, "Não foi possível realizar o saque"
}

func (c *ContaPoupanca) Depositar(deposito float64) (float64, string) {
	if deposito < 0 {
		return c.saldo, "Não foi possível realizar o deposito"
	}
	c.saldo += deposito
	return c.saldo, ""
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
