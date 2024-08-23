package handlers

import (
	"errors"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"log"
	"net/http"
	"sort"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	filtered := utils.ArtistsJson

	search := r.URL.Query().Get("search")

	if search != "" {
		filtered = utils.Filter(search)
	}

	if r.Method != "GET" {
		ErrorHandler(w, errors.New("400 | Bad request"))
		return
	}

	// Get unique member counts
	memberCounts := make(map[int]bool)
	for _, artist := range utils.ArtistsJson {
		memberCounts[len(artist.Members)] = true
	}

	data := datatypes.ArtistsPage{
		Artists:      filtered,
		Locations:    utils.GetAllLocations(utils.LocationsJson),
		MemberCounts: make([]int, 0, len(memberCounts)),
	}

	for count := range memberCounts {
		data.MemberCounts = append(data.MemberCounts, count)
	}
	sort.Ints(data.MemberCounts)

	t, err := template.ParseFiles("templates/artists.html", "templates/artist_block.html")
	if err != nil {
		log.Println(utils.RED, "Error parsing artist template", utils.NONE)
		ErrorHandler(w, errors.New("500 | Internal server error: Could not parse template"))
		return
	}
	t.Execute(w, data)
}