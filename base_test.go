package base

import (
	"testing"
)

func TestEncode(t *testing.T) {
	const alphabet = "23456789BCDFGHJKLMNPQRSTVWXZ"
	{
		for i := 0; i <= 500; i++ {
			input := uint(i)
			output := ToString(input, alphabet)
			reversed := ToUInt(output, alphabet)
			if !(reversed == input) {
				t.Errorf("ToString/ToUInt test failed! Input %v, output %v, reversed %v.", input, output, reversed)
			}
		}
	}
}

func TestExample(t *testing.T) {
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	{
		input := uint(195948557195948557)
		output := ToString(input, alphabet)
		reversed := ToUInt(output, alphabet)

		if !(reversed == input) {
			t.Errorf("ToString/ToUInt test failed! Input: %v, output: %v, reversed: %v.", input, output, reversed)
		}
	}
}
