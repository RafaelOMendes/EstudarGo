package contas

type ContaCorrente struct {
	Titular    string
	NumAgencia int
	NumConta   int
	Saldo      float64
}

func (c *ContaCorrente) Sacar(saque float64) (float64, string) {
	if saque <= c.Saldo && saque > 0 {
		c.Saldo -= saque
		return c.Saldo, ""
	}
	return c.Saldo, "Não foi possível realizar o saque"
}

func (c *ContaCorrente) Depositar(deposito float64) (float64, string) {
	if deposito < 0 {
		return c.Saldo, "Não foi possivel realizar o deposito"
	}
	c.Saldo += deposito
	return c.Saldo, ""
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
