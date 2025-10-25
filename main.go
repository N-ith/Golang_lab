package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	lab1()
	line()
	lab2()
	line()
	lab3()
	line()
	lab4()
	line()
	lab5()
	line()
	lab6()
	line()
}

func line() {
	for i := 0; i < 30; i++ {
		print("-")
	}
	print("\n")

}

//------------------------------------------------- Lab 1

func lab1() {
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

//------------------------------------------------- Lab 2

func lab2() {
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

//------------------------------------------------- Lab 3

func lab3() {

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

// ------------------------------------------------- Lab 4
func lab4() {
	print("Lab4: \n")
	for {
		var choice int

		print("======== MIni Calculator ========\n")
		print("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit\n")
		print("Choose: ")
		fmt.Scanln(&choice)

		if choice == 6 {
			print("Exiting...\n")
			break
		}

		a, b , _:= inp_handle()
		
		switch choice {
		case 1:
			fmt.Printf("Result for %v+%v=%v\n", a, b, add(a, b))
		case 2:
			fmt.Printf("Result for %v-%v=%v\n", a, b, sub(a, b))
		case 3:
			fmt.Printf("Result for %vx%v=%v\n", a, b, multiply(a, b))
		case 4:
			if b != 0 {
				fmt.Printf("Result for %v/%v=%v\n", a, b, division(a, b))
			} else {
				fmt.Println("b cannot be 0")
			}
		case 5:
			var Ia, Ib int = int(a), int(b)
			if b != 0 {
				fmt.Printf("Result for %v%%%v=%v\n", a, b, remainder(Ia, Ib))
			} else {
				fmt.Println("b cannot be 0")
			}
		default:
			fmt.Printf("Invalid choice\n")
		}
	}
}

func inp_handle () (float64,float64, error){
// Check for input
	var a, b float64
	fmt.Print("Enter value for a: ")
	_, err := fmt.Scanln(&a)

	if err != nil {
		fmt.Println("a must be integer or float number!")
		return 0,0,err
	}

	fmt.Print("Enter value for b: ")
	_, err = fmt.Scanln(&b)

	if err != nil {
		fmt.Println("b must be integer or float number!")
		return 0,0,err
	}
	return a, b, nil
}

func add(a, b float64) float64 { return a + b }

func sub(a, b float64) float64 { return a - b }

func multiply(a, b float64) float64 { return a * b }

func division(a, b float64) float64 { return a / b }

func remainder(a, b int) int { return a % b }

//------------------------------------------------- Lab 5

func lab5() {

	print("Lab4: \n")

	print("Enter a string to start encoding: ")

	// use simple scanln or scanf or scan won't handle space
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mes := scanner.Text()

	// Define plain text data to byte
	Bmes := []byte(mes)

	fmt.Printf("Binary: %08b\n", Bmes)
	fmt.Println("Hexadecimal:", hex.EncodeToString(Bmes))
	fmt.Println("Base64:", base64.StdEncoding.EncodeToString(Bmes))
}

//------------------------------------------------- Lab 6

func lab6() {
	println("Lab6:")

	print("Enter a string to encrypt: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	plainText := scanner.Text()

	var key []byte
	print("Enter a key: ")
	fmt.Scanln(&key)
	
	encrypt := xorEncryption(plainText, key)
	fmt.Printf("Starting XOR encryption with key: %v <=> %v \n",key, base64.StdEncoding.EncodeToString([]byte(string(key))))
	
	// print the encryption as base64 for unprintable character support
	fmt.Print("Encrypted: ", base64.StdEncoding.EncodeToString([]byte(encrypt)), "\n")

	decrypt := xorEncryption(encrypt, key)
	fmt.Println("Decrypted", string(decrypt))
}

func xorEncryption(pText string, key []byte) string {

	cipher := make([]byte, len(pText))

	for i := range pText {
		cipher[i] = pText[i] ^ key[i % len(key)]
	}

	return string(cipher)
}
