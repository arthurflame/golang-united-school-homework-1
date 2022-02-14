package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	rendered := GetMessage()
	if GetMessage() != rendered {
		t.Error("GetMessage", rendered)
	}

}