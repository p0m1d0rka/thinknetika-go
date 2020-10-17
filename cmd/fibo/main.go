package main

import (
	countfibonachi "fibo/countfibonachi"
	"flag"
	"fmt"
)

func main() {
	// why i cant use nFlag in function CountFibonachi()?
	// var nFlag = flag.Int("n", 1, "number of finobacci to count to")
	var fiboNumber int

	flag.IntVar(&fiboNumber, "n", 1, "number of fibonacci to count to")
	flag.Parse()
	var result int
	result = countfibonachi.CountFibonachi(fiboNumber)
	if result > 0 {
		fmt.Printf("Fibo number of %v = %v", fiboNumber, result)
	} else {
		fmt.Printf("Too large number %v number shoud be less then 20", fiboNumber)
	}

}
