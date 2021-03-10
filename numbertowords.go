// Thispackage the English words for number input
package numbertowords

import "errors"

var words = [21]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"Sixteen",
	"seventeen",
	"Eighteen",
	"Ninteen",
	"twenty",
}

var tenwords = [10]string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

//this is the max number and this program validates
const maxNumber = 99999

//the purpose of this function is to Convert is the function to convert the  number to words in english
func Convert(number int) (string, error) {
	if number < 0 || number > maxNumber {
		return "", errors.New("The number is not valid")
	}

	result := ""
	thousends := number / 1000
	if thousends > 0 {
		result, _ = Convert(thousends)
		result = result + " thousend"
		number = number % 1000
		if number == 0 {
			return result, nil

		}
	}

	hundreds := number / 100

	if hundreds > 0 {
		result = words[hundreds] + " hundred "
		number = number % 100
		if number == 0 {
			return result, nil
		}
	}
	tens := number / 10
	units := number % 10

	if tens < 2 {
		return result + words[number], nil
	}

	if units == 0 {
		return result + tenwords[tens], nil
	}
	return result + tenwords[tens] + "" + words[units], nil
}
