package main

import "fmt"

func fibonacci(number int) []int {
	var list []int
	num1 := 0
	num2 := 1
	for {
		num3 := num1 + num2
		fmt.Printf("%d \n", num3)
		list = append(list, num3)
		if num3 > number {
			break
		}
		num1 = num2
		num2 = num3

	}
	return list
}

func IsNumberInFibonacci(number int) {
	list := fibonacci(number)
	for i := 0; i < number; i++ {
		if list[i] == number {
			fmt.Printf("seu numero %d está na sequencia", number)
		}
	}

}

func main() {
	number := 2
	IsNumberInFibonacci(number)

}
