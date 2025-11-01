package lab2

import (
	"fmt"
	"lab/lab2/utils/crack"
)

func ex2() {
	fmt.Println("Exercise 2:")
	passlist := "lab2/utils/nord_vpn.txt"
	vfile := "lab2/results/verbose_ex2.txt"
	pass := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	algo := crack.Sha1

	passwd_crack(passlist, vfile, pass, algo)
}