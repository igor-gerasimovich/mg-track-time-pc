package utils

/**
 * SplitLast returns the last part of splitting by rune
 * Examples:
 *   val: name1-name2-name3, delimiter: -
 *       => result: name3
 *   val: name1|someData, delimiter: |
 *       => result: someData
 */
func SplitLast(str string, rune uint8) string {
	newStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == rune {
			break
		}
		newStr = string(str[i]) + newStr
	}

	return newStr
}

/**
 * SplitLastByte returns the last part of splitting by rune
 * Examples:
 *   val: name1-name2-name3, delimiter: -
 *       => result: name3
 *   val: name1|someData, delimiter: |
 *       => result: someData
 */
func SplitLastByte(str []byte, rune uint8) []byte {
	cursor := 0

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == rune {
			cursor = i + 1
			break
		}
	}

	return str[cursor:]
}
