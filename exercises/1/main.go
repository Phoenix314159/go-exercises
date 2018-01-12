package main

import "fmt"

func main() {
	h, even := half(8)
	fmt.Println(h, even)
}
func half(n int) (int, bool) {
	h := n / 2
	if n%2 == 0 {
		return h, true
	}
	return h, false
}
