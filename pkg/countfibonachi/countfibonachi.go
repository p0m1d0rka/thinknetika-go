package countfibonachi

// CountFibonachi count fubonachi number
func CountFibonachi(n int) int {
	if n > 20 {
		return -1
	}
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	return CountFibonachi(n-1) + CountFibonachi(n-2)
}
