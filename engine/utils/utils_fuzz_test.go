package utils

import "testing"

func FuzzGetChildNumsFromHex(f *testing.F) {
	// Seed corpus with valid and edge case inputs
	f.Add("")
	f.Add("ababab")
	f.Add("8bb0cf6eb9b17d0f7d22b456f121257dc1254e1f01665370476383ea776df414")
	f.Add("00000000")
	f.Add("ffffffff")
	f.Add("12345678")
	f.Add("abc")
	f.Add("invalid")
	f.Add("zzzzzzzz")

	f.Fuzz(func(t *testing.T, input string) {
		// The function should never panic, only return errors
		result, err := GetChildNumsFromHex(input)
		if err == nil {
			// If no error, result should be non-nil
			if result == nil {
				t.Errorf("GetChildNumsFromHex returned nil result without error")
			}
		}
	})
}
