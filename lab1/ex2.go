package lab1

import "fmt"

func ex2() {
	fmt.Println("Lab 2:")

	var num1 int
	var num2 int
	var boolean bool

	print("Integer value for num1: ")
	fmt.Scanln(&num1)
	print("Integer value for num2: ")
	fmt.Scanln(&num2)

	fmt.Printf("num1: %v, num2: %v \n", num1, num2)

	boolean = (num1 > 0) && (num2 > 0)
	println("num1&&num2:", and(num1, num2))

	boolean = (num1 > 0) || (num2 > 0)
	println("num1||num2:", boolean)

	boolean = (num1 > 0) != (num2 > 0)
	println("num1!=num2:", boolean)
}

func and(num1, num2 int) bool {
	return (num1 > 0) && (num2 > 0)
}
