package main

// For any N, let f(N) be the last five digits before the trailing zeroes in N!.
// For example,
// 9! = 362880 so f(9)=36288
// 10! = 3628800 so f(10)=36288
// 20! = 2432902008176640000 so f(20)=17664

// Find f(1,000,000,000,000)

import (
	"big"
	"fmt"
)

func Factorial(val *big.Int) *big.Int {
	result := big.NewInt(int64(1))
	
	for mult := big.NewInt(int64(1)); mult.Cmp(val) != 1; mult = mult.Add(mult, big.NewInt(int64(1))) {
		result = result.Mul(result, mult) 
	} 
	return result
}

func FindLastNonZeroDigit(val *big.Int) int {
	sval := val.String()
	slen := len(sval)
	fmt.Println("Len: ", slen)
	lastNonZero := slen
	
	for i := slen - 1; i >= 0; i-- {
		c := string(sval[i])		
		if c != string("0"){
			fmt.Println("i: ", i)
			lastNonZero = i
			break
		}
	}
	
	return lastNonZero
}


func main() {
	//bi := big.NewInt(12345678900000000)
	b := big.NewInt(int64(1000000))
	b = b.Mul(b, b)
	fmt.Println("B: ", b.String())
	
	bi := Factorial(b) //big.NewInt(int64(100))
	fmt.Println("Factorial: ", bi.String())
	idx := FindLastNonZeroDigit(bi)
	fmt.Println("Last NonZero: ", idx)
	
	
	sbi := bi.String()
	subsbi := sbi[idx-4:idx+1]
	
	fmt.Println("Substring: ", subsbi)
	
}