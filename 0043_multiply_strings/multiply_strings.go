package multiply_strings

import "strconv"

func sumResults(results [][]int) string {
	// Initialize answer as a number from results.
	answer := results[len(results)-1]
	results = results[:len(results)-1]
	var newAnswer []int
	// Sum each digit from answer and result
	for _, result := range results {
		newAnswer = nil
		var carry int
		for i := 0; i < len(answer) || i < len(result); i++ {
			var digit1, digit2 int
			// If answer is shorter than result or vice versa, use 0 as the current digit.
			if i < len(result) {
				digit1 = result[i]
			} else {
				digit1 = 0
			}
			if i < len(answer) {
				digit2 = answer[i]
			} else {
				digit2 = 0
			}
			// Add current digits of both numbers.
			sum := digit1 + digit2 + carry
			// Set carry equal to the tens place digit of sum.
			carry = sum / 10
			// Append the ones place digit of sum to answer.
			newAnswer = append(newAnswer, sum%10)
		}
		if carry > 0 {
			newAnswer = append(newAnswer, carry)
		}
		answer = newAnswer
	}
	// Convert answer to a string.
	finalAnswer := ""
	for _, digit := range answer {
		finalAnswer += strconv.Itoa(digit)
	}
	return finalAnswer
}

func multiplyOneDigit(
	firstNumber string,
	secondNumberDigit rune,
	numZeros int,
) []int {
	// Insert zeros at the beginning based on the current digit's place.
	currentResult := make([]int, numZeros)
	var carry int
	// Multiply firstNumber with the current digit of secondNumber.
	for _, firstNumberDigit := range firstNumber {
		multiplication := (int(secondNumberDigit-'0') * int(firstNumberDigit-'0')) + carry
		// Set carry equal to the tens place digit of multiplication.
		carry = multiplication / 10
		// Append last digit to the current result.
		currentResult = append(currentResult, multiplication%10)
	}
	if carry > 0 {
		currentResult = append(currentResult, carry)
	}
	return currentResult
}

func multiply(firstNumber, secondNumber string) string {
	if firstNumber == "0" || secondNumber == "0" {
		return "0"
	}
	// Reverse both numbers.
	firstNumber = reverseString(firstNumber)
	secondNumber = reverseString(secondNumber)
	// For each digit in secondNumber, multipy the digit by firstNumber and
	// store the multiplication result (reversed) in results.
	var results [][]int
	for i, secondNumberDigit := range secondNumber {
		results = append(
			results,
			multiplyOneDigit(firstNumber, secondNumberDigit, i),
		)
	}
	// Add all the results in the results array, and store the sum in the answer string.
	answer := sumResults(results)
	// answer is reversed, so reverse it to get the final answer.
	return reverseString(answer)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
