package euler


func fibbProblem2(n int) int {
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
