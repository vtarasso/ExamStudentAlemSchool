package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) == 2 {
		res := ""
		arg := os.Args[1]
		for i, w := range arg {
			if w == 'a' || w == 'e' || w == 'i' || w == 'o' || w == 'u' || w == 'A' || w == 'E' || w == 'I' || w == 'O' || w == 'U' {
				res = arg[i:] + arg[:i] + "ay"
				break
			} else {
				res = "No vowels"
			}
		}
		if res == "" {
			return
		}
		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune(10)
	}
}