package main

import "fmt"

type ContaCorrente struct {
    titular string
    saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    if valorDoSaque > 0 && valorDoSaque <= c.saldo {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func main() {
    contaDaSilvia := ContaCorrente{
        titular: "Silvia",
        saldo:   500,
    }

    fmt.Println(contaDaSilvia.saldo) // Imprime o saldo inicial da conta
}