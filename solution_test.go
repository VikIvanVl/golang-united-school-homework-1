package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	value := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})
	message := GetMessage()

	if !strings.EqualFold(message, value) {
		t.Errorf(" Error unexpected result:\n\tValue: %q\n\tGot: %q", value, message)
	}
}
