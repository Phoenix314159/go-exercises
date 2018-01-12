package main

import "os"



func main() {
	src, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	src2, err := os.Stat("src.txt")
	if err != nil {
		panic(err)
	}
	size := src2.Size()
	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	bs := make([]byte, size)
	src.Read(bs)
	dst.Write(bs)
}
