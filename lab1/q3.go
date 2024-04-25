package lab1

import (
	"regexp"
)

func LunhFomularChecker(input string) bool {
	// remove all not number characters
	re := regexp.MustCompile(`\D`)
	input = re.ReplaceAllString(input, "")
	// reverse the input
	reversedInput := reverse(input)
	// calculate the sum
	sum := 0
	for i, char := range reversedInput {
		charInt := int(char - '0')
		if i%2 == 1 {
			double := int(charInt) * 2
			if double < 9 {
				sum += double
			} else {
				sum += double - 9
			}
		} else {
			sum += int(charInt)
		}
	}
	return sum%10 == 0
}
func reverse(input string) string {
	reversed := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	return reversed
}

/*
func main() {
	input := " 4539 3195 0343 6467"
	if LunhFomularChecker(input) {
		println("Valid")
	} else {
		println("Invalid")
	}
}
*/
