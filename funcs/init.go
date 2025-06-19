package groupie

import (
	"html/template"
	"log"
)

var (
	tmplHome         *template.Template
	tmplArtistDetail *template.Template
	tmpl404          *template.Template
)

func Init() {
	var err error
	tmplHome, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing index.html: %v", err)
	}

	tmplArtistDetail, err = template.ParseFiles("templates/ArtistDetail.html")
	if err != nil {
		log.Fatalf("Error parsing ArtistDetail.html: %v", err)
	}
	tmpl404, err = template.ParseFiles("templates/404.html")
	if err != nil {
		log.Fatalf("Error parsing 404.html: %v", err)
	}
}