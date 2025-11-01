package lab1

import (
	"fmt"
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"os"
)

func ex5() {

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
