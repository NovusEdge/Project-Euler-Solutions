package euler

import (
	"fmt"
	"time"
)

func Problem002() {
	ans, i, start := 0, 0, time.Now()
	k := 1

	for k < 4000000 {
		ans += fibbProblem2(i)
		i++
		k = fibbProblem2(i)
	}

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 2 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}
