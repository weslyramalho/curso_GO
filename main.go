package main

import (
	"fmt"

	"github.com/weslyramalho/curso-go/matematica"
)

type ID int

var (
	a ID = 0
)

type pessoa interface {
	Desativar()
}
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}
func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("RESUSLTADO: ", s)

	//memoria -> endereço -> valor
	/*
		a := 10
		var ponteiro *int = &a
		*ponteiro = 20
		b := &a
		*b = 30

		wesly := Cliente{
			Nome:  "Wesly",
			Idade: 36,
			Ativo: true,
		}
		wesly.Desativar()
	*/
	//fmt.Print(wesly.Nome)
	//fmt.Println(sum(1, 5, 7, 8, 12, 45, 14))
}

// somar varios numeros
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
