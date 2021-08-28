package euler

import (
	"fmt"
	"time"
)


func Problem006() {

	start := time.Now()
	ans := sumSquare_Problem006(100) - squareSum_Problem006(100)
	end := time.Now()

	fmt.Printf("\nAnswer to Problem 6 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}


// reports the sum of the first *n* squares (used in problem 6)
func squareSum_Problem006(n int) int {
	k := 0
	for i := 1; i <= n; i++ {
		k += i * i
	}
	return k
}

//reports the square of the sum of first n numbers (used in problem 6)
func sumSquare_Problem006(n int) int {
	var k int
	k = n * (n + 1) / 2
	return k * k
}
