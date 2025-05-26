package main

import (
	"fmt"
	"math"
)

func hammingDistance(x int, y int) int {
	return countBinaryOnes(x ^ y)
}

func countBinaryOnes(num int) int {
	if num == 0 {
		return 0
	}

	count := 0

	for curNum, powOfTwo := float64(num), float64(0); curNum > 0; powOfTwo++ {
		twoToPow := math.Pow(2, powOfTwo)
		if int64(curNum/twoToPow)%2 != 0 {
			count++
			curNum -= twoToPow
		}
	}

	return count
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
