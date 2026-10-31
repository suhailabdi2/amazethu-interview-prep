package main

// fibonacci sequence
func fibonacci(n int) (result int) {
	seq := []int{0, 1}
	for len(seq) <= n {
		seq = append(seq, (seq[len(seq)-1])+seq[len(seq)-2])
	}
	return seq[len(seq)-1]
}
func FibOptimal(n int) int {
	if n == 0 {
		return 0
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
