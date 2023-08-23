package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ar := os.Args
	if len(ar) != 2 {
		os.Exit(0)
	}
	n, er := strconv.Atoi(ar[1])
	if er != nil {
		os.Exit(0)
	}
	if n < 2 {
		return
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Print(i, "*")
			n /= i
			i = 1
		}
	}
	fmt.Print(n)
	fmt.Println()
}

/*
package main
import (
	"fmt"
	"os"
	"strconv"
)
func fprime(value int) {
	if value == 1 || value == 0 {
		return
	}
	i := 2
	for value > 1 {
		if value%i == 0 {
			fmt.Print(i)
			value = value / i
			if value > 1 {
				fmt.Print("*")
			}
			i--
		}
		i++
	}
	fmt.Println()
}
func main() {
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			fprime(i)
		}
	}
}
*/
