package piscine

func Itoa(n int) string {
		res := ""
	if n == 0 {
		res = ""
	} else if n < 0 {
		for n != 0 {
			res = string(rune(-n%10 + 48)) + res
			n = n / 10
		}
		res = string('-') + res
	} else {
		for n != 0 {
			res = string(rune(n%10+48)) + res
			n = n / 10
		}
	}
	return res
}
