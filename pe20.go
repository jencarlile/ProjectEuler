package main

// Project Euler #20
// #n! means n  (n  1)  ...  3  2  1
// For example, 10! = 10  9  ...  3  2  1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
// Find the sum of the digits in the number 100!

import (
	"big"
	"flag"
	"fmt"
	"strconv"
)

var n *int = flag.Int("n", 100, "Number for which to compute sum of digits of n!")

func Factorial(val int) *big.Int {
	result := big.NewInt(int64(1))
	for mult := 1; mult <= val; mult++ {
		result = result.Mul(result, big.NewInt(int64(mult))) 
	} 
	return result
}

func SumOfDigits(s string) int {
	var sum int
	sum = 0
	for i := 0; i < len(s); i++ {
		digit, _ := strconv.Atoi(string(s[i]))
		sum += digit
	}
	return sum
} 


func main() {
	// Parse flag to get input value n
	flag.Parse()
	fmt.Println("N: ", *n)
	
	// Compute factorial of input value n
	nf := Factorial(*n)
	fmt.Println("Factorial: ", nf)
	
	// Convert factorial result to string
	nfstr := nf.String()
	fmt.Println("String: ", nfstr)
	
	// Find the sum of digits in factorial result
	sum := SumOfDigits(nfstr)
	fmt.Println("Sum: ", sum)
}