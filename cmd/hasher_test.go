package cmd

import "testing"

func TestTruncateLabel(t *testing.T) {
	cases := []struct {
		name string
		in   string
		max  int
		want string
	}{
		{"shorter than max", "abc", 8, "abc"},
		{"exact max", "abcdefgh", 8, "abcdefgh"},
		{"ascii truncated", "abcdefghij", 8, "abcdefgh"},
		{"multibyte not split mid-rune", "héllo world", 8, "héllo wo"},
		{"emoji not split mid-rune", "🙂🙂🙂🙂🙂🙂🙂🙂🙂🙂", 8, "🙂🙂🙂🙂🙂🙂🙂🙂"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := truncateLabel(tc.in, tc.max)
			if got != tc.want {
				t.Errorf("truncateLabel(%q, %d) = %q, want %q", tc.in, tc.max, got, tc.want)
			}
		})
	}
}
