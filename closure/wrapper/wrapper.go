package wrapper


func Wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
