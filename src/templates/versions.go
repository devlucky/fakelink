package templates

const v1 = `
<!DOCTYPE html>
<html prefix="og: http://ogp.me/ns#">
<head>
    <title>{{.Title}}</title>
    <meta property="og:title" content="{{.Title}}" />
    <meta property="og:type" content="{{.Type}}" />
    <meta property="og:url" content="{{.Url}}" />
    <meta property="og:image" content="{{.Image}}" />
</head>
</html>
`
