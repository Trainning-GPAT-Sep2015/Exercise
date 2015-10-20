package parseint_test

import (
	"parseInt"
	"testing"
)

func TestParse(t *testing.T) {
	testcases := []struct {
		in   string
		want int
	}{
		{"100000", 100000},
		{"-29111", -29111},
		{"2822188", 2822188},
		{"2", 2},
		{"282218111", 282218111},
	}

	for _, c := range testcases {
		got, _ := parseint.Parse(c.in)
		if got != c.want {
			t.Errorf("ParseInt(%q) == %q, want %d", c.in, got, c.want)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseint.Parse("112")
	}
}
