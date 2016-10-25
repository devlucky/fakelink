package links

import (
	"github.com/devlucky/fakelink/src/templates"
	"strings"
	"testing"
)

func TestValidNewLink(t *testing.T) {
	values := templates.Values{Title: "some-title"}
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

func TestInvalidNewLink(t *testing.T) {
	if _, err := NewLink(templates.Values{}, true); err == nil {
		t.Error("Expected NewLink to fail if passed a missing title")
	}
}

func TestSlugGeneration(t *testing.T) {
	l, err := NewLink(templates.Values{Title: "An Extravagant Title! :)"}, true)
	if err != nil {
		t.Fatal("Unexpected error creating a new link", err)
	}

	s1, s2 := generateSlug(l), generateSlug(l)
	if s1 == s2 {
		t.Error("Expected two slugs to be different")
	}

	if !strings.Contains(s1, "an-extravagant-title") {
		t.Error("Expected the slug to contain part of the title")
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
