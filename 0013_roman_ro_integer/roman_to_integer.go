package roman_ro_integer

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0

	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
			result = result - romanMap[string(s[i])]

			continue
		}

		if string(s[i]) == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
			result = result - romanMap[string(s[i])]

			continue
		}

		if string(s[i]) == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
			result = result - romanMap[string(s[i])]

			continue
		}

		result = result + romanMap[string(s[i])]
	}

	return result + romanMap[string(s[len(s)-1])]
}
