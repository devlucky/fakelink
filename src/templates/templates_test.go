package templates

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("The code panicked")
		}
	}()

	Get()
}

func TestExecuteTemplateWithValues(t *testing.T) {
	values := &Values{
		Title:       "some-title",
		Description: "some-description",
		SiteName:    "some-site-name",
		Type:        "some-type",
		Url:         "some-url",
		Image:       "some-image",
	}

	buf := new(bytes.Buffer)
	Get().Execute(buf, values)
	generatedTemplate := buf.String()

	expectToContain(
		t,
		generatedTemplate,
		values.Title,
		values.Description,
		values.Type,
		values.Url,
		values.Image,
	)
}

func TestExecuteTemplateWithoutValues(t *testing.T) {
	values := &Values{}

	buf := new(bytes.Buffer)
	Get().Execute(buf, values)
	generatedTemplate := buf.String()

	expectNotToContain(
		t,
		generatedTemplate,
		"title",
		"description",
		"site_name",
		"type",
		"url",
		"image",
	)
}

func expectToContain(t *testing.T, template string, values ...string) {
	for _, value := range values {
		if !strings.Contains(template, value) {
			t.Error(fmt.Sprintf("Expected generated template to include %s", value))
		}
	}
}

func expectNotToContain(t *testing.T, template string, tags ...string) {
	for _, tag := range tags {
		if strings.Contains(template, fmt.Sprintf("og:%s", tag)) {
			t.Error(fmt.Sprintf("Expected generated template not to include %s when one was not specified", tag))
		}
	}
}
