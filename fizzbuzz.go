package fizzbuzz

import "fmt"

func Fizzbuzz(number int) string {
	if number%15 == 0 {
		return "fizzbuzz"
	} else if number%3 == 0 {
		return "fizz"
	} else if number%5 == 0 {
		return "buzz"
	}
	return fmt.Sprintf("%d", number)
}
