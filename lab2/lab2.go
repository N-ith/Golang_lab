package lab2

import (
	"fmt"
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
