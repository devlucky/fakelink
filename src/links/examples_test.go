package links

import (
	"reflect"
	"testing"
)

func TestRandomLink(t *testing.T) {
	random := false

	for i := 0; !random && i < 5; i++ {
		l1, l2 := RandomLink(), RandomLink()
		if !reflect.DeepEqual(l1, l2) {
			random = true
		}
	}

	if !random {
		t.Error("Expected RandomLink to provide random links")
	}
}
