package links

import (
	"github.com/devlucky/fakelink/src/templates"
	"testing"
	"os"
	"reflect"
)

/*
	Generic test suite for stores
*/

func behavesLikeAStore(t *testing.T, adapter Store) {
	adapter.clear()
	testFindMissing(t, adapter)

	adapter.clear()
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

	if !reflect.DeepEqual(link.Values, values) {
		t.Error("Expected the link's values to be exactly the sames we stored")
	}
}

/*
	All implementations comply with the expected behavior
*/

func TestInMemoryStore(t *testing.T) {
	adapter := NewInMemoryStore()
	behavesLikeAStore(t, adapter)
}

func TestRedisStore(t *testing.T) {
	adapter := NewRedisStore(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASS"),
		0,
	)
	behavesLikeAStore(t, adapter)
}