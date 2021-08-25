package euler

import (
	"fmt"
	"time"
)

func Problem001() {

	start := time.Now()
	ans := 0
	for i := 2; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			ans += i
		}
	}
	end := time.Now()
	fmt.Printf("\nAnswer to Problem 1 : %d\n", ans)
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())

}
