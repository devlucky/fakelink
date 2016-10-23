package links

import (
	"github.com/devlucky/fakelink/src/templates"
	"testing"
)

/*
	Generic test suite for stores
*/

func behavesLikeAStore(t *testing.T, adapter Store) {
	testFindMissing(t, adapter)
	testCreateAndFind(t, adapter)
}

func testFindMissing(t *testing.T, adapter Store) {
	link := adapter.Find("missing")
	if link != nil {
		t.Error("Expected .Find on a missing link to be nil")
	}
}

func testCreateAndFind(t *testing.T, adapter Store) {
	values := &templates.Values{Title: "something"}
	link, err := NewLink(values)
	if err != nil {
		t.Error("Not expecting .NewLink to fail. Instead, got", err)
	}

	slug := adapter.Create(link)

	link = adapter.Find(slug)
	if link == nil {
		t.Error("Expected .Find to find a link after .Create")
	}

	if link.Values != values {
		t.Error("Expected the link's values to be exactly the sames we stored")
	}
}

/*
	The in-memory implementation complies with the expected behavior
*/

func TestInMemoryStore(t *testing.T) {
	adapter := NewInMemoryStore()
	behavesLikeAStore(t, adapter)
}
