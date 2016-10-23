package templates

import (
	"errors"
	"fmt"
	"html/template"
)

// A Store allows saving and storing HTML templates
type Store interface {
	// Get a particular template, by version
	GetTemplate(version string) (*template.Template, error)
}

/*
	In-memory
*/

// In-memory implementation of a template store
type InMemoryStore struct {
	templates map[string]*template.Template
}

// Create a new in-memory store
func NewInMemoryStore() *InMemoryStore {
	store := &InMemoryStore{make(map[string]*template.Template)}

	// TODO: Make this lazy
	store.initializeAllVersions()

	return store
}

func (store *InMemoryStore) GetTemplate(version string) (*template.Template, error) {
	if t, ok := store.templates[version]; ok {
		return t, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find template %s", version))
}

// This is the eager-loading version for templates
func (store *InMemoryStore) initializeAllVersions() {
	store.addTemplate("v1", v1)
}

// Adds a template with a certain name and content. If the template cannot be loaded, the system panics.
// Thus, it is supposed to be used internally, being certain it will work
func (store *InMemoryStore) addTemplate(name string, templateStr string) {
	t, err := template.New(name).Parse(templateStr)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error parsing the template: %s", err))
	}
	store.templates[name] = t
}
