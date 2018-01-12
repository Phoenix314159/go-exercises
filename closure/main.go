package main

import (
	"fmt" // file level scope
	"./wrapper"
)
var x int  //package level scope -- sets x to zero value
func increment() int {
	x++
	return x
}

func main () {
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

	increment2 := wrapper.Wrapper()
	fmt.Println(increment2())
	fmt.Println(increment2())
}
