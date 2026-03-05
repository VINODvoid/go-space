package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Add arguments like number operator number")
		return
	}
	a := os.Args[1]
	b := os.Args[3]
	op := os.Args[2]
	aa, _ := strconv.ParseFloat(a, 64)
	bb, _ := strconv.ParseFloat(b, 64)
	switch op {
	case "+":
		fmt.Println(aa + bb)
	case "-":
		fmt.Println(aa - bb)
	case "*":
		fmt.Println(aa * bb)
	case "/":
		if bb == 0 {
			fmt.Println("cannot divide by zero")
			return
		}
		fmt.Println(aa / bb)
	default:
		fmt.Println("Invalid operator")
	}
}
