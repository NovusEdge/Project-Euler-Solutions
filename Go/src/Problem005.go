package euler

import (
	"fmt"
	"time"
)


func Problem005() {

	start := time.Now()

	ans := lcm_Problem005(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 5 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

//gcd_Problem005 : greatest common divisor (gcd_Problem005) via Euclidean algorithm
func gcd_Problem005(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

//lcm_Problem005 : find Least Common Multiple (lcm_Problem005) via gcd_Problem005
func lcm_Problem005(a, b int, integers ...int) int {
	result := a * b / gcd_Problem005(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm_Problem005(result, integers[i])
	}

	return result
}
