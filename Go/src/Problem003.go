package euler

import (
	"fmt"
	"time"
    "math"
)

func Problem003() {
	start := time.Now()
	ans := maxPrime_Problem003(factors_Problem003(600851475143))
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 3 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

//Factors returns an array of the factors of a number
func factors_Problem003(n int64) (res []int64) {
	upper := int64((math.Sqrt(math.Abs(float64(n)))))
	for i := int64(1); i < upper; i++ {
		if n%i == 0 {
			res = append(res, i)
			res = append(res, int64(n/i))
		}
	}
	return
}

//reports the largest prime number that is in the array [arr] (used in problem 3)
func maxPrime_Problem003(arr []int64) int64 {
	flag := arr[0]
	for _, i := range arr {
		if i > flag && isPrime_Problem003(i) {
			flag = i
		}
	}
	return flag
}

//isPrime_Problem003 reports if a number is a prime number or not
func isPrime_Problem003(n int64) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	}
	upper := int64(math.Trunc(math.Sqrt(float64(n))))
	for i := int64(2); i <= upper; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
