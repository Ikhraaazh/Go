package main

import "fmt"

func main() {
	var i int = 256
	var f float32 = float32(i)
	fmt.Println(f)

	var a int64 = 100000
	var b int32 = int32(a) // int64 -> int32
	var c int16 = int16(b) // int32 -> int16
	var d int8 = int8(c)   // int16 -> int8 (bisa overflow)

	fmt.Println(a, b, c, d)
}
