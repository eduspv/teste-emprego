package main

import (
	"fmt"
	"strings"
)

func stringCounter(str string) int {
	count := 0
	char1 := 'a'
	char2 := 'A'
	for _, c := range str {
		if c == char1 || c == char2 {
			count++
		}
	}
	return count
}

func StringContainsA(str string) {
	if strings.Contains(str, "A") {
		fmt.Println("sua string contem a letra (A)")
	} else if strings.Contains(str, "a") {
		fmt.Println("sua string contem a letra (a)")
	}

}

func main() {
	str := "eu sou vasco da gama meu bem"
	StringContainsA(str)
	fmt.Printf("Quantidade de vezes que aparece o caracter a ou A: %d", stringCounter(str))

}
