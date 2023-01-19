package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}
	//mostra apenas os dados sem o indice
	for _, v := range numeros {
		println(v)
	}
	//mostrar somente os inidices
	for k := range numeros {
		println(k)
	}

	//pegar numeros menores que 10
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	//lupe inifito
	for {
		println("ola mundo")
	}

}
