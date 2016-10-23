package links

import "github.com/devlucky/fakelink/src/templates"

// A Link represents a certain template version and values. They are user-generated
type Link struct {
	Values  *templates.Values
}

// Get a random link with random values from a defined set of mocks
func RandomLink() (l *Link) {
	// TODO: Make this random
	l.Values = &templates.Values{
		Title: "Sharknado",
		Type:  "website",
		Url:   "http://www.imdb.com/title/tt2724064/",
		Image: "https://images-na.ssl-images-amazon.com/images/M/MV5BOTE2OTk4MTQzNV5BMl5BanBnXkFtZTcwODUxOTM3OQ@@._V1_SY1000_CR0,0,712,1000_AL_.jpg",
	}

	return
}

func NewLink(values *templates.Values) (*Link, error) {
	// TODO: Do some validations here, in case of injection

	return &Link{values}, nil
}
