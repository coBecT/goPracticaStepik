package main

func main() {

}
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
