package links

import (
	"github.com/devlucky/fakelink/src/templates"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
)

// A Link represents a certain template version and values. They are user-generated
type Link struct {
	Private bool              `json:"private"`
	Values  *templates.Values `json:"values"`
}

func NewLink(values *templates.Values, private bool) (*Link, error) {
	// TODO: Do some validations here, in case of injection
	link := &Link{
		Values:  values,
		Private: private,
	}
	return link, nil
}

// TODO: Base this on the title instead of generating a UUID
func generateSlug(link *Link) string {
	slug := uuid.NewV4().String()

	// Set all possible flags
	if link.Private {
		slug = setFlags(slug, privateFlag)
	}

	return slug
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
