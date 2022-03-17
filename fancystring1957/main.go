package fancystring1957

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}

	chars := []rune(s)

	precedingChar1 := chars[0]
	precedingChar2 := chars[1]

	result := []rune{precedingChar1, precedingChar2}

	for _, c := range chars[2:] {
		if !(precedingChar1 == precedingChar2 && precedingChar2 == c) {
			result = append(result, c)
		}
		precedingChar1 = precedingChar2
		precedingChar2 = c
	}

	return string(result)
}
