package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		n1, _ := strconv.Atoi(args[0])
		n2, _ := strconv.Atoi(args[1])
		if n1 <= 0 && n2 <= 0 {
			return
		}
		if n1 <= 0 || n2 <= 0 {
			return
		}
		if n1 > 0 && n2 > 0 {
			for n1 != n2 {
				if n1 > n2 {
					n1 = n1 - n2
				} else {
					n2 = n2 - n1
				}
			}
		}
		fmt.Println(n1)
	}
}


/*package main

import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	if len(os.Args) != 3 {
		return
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	if n1 < 0 || n2 < 0 {
		return
	}
	if n1 < n2 {
		n1, n2 = n2, n1
	}
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	fmt.Println(n1)
}
*/