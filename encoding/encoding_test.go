package encoding

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		given    string
		expected string
	}{
		{"foo", "foo"},
		{"a b", "a%20b"},
		{"'%", "%27%25"},
	}

	for i, test := range tests {
		got := URIEncode(test.given)
		if got != test.expected {
			t.Errorf("test[%d]: expected %s, got %s", i, test.expected, got)
		}
	}
}
