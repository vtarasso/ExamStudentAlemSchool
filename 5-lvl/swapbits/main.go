package main

import "fmt"

func main() {
	fmt.Println(SwapBits(128)) // 0100 0000 --> 0000 0100
}

func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}
