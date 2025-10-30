package lab2

import (
	"bufio"
	"fmt"
	"io"
	"lab/lab2/utils/crack"
	"os"
	"strings"
)

func Run() {

	ex0()
	line()
	ex1()
	line()
	ex2()
	line()
	ex3()
	line()
}

func line() {
	for i := 0; i < 30; i++ {
		print("-")
	}
	fmt.Println("")
}

func ex0() {
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

func readfromfile(passlist string) []string {

	f, err := os.Open(passlist)
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	var rl []string

	for {
		l, err := r.ReadString('\n')

		//EOF means end of file, let the loop break at the end of the file
		if err == io.EOF {
			break
		}
		check(err)

		// remove the newline both linux \n and windows \rn
		l = strings.TrimRight(l, "\n")

		rl = append(rl, l)
	}
	return rl
}

func openfiletowrite(vfile string) *os.File {

	// os.O_CREATE: create file if it doesn't exist
	// os.O-WRONLY: open for writing
	// os.O_TRUNC: overwrite
	f, err := os.OpenFile(vfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("Error: ", err, "\n")
		return nil
	}

	return f
}

func passwd_crack(passlist, vfile, pass string, algo func(string) string) {

	openfile := openfiletowrite(vfile)

	rl := readfromfile(passlist)

	for i := 0; i < len(rl); i++ {
		if algo(rl[i]) == pass {
			fmt.Println("Password is: ", rl[i])
			fmt.Fprintf(openfile, "%v. Checking: %v\n", i+1, rl[i])
			fmt.Fprintf(openfile, "Password is: %v\n", rl[i])
			break
		} else if i == len(rl)-1 {
			fmt.Println("Passowrd is not in the password list. :(")
			fmt.Fprintf(openfile, "Password is not in the password list. :(")
		} else {
			fmt.Fprintf(openfile, "%v. Checking: %v\n", i+1, rl[i])
		}
	}
	openfile.Close()
}

func ex1() {
	fmt.Println("Exercise 1:")
	passlist := "lab2/utils/nord_vpn.txt"
	vfile := "lab2/results/verbose_ex1.txt"
	pass := "6a85dfd77d9cb35770c9dc6728d73d3f"
	algo := crack.Md5

	passwd_crack(passlist, vfile, pass, algo)
}

func ex2() {
	fmt.Println("Exercise 2:")
	passlist := "lab2/utils/nord_vpn.txt"
	vfile := "lab2/results/verbose_ex2.txt"
	pass := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	algo := crack.Sha1

	passwd_crack(passlist, vfile, pass, algo)
}

func ex3() {
	fmt.Println("Exercise 3:")
	passlist := "lab2/utils/nord_vpn.txt"
	vfile := "lab2/results/verbose_ex3.txt"
	pass := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	algo := crack.Sha512

	passwd_crack(passlist, vfile, pass, algo)
}
