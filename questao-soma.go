package main

import "fmt"

func main() {
	indice := 12
	soma := 0
	k := 1
	for k < indice {
		k = k + 1
		soma = soma + k
		fmt.Println(soma)
	}
}

//resultado seria 77
