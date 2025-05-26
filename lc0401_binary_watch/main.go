package main

import (
	"math/bits"
	"strconv"
)

var hours = [4]int{8, 4, 2, 1}
var minutes = [6]int{32, 16, 8, 4, 2, 1}

func readBinaryWatch(turnedOn int) []string {
	if turnedOn == 0 {
		return []string{"0:00"}
	}

	var res []string

	var i uint32
	for i = 0; i < 1024; i++ { // 1024 is 2^10, 10 here is the num of LEDs
		if bits.OnesCount32(i) == turnedOn {
			strRunes := []rune(numToByteStr(i))
			ts := timeStr(&strRunes)
			if ts != "" {
				res = append(res, ts)
			}
		}
	}

	return res
}

func numToByteStr(num uint32) string {
	s := strconv.FormatInt(int64(num), 2)
	sLen := len(s)
	if sLen < 10 {
		for i := 0; i < 10-sLen; i++ {
			s = "0" + s
		}
	}

	return s
}

// runes is a string of bits like 10010110
func timeStr(runes *[]rune) string {
	var hour int
	for i := 0; i < 4; i++ {
		if (*runes)[i] == '1' { // 1 is a bit (LED on)
			hour += hours[i]
			if hour >= 12 {
				return ""
			}
		}
	}

	var minute int
	for i := 4; i < 10; i++ {
		if (*runes)[i] == '1' { // 1 is a bit (LED on)
			minute += minutes[i-4]
			if minute > 59 {
				return ""
			}
		}
	}

	minuteStr := strconv.Itoa(minute)
	if len(minuteStr) == 1 {
		minuteStr = "0" + minuteStr
	}
	return strconv.Itoa(hour) + ":" + minuteStr
}
