package lab1

import "fmt"

func ex1() {
	
	fmt.Println("Lab1: ")

	var number int = 10

	number += 1
	fmt.Println(". 10+1=", number)

	number -= 1
	fmt.Println(". 10-1=", number)

	number *= 2
	fmt.Println(". 10*2=", number)

	number /= 2
	fmt.Println(". 10/2=", number)

	number %= 2
	fmt.Println(". 10%2=", number)
}
