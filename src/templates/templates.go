// Package templates contains the main templating logic
// the application uses to provide user-generated htmls
package templates

import (
	"fmt"
	"html/template"
)

// Values describe all the possible OpenGraph attributes a compliant website might have
type Values struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SiteName    string `json:"site_name"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Image       string `json:"image"`
}

const templateStr = `
<!DOCTYPE html>
<html prefix="og: http://ogp.me/ns#">
<head>
    {{if .Title}}
    <title>{{.Title}}</title>
    <meta property="og:title" content="{{.Title}}" />
    {{end}}

    {{if .SiteName}}<meta property="og:site_name" content="{{.SiteName}}" />{{end}}
    {{if .Description}}<meta property="og:description" content="{{.Description}}" />{{end}}
    {{if .Type}}<meta property="og:type" content="{{.Type}}" />{{end}}
    {{if .URL}}<meta property="og:url" content="{{.URL}}" />{{end}}
    {{if .Image}}<meta property="og:image" content="{{.Image}}" />{{end}}
</head>
</html>
`

// Get the OpenGraph template we will be using
func Get() *template.Template {
	t, err := template.New("template").Parse(templateStr)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error parsing the template: %s", err))
	}

	return t
}
