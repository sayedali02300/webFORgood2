package groupie

import (
	"encoding/json"

	"log"
	"net/http"
	"strconv"
	"strings"
)

// HomePage displays all artists or search results
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Handle404(w, r)
		return
	}

	log.Println("Endpoint Hit: HomePage")

	searchQuery := r.URL.Query().Get("search")
	allArtists := ArtistData()

	var filtered []Artist
	if searchQuery != "" {
		lowerSearch := strings.ToLower(searchQuery)
		for _, artist := range allArtists {
			if strings.Contains(strings.ToLower(artist.Name), lowerSearch) {
				filtered = append(filtered, artist)
			}
		}
	} else {
		filtered = allArtists
	}

	data := struct {
		Artists     []Artist
		SearchQuery string
	}{
		Artists:     filtered,
		SearchQuery: searchQuery,
	}

	if err := tmplHome.Execute(w, data); err != nil {
		http.Error(w, "Failed to render homepage", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}

// PostPage displays a detailed artist page based on submitted ID
func PostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	artistID := r.FormValue("ArtistId")
	id, err := strconv.Atoi(artistID)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var foundArtist Artist
	for _, a := range ArtistData() {
		if int(a.Id) == id {
			foundArtist = a
			break
		}
	}

	var foundRelation Relation
	for _, rel := range LocationData() {
		if rel.ID == id {
			foundRelation = rel
			break
		}
	}

	type ArtistWithLocations struct {
		Artist   Artist
		Relation Relation
	}

	err = tmplArtistDetail.Execute(w, ArtistWithLocations{Artist: foundArtist, Relation: foundRelation})
	if err != nil {
		http.Error(w, "Failed to render artist detail", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}

}

// SuggestHandler returns autocomplete suggestions for artist names
func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	allArtists := ArtistData()

	var suggestions []string
	for _, artist := range allArtists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			suggestions = append(suggestions, artist.Name)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(suggestions); err != nil {
		http.Error(w, "Failed to encode suggestions", http.StatusInternalServerError)
	}
}
