package math

func NStar(n int) int {
	if n <= 1 {
		return n
	}
	return n * NStar(n-1)
}
