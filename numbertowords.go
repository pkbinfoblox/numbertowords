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

const MaxNumber = 99999

//the number to words package will convert the input number to string

func Convert(number int) (string, error) {
	if number < 0 || number > 99999 {
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
