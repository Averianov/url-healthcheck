package utils

import (
	"testing"
)

func TestReadJSON(t *testing.T) {
	urls, err := ReadJSON("test.json")
	if err != nil {
		t.Errorf("error read json file: %v", err)
	}
	t.Logf("%v", urls)
}
