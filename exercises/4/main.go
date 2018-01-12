package main

import "fmt"

func largest(numbers ...int) int {
	var largest int // sets to zero value which is 0
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	l := largest(2, 3, 6, 7, 45, 67, 897)
	fmt.Println(l)
}
