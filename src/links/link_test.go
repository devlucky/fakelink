package links

import (
	"github.com/devlucky/fakelink/src/templates"
	"reflect"
	"testing"
)

func TestValidNewLink(t *testing.T) {
	values := &templates.Values{Title: "some-title"}
	link, err := NewLink(values)
	if err != nil {
		t.Error("Expected NewLink not to fail")
	}

	if link.Values != values {
		t.Error("Expected NewLink to create link with the supplied values")
	}
}

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
