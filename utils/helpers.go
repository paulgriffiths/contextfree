package utils

// IntRange returns an array of integers in the range 0...n-1.
func IntRange(n int) []int {
	r := []int{}
	for i := 0; i < n; i++ {
		r = append(r, i)
	}
	return r
}
