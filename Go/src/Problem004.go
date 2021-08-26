package euler

import (
	"fmt"
	"time"
    "sort"
)

func Problem004() {
	start := time.Now()
	productArray := []int{}

	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			if isPallindrome_Problem004(i * j) {
				productArray = append(productArray, i*j)
			}
		}
	}

	end := time.Now()

	fmt.Printf("\nAnswer to Problem 4 : %d\n", max_Problem004(productArray))
	fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}

func isPallindrome_Problem004(num int) bool {
    s := fmt.Sprintf("%d", num)
	return s == reverse_Problem004(s)
}

func max_Problem004(productArray []int) int {

    sort.Ints(productArray)

    return productArray[len(productArray)-1]
}

func reverse_Problem004(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) { ret += string(r) }(v)
	}
	return
}
