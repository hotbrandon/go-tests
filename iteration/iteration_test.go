package iteration

import (
	"strings"
	"testing"
)

func Repeat(s string, count int) string {
	var sb strings.Builder
	for i := 0; i < count; i++ {
		sb.WriteString(s)
	}

	return sb.String()
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, repteated: %q", expected, repeated)
	}

}
