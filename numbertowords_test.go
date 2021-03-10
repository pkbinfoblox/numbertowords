package numbertowords

import "testing"

func TestInvalidInput(t *testing.T) {

	result, err := Convert(-1)

	if result != "" || err == nil {
		t.Fatal("The wrong result for -1")
	}

	result, err = Convert(100000)
	if result != "" || err == nil {
		t.Fatal("The wrong result for -1")
	}

}

func TestUnits(t *testing.T) {
	testcase := map[int]string{
		0:   "zero",
		1:   "one",
		2:   "two",
		3:   "three",
		4:   "four",
		5:   "five",
		6:   "six",
		7:   "seven",
		8:   "eight",
		9:   "nine",
		90:  "ninety",
		915: "nine hundred fifteen",
	}
	for key, val := range testcase {
		result, err := Convert(key)
		if result != val || err != nil {
			t.Fatal("failed test for input", val, result, key)
		}
	}
}
