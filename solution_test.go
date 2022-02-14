package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	want := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})
	got := GetMessage()

	if want != got {
		t.Errorf("GetMessage() = %v, want %v", got, want)
	}
}