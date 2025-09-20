package util

import "unicode"

func IsValidShipName(shipName string) bool {
	runes := []rune(shipName)
	for r := 0; r < len(runes); r++ {
		if !unicode.IsDigit(runes[r]) {
			return true
		}
	}
	return false
}
