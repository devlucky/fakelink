package helpers

import "testing"

func TestStringInSlice(t *testing.T) {
	slice := []string{"one", "three"}
	if !StringInSlice("one", slice) {
		t.Error("Expected string 'one' to be in the slice")
	}

	if StringInSlice("two", slice) {
		t.Error("Expected string 'two' not to be in the slice")
	}
}
