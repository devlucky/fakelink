package links

import (
	"github.com/devlucky/fakelink/src/helpers"
	"github.com/devlucky/fakelink/src/templates"
	"os"
	"reflect"
	"testing"
)

/*
	Generic test suite for stores
*/

func behavesLikeAStore(t *testing.T, store Store) {
	store.clear()
	testFindMissing(t, store)

	store.clear()
	testCreateAndFind(t, store)

	store.clear()
	testFindRandom(t, store)
}

func testFindMissing(t *testing.T, store Store) {
	link := store.Find("missing")
	if link != nil {
		t.Error("Expected .Find on a missing link to be nil")
	}
}

func testFindRandom(t *testing.T, store Store) {
	createLinks(t, store, 1, true)
	slug := store.FindRandom()
	if slug != "" {
		t.Error("Expected .FindRandom to return nil when the store is empty, or contains only private links")
	}

	slugs := createLinks(t, store, 10, false)

	random := false

	for i := 0; !random && i < 10; i++ {
		s1, s2 := store.FindRandom(), store.FindRandom()
		if s1 != s2 {
			random = true
		}

		if !helpers.StringInSlice(s1, slugs) || !helpers.StringInSlice(s2, slugs) {
			t.Error("Expected .FindRandom to return existing slugs")
		}
	}

	if !random {
		t.Error("Expected .FindRandom to provide random link slugs")
	}
}

func testCreateAndFind(t *testing.T, store Store) {
	values := templates.Values{Title: "something"}
	link, err := NewLink(values, true)
	if err != nil {
		t.Fatal("Not expecting .NewLink to fail. Instead, got", err)
	}

	slug := store.Create(link)

	link = store.Find(slug)
	if link == nil {
		t.Error("Expected .Find to find a link after .Create")
	}

	if !reflect.DeepEqual(link.Values, values) {
		t.Error("Expected the link's values to be exactly the sames we stored")
	}
}

func createLinks(t *testing.T, store Store, n int, private bool) []string {
	slugs := make([]string, n)

	for i := 0; i < n; i++ {
		link := RandomLink()
		link.Private = private
		slugs = append(slugs, store.Create(link))
	}

	return slugs
}

/*
	All implementations comply with the expected behavior
*/

func TestInMemoryStore(t *testing.T) {
	store := NewInMemoryStore()
	behavesLikeAStore(t, store)
}

func TestRedisStore(t *testing.T) {
	store := NewRedisStore(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASS"),
	)
	behavesLikeAStore(t, store)
}
