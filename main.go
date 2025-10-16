package main

import (
	"fmt"
)

func main () {
	lab1()
	lab2()
	lab3()
}

func line () {

	for i:=0; i<30; i++ {
		print("-")
	}
	print("\n")

}

func lab1 () {
// Task 1: Write a Go program to demonstrate the use of assignment operators. 
// The program should take two integer inputs and perform various assignment operations
// such as =, +=, -=, *=, /=, and %=. Display the result after each operation.
	fmt.Println("Lab1: ")
	
	var number int = 10
	
	number+=1
	fmt.Println(". 10+1=", number)

	number-=1
	fmt.Println(". 10-1=", number)

	number*=2
	fmt.Println(". 10*2=", number)
	
	number/=2
	fmt.Println(". 10/2=", number)

	number%=2
	fmt.Println(". 10%2=", number)
	line()
}

func lab2() {
// 	• Write a Go program to demonstrate the use of logical
// operators such as &&, ||, and !.
// • The program should take two integer inputs and evaluate
// logical expressions like:
// • both positive (&&)
// • one greater than the other (||)
// • not equal (!)

// Display whether each condition is true or false.

	fmt.Println("Lab 2:")

	var num1 int
	var num2 int
	var boolean bool

	print("Integer value for state1: ")
	fmt.Scanln(&num1)
	print("Integer value for state2: ")
	fmt.Scanln(&num2)


	fmt.Printf("state1: %v, state2: %v \n", num1, num2)

	boolean = (num1>0)&&(num2>0)
	println("state1&&state2", and(num1, num2))

	boolean = (num1>0)||(num2>0)
	println("state1||state2", boolean)
	
	boolean = (num1>0)!=(num2>0)
	println("state1!=state2", boolean)
	line()
}

func and (num1, num2 int) bool {
	return (num1>0) && (num2>0)
}


func lab3 () {
// 	• Write a Go program to demonstrate the use of bitwise and
// assignment operators. The program should perform AND, OR, XOR,
// NOT, left shift, and right shift operations on two integers using
// functions and display the results. Also, show the effect of
// assignment operators on variable values.
// • Pseudocode:
// main() {
// myXor(a,b)
// myNOT(a,b)
// myOR(a,b)
// myAND(a,b)
// }




}

