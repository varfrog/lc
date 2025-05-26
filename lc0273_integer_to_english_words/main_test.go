package main

import (
	"testing"
)

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{0, "Zero"},
		{1, "One"},
		{9, "Nine"},
		{10, "Ten"},
		{11, "Eleven"},
		{15, "Fifteen"},
		{20, "Twenty"},
		{21, "Twenty One"},
		{30, "Thirty"},
		{42, "Forty Two"},
		{58, "Fifty Eight"},
		{99, "Ninety Nine"},

		{100, "One Hundred"},
		{101, "One Hundred One"},
		{110, "One Hundred Ten"},
		{115, "One Hundred Fifteen"},
		{200, "Two Hundred"},
		{999, "Nine Hundred Ninety Nine"},

		{1000, "One Thousand"},
		{1001, "One Thousand One"},
		{1010, "One Thousand Ten"},
		{1100, "One Thousand One Hundred"},
		{1234, "One Thousand Two Hundred Thirty Four"},
		{999999, "Nine Hundred Ninety Nine Thousand Nine Hundred Ninety Nine"},

		{1000000, "One Million"},
		{1000001, "One Million One"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{1999999, "One Million Nine Hundred Ninety Nine Thousand Nine Hundred Ninety Nine"},

		{1000000000, "One Billion"},
		{2147483647, "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven"},
	}

	for _, tt := range tests {
		result := numberToWords(tt.num)
		if result != tt.expected {
			t.Errorf("numberToWords(%d) = %q, want %q", tt.num, result, tt.expected)
		}
	}
}
