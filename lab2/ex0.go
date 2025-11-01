package lab2

import (
	"bufio"
	"fmt"
	"lab/lab2/utils/crack"
	"os"
)

func ex0 () {
	fmt.Println("Exercise 1: ")
	var i1, i2 string

	fmt.Print("Please input value 1: ")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	i1 = scanner1.Text()

	fmt.Print("Please input value 2: ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	i2 = scanner2.Text()

	md51 := crack.Md5(i1)
	md52 := crack.Md5(i2)

	sha11 := crack.Sha1(i1)
	sha12 := crack.Sha1(i2)

	sha2561 := crack.Sha256(i1)
	sha2562 := crack.Sha256(i2)

	sha31 := crack.Sha3(i1)
	sha32 := crack.Sha3(i2)

	fmt.Printf("Hash(MD5): \nOutput A: %v \noutput B: %v \n=> ", md51, md52)
	checkhash(md51, md52)
	fmt.Printf("Hash(SHA1): \nOutput A: %v \noutput B: %v \n=> ", sha11, sha12)
	checkhash(sha11, sha12)
	fmt.Printf("Hash(SHA256): \nOutput A: %v \noutput B: %v \n=> ", sha2561, sha2562)
	checkhash(sha2561, sha2562)
	fmt.Printf("Hash(SHA3): \nOutput A: %v \noutput B: %v \n=> ", sha31, sha32)
	checkhash(sha31, sha32)
}


func checkhash(i1, i2 string) {

	if i1 == i2 {
		println("Matched! HASH BROKE!!!")
	} else {
		println("Not Matched!")
	}
}

func check(err error) {
	if err != nil {
		fmt.Print("Error: ", err, "\n")
	}
}
