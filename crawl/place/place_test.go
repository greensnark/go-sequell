package place

import (
	"testing"
)

func TestCanonicalPlace(t *testing.T) {
	cases := [][]string{
		{"shoal:2", "Shoals:2"},
		{"Shoals:4", "Shoals:4"},
		{"V:4", "Vaults:4"},
		{"ice", "IceCv"},
		{"Labyrinth", "Lab"},
	}
	for _, test := range cases {
		actual := CanonicalPlace(test[0])
		if actual != test[1] {
			t.Errorf("CanonicalPlace(%v) == %v, expected %v",
				test[0], actual, test[1])
		}
	}
}

func TestStripPlaceDepth(t *testing.T) {
	cases := [][]string{
		{"Cave:5", "Cave"},
		{"Melkor", "Melkor"},
	}
	for _, test := range cases {
		actual := StripPlaceDepth(test[0])
		if actual != test[1] {
			t.Errorf("StripPlaceDepth(%v) == %v, expected %v",
				test[0], actual, test[1])
		}
	}
}