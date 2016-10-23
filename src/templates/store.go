package templates

import (
	"fmt"
	"html/template"
	"errors"
)

type Store struct {
	templates map[string]*template.Template
}

func NewStore() (*Store) {
	store := &Store{make(map[string]*template.Template)}

	// TODO: Make this lazy
	store.initializeAllVersions()

	return store
}

// Get a particular
func (store *Store) GetTemplate(version string) (*template.Template, error) {
	if t, ok := store.templates[version]; ok {
		return t, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find template %s", version))
}

// This is the eager-loading version for templates
func (store *Store) initializeAllVersions() {
	store.addTemplate("v1", v1)
}

// Adds a template with a certain name and content. If the template cannot be loaded, the system panics.
// Thus, it is supposed to be used internally, being certain it will work
func (store *Store) addTemplate(name string, templateStr string) {
	t, err := template.New(name).Parse(templateStr)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error parsing the template: %s", err))
	}
	store.templates[name] = t
}