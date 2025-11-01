package lab1

import "fmt"

func ex3() {

	println("Lab3: ")
	var num1 int = 2
	var num2 int = 3

	fmt.Printf("num1: %v, num2: %v \n", num1, num2)
	fmt.Println("myXOR:", myXOR(num1, num2))
	fmt.Println("myOR:", myOR(num1, num2))
	fmt.Println("myAND:", myAND(num1, num2))
	fmt.Println("myNOT:", myNOT(num1))
	fmt.Println("leftShift:", leftShift(num1, num2))
	fmt.Println("rightShift:", rightShift(num1, num2))

}

func myXOR(a, b int) int { return a ^ b }

func myOR(a, b int) int { return a | b }

func myAND(a, b int) int { return a & b }

func myNOT(a int) int { return ^a }

func leftShift(a, b int) int { return a << b }

func rightShift(a, b int) int { return a >> b }
