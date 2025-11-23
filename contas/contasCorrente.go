package contas

type ContaCorrente struct {
	Titular string
	NumeroAgencia int
	NumeroConta int
	Saldo   float64
}

type contaDaSilvia struct {
	titular string
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Valor de depósito inválido", c.Saldo
	}
}

func(c * ContaCorrente) Transferir(valorDaTransferencia float64, ContaDestino  *ContaCorrente) bool{
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		ContaDestino.Depositar(valorDaTransferencia)
		return true
	} else{
		return false
	}
}