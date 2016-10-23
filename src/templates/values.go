package templates

// Values describe all the possible OpenGraph attributes a compliant website might have
type Values struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Type  string `json:"type"`
	Url   string `json:"url"`
	Image string `json:"image"`
}
