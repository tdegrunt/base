package base

import (
	"strings"
)

// Converts a string (in base defined by alphabet) to an uint (base 10)
//
// Example: short := ToUInt("etrGeuS3YW", "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// Output: "195948557195948557"
func ToUInt(val string, alphabet string) uint {
	var num uint

	base := uint(len(alphabet))

	for _, v := range val {
		num = num*base + uint(strings.IndexRune(alphabet, v))
	}

	return num
}

// Converts an uint (base 10) to a different base (defined by alphabet)
//
// Example: short := ToString(195948557195948557, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// Output: "etrGeuS3YW"
func ToString(num uint, alphabet string) string {
	base := uint(len(alphabet))
	result := ""
	if num == 0 {
		return alphabet[0:1]
	}
	for num > 0 {
		pos := uint(num % base)
		digit := alphabet[pos : pos+1]
		result = digit + result
		num /= base
	}
	return result
}
