package main

import (
	"fmt"
	"math"
	"strings"
)

var numToWords = map[int]string{
	0:  "Zero",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

var thousandPowers = map[int]string{
	1: "Thousand",
	2: "Million",
	3: "Billion",
	4: "Trillion",
	5: "Quadrillion",
	6: "Quintillion",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	thousandsPower := 0
	for n := num; n >= 1000; {
		thousandsPower++
		n /= 1000
	}
	res, err := getStringForThousandsPower(num, thousandsPower)
	if err != nil {
		panic(err)
	}
	return res
}

func getStringForThousandsPower(num int, power int) (string, error) {
	if num < 1000 {
		return getStringForSubThousand(num)
	}

	powerStr, ok := thousandPowers[power]
	if !ok {
		return "", fmt.Errorf("thousandPowers[%d] not found", power)
	}

	var strParts []string

	// Get the first digits. For [1000; 9999], the 1st digit. For [10000; 99999], the first 2 digits. For [100000; 999999], the first 3 digits.
	var firstDigits int
	for firstDigits = num; firstDigits > 999; firstDigits = firstDigits / 1000 {
	}

	firstDigitsStr, err := getStringForSubThousand(firstDigits)
	if err != nil {
		return "", fmt.Errorf("getStringForSubThousand(%d): %w", firstDigits, err)
	}
	strParts = append(strParts, firstDigitsStr, powerStr)

	numWithFirstDigitsRemoved := num % int(math.Pow(1000, float64(power)))
	if numWithFirstDigitsRemoved > 0 {
		restOfString, err := getStringForThousandsPower(numWithFirstDigitsRemoved, power-1)
		if err != nil {
			return "", fmt.Errorf("getStringForThousandsPower(%d): %w", numWithFirstDigitsRemoved, err)
		}
		strParts = append(strParts, restOfString)
	}

	return strings.Join(strParts, " "), nil
}

func getStringForSubThousand(num int) (string, error) {
	if num >= 1000 {
		return "", fmt.Errorf("Numbers < 1000 expected")
	}

	if num < 20 { // Numbers [0; 19]
		str, ok := numToWords[num]
		if !ok {
			return "", fmt.Errorf("Could not find word for number %d", num)
		}
		return str, nil
	}

	if num < 100 { // Numbers [20; 99]
		subHundredStr, err := subHundredString(num)
		if err != nil {
			return "", fmt.Errorf("subHundredString for %d: %w", num, err)
		}
		return subHundredStr, nil
	}

	// Numbers [100; 999]
	// How many hundreds?
	numHundreds := num / 100
	numHundredsStr, ok := numToWords[numHundreds]
	if !ok {
		return "", fmt.Errorf("Could not find word for number %d", numHundreds)
	}

	firstDigitRemovedFromHundreds := num - 100*numHundreds
	if firstDigitRemovedFromHundreds == 0 { // E.g. 200
		return numHundredsStr + " Hundred", nil
	}
	subHundredStr, err := subHundredString(firstDigitRemovedFromHundreds)
	if err != nil {
		return "", fmt.Errorf("subHundredString for %d: %w", num, err)
	}

	return numHundredsStr + " Hundred " + subHundredStr, nil
}

func subHundredString(num int) (string, error) {
	if num < 20 {
		if str, ok := numToWords[num]; ok {
			return str, nil
		}
	}

	tenMultiplier := num / 10
	tensWord, ok := numToWords[tenMultiplier*10]
	if !ok {
		return "", fmt.Errorf("Could not find word for number %d", tenMultiplier)
	}

	singleDigit := num % 10
	if singleDigit == 0 {
		return tensWord, nil
	}
	singleDigitWord, ok := numToWords[singleDigit]
	if !ok {
		return "", fmt.Errorf("Could not find word for number %d", singleDigit)
	}

	return tensWord + " " + singleDigitWord, nil
}
