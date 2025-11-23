package main

import "fmt"

type ContaCorrente struct {
	titular string
	saldo   float64
}

type contaDaSilvia struct {
	titular string
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string{
	podeDepositar := valorDoDeposito >= 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso"
	} else {
		return "Valor de depósito inválido"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{
		titular: "Silvia",
		saldo:   500,
	}

	fmt.Println(contaDaSilvia.saldo)
	contaDaSilvia.Depositar(-1000)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDaSilvia.Sacar(1200))
}
