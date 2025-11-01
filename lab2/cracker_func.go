package lab2

// This file handles all of the password cracking operation,
// main function is passwd_crack
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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