package links

import (
	"github.com/devlucky/fakelink/src/templates"
	"reflect"
	"testing"
)

func TestValidNewLink(t *testing.T) {
	values := &templates.Values{Title: "some-title"}
	link, err := NewLink(values, true)
	if err != nil {
		t.Error("Expected NewLink not to fail")
	}

	if link.Values != values {
		t.Error("Expected NewLink to create link with the supplied values")
	}

	if link.Private != true {
		t.Error("Expected NewLink to create link with the supplied privacy")
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

func TestSlugGeneration(t *testing.T) {
	l, err := NewLink(&templates.Values{}, true)
	if err != nil {
		t.Fatal("Unexpected error creating a new link", err)
	}

	s1, s2 := generateSlug(l), generateSlug(l)
	if s1 == s2 {
		t.Error("Expected two slugs to be different")
	}

	if !hasFlag(generateSlug(l), privateFlag) {
		t.Error("Expected slug for private link to contain private flag")
	}

	l.Private = false
	if hasFlag(generateSlug(l), privateFlag) {
		t.Error("Expected slug for public link not to contain private flag")
	}
}

func TestSlugFlags(t *testing.T) {
	slug := "some-slug"
	slug = setFlags(slug, privateFlag)

	if !hasFlag(slug, privateFlag) {
		t.Error("Expected slug to have private flag")
	}
}
