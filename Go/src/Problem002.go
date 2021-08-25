package euler

import (
	"fmt"
	"time"
)

func Problem002() {
	ans, i, start := 0, 0, time.Now()
	k := 1

	for k < 4000000 {
		ans += fibb_Problem2(i)
		i++
		k = fibb_Problem2(i)
	}

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 2 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func fibb_Problem2(n int) int {
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a += b
		a, b = b, a
	}
	if a%2 == 0 {
		return a
	}
	return 0
}
