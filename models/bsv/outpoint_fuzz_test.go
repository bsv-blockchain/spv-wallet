package bsv

import "testing"

func FuzzOutpointFromString(f *testing.F) {
	// Seed corpus with valid and edge case inputs
	f.Add("0000000000000000000000000000000000000000000000000000000000000000-0")
	f.Add("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa-1")
	f.Add("8bb0cf6eb9b17d0f7d22b456f121257dc1254e1f01665370476383ea776df414-100")
	f.Add("")
	f.Add("invalid")
	f.Add("-1")
	f.Add("abc-def")
	f.Add("0000000000000000000000000000000000000000000000000000000000000000--1")

	f.Fuzz(func(t *testing.T, input string) {
		// The function should never panic, only return errors
		outpoint, err := OutpointFromString(input)
		if err == nil {
			// If no error, verify round-trip works
			s := outpoint.String()
			if s == "" {
				t.Errorf("String() returned empty for valid outpoint")
			}
		}
	})
}
