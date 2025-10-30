package lab1

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func Run() {
	ex1_1()
	line()
	ex2_1()
	line()
	ex3_1()
	line()
	ex4_1()
	line()
	ex5_1()
	line()
	ex6_1()
	line()
}

func line() {

	for i := 0; i < 30; i++ {
		print("-")
	}
	print("\n")

}

//------------------------------------------------- Lab 1

func ex1_1() {
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

func ex2_1() {
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

func ex3_1() {

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
func ex4_1() {
	print("Lab4: \n")
	for {
		var choice int
		var a, b float32

		print("======== MIni Calculator ========\n")
		print("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit\n")
		print("Choose: ")
		fmt.Scanln(&choice)

		if choice == 6 {
			print("Exiting...\n")
			break
		}

		print("Enter a: ")
		fmt.Scanln(&a)
		print("Enter b: ")
		fmt.Scanln(&b)

		switch choice {
		case 1:
			fmt.Printf("Result for %v+%v=%v\n", a, b, add(a, b))
		case 2:
			fmt.Printf("Result for %v-%v=%v\n", a, b, sub(a, b))
		case 3:
			fmt.Printf("Result for %vx%v=%v\n", a, b, multiply(a, b))
		case 4:
			fmt.Printf("Result for %v/%v=%v\n", a, b, division(a, b))
		case 5:
			var Ia, Ib int = int(a), int(b)
			fmt.Printf("Result for %v%%%v=%v\n", a, b, remainder(Ia, Ib))
		default:
			fmt.Printf("Invalid choice")
		}
	}
}

func add(a, b float32) float32 { return a + b }

func sub(a, b float32) float32 { return a - b }

func multiply(a, b float32) float32 { return a * b }

func division(a, b float32) float32 { return a / b }

func remainder(a, b int) int { return a % b }

//------------------------------------------------- Lab 5

func ex5_1() {

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

// func lab6() {

// 	println("Lab6:")
// 	print("Enter a string to encrypt: ")

// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	plainText := scanner.Text()

// 	print("Enter a key to start the encryption: ")
// 	var key []byte
// 	fmt.Scanln(&key)
// 	// key := []byte(keyStr)

// 	encrypt := xorEncryption(plainText, key)
// 	println("Encrypted:",encrypt)

// 	de_encrypt, err := base64.StdEncoding.DecodeString(encrypt)
// 	if err != nil {
// 		print("error:", err)
// 		return
// 	}


// 	decrypt:= xorEncryption(string(de_encrypt), key)
// 	de_decrypt, err := base64.StdEncoding.DecodeString(decrypt)
// 	if err != nil {
// 		fmt.Print("Decrypt error: ", err, "\n")
// 		return
// 	}

// 	fmt.Println("Decrypted", string(de_decrypt))
// }

// func xorEncryption(pText string, key []byte) string {

// 	cipher := []byte(pText)

// 	for i := 0; i < len(pText); i++ {
// 		cipher[i] = pText[i] ^ key[i % len(key)]
// 	}

// 	return base64.StdEncoding.EncodeToString(cipher)
// }



func ex6_1() {

	println("Lab6:")

	print("Enter a string to encrypt: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	plainText := scanner.Text()

	print("Enter a key to start the encryption: ")
	var key byte
	fmt.Scanln(&key)

	encrypt := xorEncryption(plainText, key)
	println("Encrypted:",encrypt)

	de_encrypt, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		print("error:", err)
		return
	}

	decrypt:= xorEncryption(string(de_encrypt), key)
	de_decrypt, err := base64.StdEncoding.DecodeString(decrypt)
	if err != nil {
		fmt.Print("Decrypt error: ", err, "\n")
		return
	}

	fmt.Println("Decrypted", string(de_decrypt))
}

func xorEncryption(pText string, key byte) string {

	cipher := []byte(pText)
	print(key)
	for i := 0; i < len(pText); i++ {
		cipher[i] = pText[i] ^ key
	}

	return base64.StdEncoding.EncodeToString(cipher)
}
