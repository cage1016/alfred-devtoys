package lib

import (
	"strconv"
	"strings"
)

//BinToOct
func BinToOct(b string) string {
	return baseConvert(b, 2, 8)
}

//BinToDec
func BinToDec(b string) string {
	return baseConvert(b, 2, 10)
}

//BinToHex
func BinToHex(b string) string {
	return strings.ToUpper(baseConvert(b, 2, 16))
}

//DecToBin
func DecToBin(b string) string {
	return baseConvert(b, 10, 2)
}

//DecToOct
func DecToOct(d string) string {
	return baseConvert(d, 10, 8)
}

//DecToHex
func DecToHex(b string) string {
	return strings.ToUpper(baseConvert(b, 10, 16))
}

//OctToBin
func OctToBin(b string) string {
	return baseConvert(b, 8, 2)
}

//OctToDec
func OctToDec(b string) string {
	return baseConvert(b, 8, 10)
}

//OctToHex
func OctToHex(b string) string {
	return strings.ToUpper(baseConvert(b, 8, 16))
}

//HexToBin
func HexToBin(b string) string {
	return baseConvert(b, 16, 2)
}

//HexToOct
func HexToOct(b string) string {
	return baseConvert(b, 16, 8)
}

//HexToDec
func HexToDec(b string) string {
	return baseConvert(b, 16, 10)
}

// baseConvert("12312", 10, 16)
// [2- 36]
func baseConvert(number string, fromBase, toBase int) string {
	i, err := strconv.ParseInt(number, fromBase, 0)
	if err != nil {
		return ""
	}

	return strconv.FormatInt(i, toBase)
}
