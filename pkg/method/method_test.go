package method

import (
	"testing"
)

func TestCheckStatusCode(t *testing.T) {
	err := CheckStatusCode("http://ya.ru")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestCheckText(t *testing.T) {
	err := CheckText("http://ya.ru")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
