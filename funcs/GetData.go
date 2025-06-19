package groupie

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)



type Artist struct {
	Id           uint     `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate uint     `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	ID             int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}

type LocationResponse struct {
    Index []Relation `json:"index"`
}

func ArtistData() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Failed to fetch artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	var artists []Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return artists
}

func LocationData() []Relation {
    resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
    if err != nil {
        log.Fatalf("Failed to fetch relation data: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response: %v", err)
    }

    var locResp LocationResponse
    if err := json.Unmarshal(body, &locResp); err != nil {
        log.Fatalf("Failed to unmarshal relation JSON: %v", err)
    }

    return locResp.Index
}
