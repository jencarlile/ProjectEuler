//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 
// 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
	"flag"
	"fmt"
)

var n *int = flag.Int("n", 1000, "Number for which to find the sum of all multiples of 3 or 5 less than n")

func main() {
	flag.Parse()
	fmt.Println("N: ", *n)

	sum := 0
	for i:= 0; i < *n; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			sum += i
		}
	}

	fmt.Println("Sum: ", sum)
}