package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable serve pára comparar
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"wesly": 1000, "joao": 2000, "maria": 3000}
	m2 := map[string]float64{"wesly": 100.0, "joao": 20.00, "maria": 300.0}
	m3 := map[string]MyNumber{"wesly": 1000, "joao": 2000, "maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))

}
