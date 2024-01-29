package contas

import "EstudarGo/curso-go-alura/banco/clientes"

type ContaCorrente struct {
	Titular    clientes.Titular
	NumAgencia int
	NumConta   int
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

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
