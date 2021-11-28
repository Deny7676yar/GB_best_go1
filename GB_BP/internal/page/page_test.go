package page

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"testing"
)

type Link struct {
	URL string
	Title string
	expectedHTTPStatus int
}



func TestNewPage(t *testing.T) {
	type DataLink struct {
		Links []Link
	}

	links:= []Link{
		{
			URL: "https://yandex.ru",
			Title: "link 1",
			expectedHTTPStatus: http.StatusOK,
		},
		{
			URL: "https://google.com",
			Title: "link 2",
			expectedHTTPStatus: http.StatusOK,
		},
		{
			URL: "https://childURL3.com",
			Title: "link 3",
			expectedHTTPStatus: http.StatusOK,
		},
	}
	tmpl, err := template.ParseFiles("../page/mocks/home.page.tmpl")
	if err != nil {
		t.Fatalf("Error on parse template: %v \n", err)
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, &DataLink{Links: links})

	pg, err := NewPage(&buffer)
	assert.NotNil(t, pg, "Get nil page.")
}

func TestPage_GetTitle(t *testing.T) {
	type DataLink struct {
		Links []Link
	}

	links:= []Link{
		{
			URL: "https://yandex.ru",
			Title: "link 1",
			expectedHTTPStatus: http.StatusOK,
		},
		{
			URL: "https://google.com",
			Title: "link 2",
			expectedHTTPStatus: http.StatusOK,
		},
		{
			URL: "https://childURL3.com",
			Title: "link 3",
			expectedHTTPStatus: http.StatusOK,
		},
	}
	tmpl, err := template.ParseFiles("../page/mocks/home.page.tmpl")
	if err != nil {
		t.Fatalf("Error on parse template: %v \n", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, &DataLink{Links: links})

	pg, err := NewPage(&buffer)

	homeTitle := "Home page"
	gotTitle := pg.GetTitle()

	assert.Equal(t, homeTitle, gotTitle, "Not equal.\n Expected: %v \n Got: %v \n", homeTitle, gotTitle)

}

func TestPage_GetLinks(t *testing.T) {
	type DataLink struct {
		Links []Link
	}

	links:= []Link{
		{
			URL: "https://www.yandex.ru",
			Title: "link 1",
			expectedHTTPStatus: http.StatusOK,
		},
		{
			URL: "https://www.google.com",
			Title: "link 2",
			expectedHTTPStatus: http.StatusOK,
		},
		{
			URL: "https://www.mozilla.org",
			Title: "link 3",
			expectedHTTPStatus: http.StatusOK,
		},
	}
	tmpl, err := template.ParseFiles("../page/mocks/home.page.tmpl")
	if err != nil {
		t.Fatalf("Error on parse template: %v \n", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, &DataLink{Links: links})

	pg, err := NewPage(&buffer)

	wantLinks := []string{"https://www.yandex.ru", "https://www.google.com", "https://www.mozilla.org"}
	gotLinks := pg.GetLinks()

	assert.Equal(t, wantLinks, gotLinks, "Not equal.\n Expected: %v \n Got: %v \n", wantLinks, gotLinks)
}
