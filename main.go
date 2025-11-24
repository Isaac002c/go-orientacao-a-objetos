package main

import (
	"fmt"
	// "github.com/Isaac002c/go-orientacao-a-objetos/clientes"
	"github.com/Isaac002c/go-orientacao-a-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64){
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface{
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(500)
	PagarBoleto(&contaDaLuiza, 200)

	fmt.Println(contaDaLuiza.ObterSaldo())
}
