package links

import (
	"github.com/satori/go.uuid"
)

// A Store allows saving and retrieving user-generated links
type Store interface {
	Find(slug string) *Link
	Create(link *Link) string
}

/*
	In-memory
*/

// In-memory implementation of a template store
type InMemoryStore struct {
	links map[string]*Link
}

// Create a new in-memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{make(map[string]*Link)}
}

func (store *InMemoryStore) Find(slug string) *Link {
	return store.links[slug]
}

func (store *InMemoryStore) Create(link *Link) string {
	slug := uuid.NewV4().String()
	store.links[slug] = link

	return slug
}
