package Parse

import (
	"testing"
)

func TestParseint(t *testing.T) {
	testcase := []struct {
		val      string
		expected int
	}{
		{"1", 1},
		{"0", 0},
		{"12", 12},
		{"123456784567823", 123456784567823},
	}

	for _, tc := range testcase {
		have, _ := Parseint(tc.val)
		if have != tc.expected {
			t.Errorf("Parse(%q) give %d, but expected %d", tc.val, have, tc.expected)
		}
	}
}

func BenchmarkParseint(b *testing.B) {
	testcase := []struct {
		val      string
		expected int
	}{
		{"1", 1},
		{"0", 0},
		{"12", 12},
		{"123456784567823", 123456784567823},
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testcase {
			Parseint(tc.val)
		}
	}
}
