package piscine

// reversebits(00100110) -> 01100100
func ReverseBits(oct byte) byte {
	var res byte
	for i := 8; i > 0; i-- {
		res = (res << 1) | (oct & 1)
		oct = (oct >> 1)
	}
	return res
}

/*
package main
import "fmt"
func main() {
	fmt.Println(ReverseBits(38))
}
func ReverseBits(oct byte) byte {
	var res byte = 0
	for i := 8; i > 0; i-- {
		res = (res << 1) | (oct & 1)
		oct = (oct >> 1)
	}
	return res
}
*/
