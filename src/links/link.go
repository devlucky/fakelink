package links

import (
	"errors"
	"fmt"
	"github.com/devlucky/fakelink/src/templates"
	"github.com/extemporalgenome/slug"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
)

// A Link represents a certain template version and values. They are user-generated
type Link struct {
	Private bool             `json:"private"`
	Values  templates.Values `json:"values"`
}

// NewLink creates a new Link from its template values.
func NewLink(values templates.Values, private bool) (*Link, error) {
	if values.Title == "" {
		return nil, errors.New("A link's title is mandatory")
	}

	link := &Link{
		Values:  values,
		Private: private,
	}
	return link, nil
}

func generateSlug(link *Link) string {
	s := fmt.Sprintf("%.80s-%s.6", slug.Slug(link.Values.Title), uuid.NewV4().String())

	// Set all possible flags
	if link.Private {
		s = setFlags(s, privateFlag)
	}

	return s
}

/*
	FLAGS:
	Links can be flagged (e.g. as private, as recent, as part of a campaign...)
	Flags are used internally to identify links and provide features or change
	the behavior seasonally depending on the link's metadata
*/

const (
	privateFlag int = 1 << iota
)

func hasFlag(slug string, flag int) bool {
	parts := strings.Split(slug, "-")
	flagCollection, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		flagCollection = 0
	}

	return flagCollection&flag != 0
}

func setFlags(slug string, flags ...int) string {
	var flagCollection int

	for _, flag := range flags {
		flagCollection = flagCollection | flag
	}

	return slug + "-" + strconv.Itoa(flagCollection)
}
