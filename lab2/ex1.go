package lab2

import (
	"fmt"
	"lab/lab2/utils/crack"
)

func ex1() {
	fmt.Println("Exercise 1:")
	passlist := "lab2/utils/nord_vpn.txt"
	vfile := "lab2/results/verbose_ex1.txt"
	pass := "6a85dfd77d9cb35770c9dc6728d73d3f"
	algo := crack.Md5

	passwd_crack(passlist, vfile, pass, algo)
}