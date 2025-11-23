package main

import(

"fmt"
"github.com/Isaac002c/go-orientacao-a-objetos/contas"

)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}


status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(status)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	// fmt.Println(contaDaSilvia.saldo)
	// status, valor := contaDaSilvia.Depositar(1000)
	// fmt.Println(status, valor)
	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDaSilvia.Sacar(1200))

}
