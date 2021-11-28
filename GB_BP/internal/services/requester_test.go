package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"testing"
	"time"
)

type Link struct {
	URL string
	Title string
}

func TestNewRequester(t *testing.T) {
	r := NewRequester(10 * time.Second)
	assert.NotNil(t, r, "Error on create new services, get nil.")
}

func TestRequester_Get(t *testing.T) {
	http.HandleFunc("/home/", func(w http.ResponseWriter, r *http.Request){
		type DataLink struct {
			Links []Link
		}

		links:= []Link{
			{
				URL: "https://yandex.ru",
				Title: "link 1",
			},
			{
				URL: "https://google.com",
				Title: "link 2",
			},
			{
				URL: "https://childURL3.com",
				Title: "link 3",
			},
		}
		tmpl, err := template.ParseFiles("../page/mocks/home.page.tmpl")
		if err != nil {
			t.Fatalf("Error on parse template: %v \n", err)
		}
		err = tmpl.Execute(w, &DataLink{Links: links})
		if err != nil {
			t.Fatalf("Error on execute template: %v \n", err)
		}
	})

	addr := "localhost:8080"
	url := "http://" +addr + "/home/"
	server := &http.Server{Addr: addr, Handler: nil}
	go func() {
		err := server.ListenAndServe()
		if err !=nil {
			t.Errorf("Error on start server: %v \n", err)
			return
		}
	}()

	time.Sleep(3 * time.Second)

	r := NewRequester(3 * time.Second)
	ctx := context.Background()

	pg, err := r.Get(ctx, url)
	if err != nil {
		t.Fatalf("Error on get page: %v \n", err)
	}

	assert.NotNil(t, pg, "Get nil page.")
}
