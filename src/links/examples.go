package links

import (
	"github.com/devlucky/fakelink/src/templates"
	"math/rand"
)

// Get a random link with random values from a defined set of mocks
func RandomLink() *Link {
	return ExampleLinks[rand.Int()%len(ExampleLinks)]
}

var ExampleLinks = []*Link{
	{
		Private: false,
		Values: &templates.Values{
			Title:       "Sharknado (TV Movie 2013)",
			Description: "Directed by Anthony C. Ferrante.  With Ian Ziering, Tara Reid, John Heard, Cassandra Scerbo. When a freak hurricane swamps Los Angeles, nature's deadliest killer rules sea, land, and air as thousands of sharks terrorize the waterlogged populace.",
			SiteName:    "IMDb",
			Type:        "video.movie",
			Url:         "http://www.imdb.com/title/tt2724064/",
			Image:       "https://images-na.ssl-images-amazon.com/images/M/MV5BOTE2OTk4MTQzNV5BMl5BanBnXkFtZTcwODUxOTM3OQ@@._V1_SY1000_CR0,0,712,1000_AL_.jpg",
		},
	},
	{
		Private: false,
		Values: &templates.Values{
			Title:       "Bloodhound Gang - The Bad Touch",
			Description: "Music video by Bloodhound Gang performing The Bad Touch. (C) 1999 Interscope Records",
			SiteName:    "YouTube",
			Type:        "video",
			Url:         "https://www.youtube.com/watch?v=xat1GVnl8-k",
			Image:       "https://i.ytimg.com/vi/xat1GVnl8-k/hqdefault.jpg",
		},
	},
	{
		Private: false,
		Values: &templates.Values{
			Title:       "EuroTrip (2004)",
			Description: "Directed by Jeff Schaffer, Alec Berg, David Mandel.  With Scott Mechlowicz, Jacob Pitts, Michelle Trachtenberg, Travis Wester. Dumped by his girlfriend, a high school grad decides to embark on an overseas adventure in Europe with his friends.",
			SiteName:    "IMDb",
			Type:        "video.movie",
			Url:         "http://www.imdb.com/title/tt0356150/",
			Image:       "https://images-na.ssl-images-amazon.com/images/M/MV5BMTIxNjcxMDUxN15BMl5BanBnXkFtZTYwNjAxNTM3._V1_.jpg",
		},
	},
	{
		Private: false,
		Values: &templates.Values{
			Title:       "Ali G Indahouse (2002)",
			Description: "Directed by Mark Mylod.  With Sacha Baron Cohen, Emilio Rivera, Gina La Piana, Dana de Celis. Ali G unwittingly becomes a pawn in the evil Chancellor's plot to overthrow the Prime Minister of Great Britain. However, instead of bringing the Prime Minister down, Ali is embraced by the nation as the voice of youth and 'realness', making the Prime Minister and his government more popular than ever.",
			SiteName:    "IMDb",
			Type:        "video.movie",
			Url:         "http://www.imdb.com/title/tt0284837/",
			Image:       "https://images-na.ssl-images-amazon.com/images/M/MV5BMTgxMTA5YmYtNTE0MC00Mzk1LWJkNTUtZjJiYzBjYjdlYTM4XkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_SY1000_CR0,0,675,1000_AL_.jpg",
		},
	},
	{
		Private: false,
		Values: &templates.Values{
			Title:       "Kakapo.js",
			Description: "A bunch of colleagues writing about swift, javascript, ruby, go, algorithms, performance and coding stories",
			SiteName:    "DevLucky",
			Type:        "website",
			Url:         "http://devlucky.github.io/kakapo-js",
			Image:       "http://devlucky.github.io/assets/images/logo.png",
		},
	},
}
