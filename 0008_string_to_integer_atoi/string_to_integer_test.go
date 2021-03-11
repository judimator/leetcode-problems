package string_to_integer_atoi

import "testing"

type TestData struct {
	input  string
	output int
}

var tests = []TestData{
	{
		"42",
		42,
	},
	{
		"   -042",
		-42,
	},
	{
		"1337c0d3",
		1337,
	},
	{
		"0-1",
		0,
	},
	{
		"words and 987",
		0,
	},
	{
		"-91283472332",
		-2147483648,
	},
	{
		"91283472332",
		2147483647,
	},
	{
		"9223372036854775808",
		2147483647,
	},
	{
		"+1",
		1,
	},
	{
		"+-12",
		0,
	},
	{
		"-+12",
		0,
	},
	{
		"21474836460",
		2147483647,
	},
	{
		"   +0 123",
		0,
	},
	{
		"4193 with words",
		4193,
	},
	{
		"0  123",
		0,
	},
	{
		"-5-",
		-5,
	},
	{
		"18446744073709551617",
		2147483647,
	},
	{
		"-000000000000001",
		-1,
	},
}

func TestMyAtoi(t *testing.T) {
	for _, data := range tests {
		if result := myAtoi(data.input); result != data.output {
			t.Error("Failed, myAtoi")
		}
	}
}
