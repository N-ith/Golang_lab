package lab2

import (
	"fmt"
	"lab/lab2/utils/crack"
)

func ex3() {
	fmt.Println("Exercise 3:")
	passlist := "lab2/utils/nord_vpn.txt"
	vfile := "lab2/results/verbose_ex3.txt"
	pass := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	algo := crack.Sha512

	passwd_crack(passlist, vfile, pass, algo)
}