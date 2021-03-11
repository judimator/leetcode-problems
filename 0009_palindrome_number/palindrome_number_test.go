package palindrome_number

import "testing"

type TestData struct {
	input  int
	output bool
}

var tests = []TestData{
	{
		121,
		true,
	},
	{
		-121,
		false,
	},
	{
		10,
		false,
	},
	{
		123456,
		false,
	},
	{
		1234321,
		true,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, data := range tests {
		if result := isPalindrome(data.input); result != data.output {
			t.Error("Failed, isPalindrome")
		}
	}
}
