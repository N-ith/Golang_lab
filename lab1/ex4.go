package lab1

import "fmt"

func ex4() {
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
