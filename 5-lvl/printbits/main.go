package main
import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) == 2 {
		num, _ := strconv.Atoi(os.Args[1])
		res := ""
		if num == 0 {
			res = "00000000"
		}
		for len(res) != 8 {
			a := num % 2
			res = string(rune(a+48)) + res
			num /= 2
		}
		for _, j := range res {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}


/*package main
import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) == 2 {
		args := os.Args[1:]
		arr := make([]int, 8)
		n, _ := strconv.Atoi(args[0])
		for i := len(arr) - 1; i >= 0; i-- {
			arr[i] = n % 2
			n = n / 2
		}
		for _, w := range arr {
			z01.PrintRune(rune(w + 48))
		}
		z01.PrintRune('\n')
	}
}
*/