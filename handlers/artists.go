package handlers

import (
	"errors"
	"groupie-tracker/datatypes"
	"groupie-tracker/utils"
	"html/template"
	"log"
	"net/http"
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

	data := datatypes.ArtistsPage{
		Artists:   filtered,
		Locations: utils.GetAllLocations(utils.LocationsJson),
	}

	t, err := template.ParseFiles("templates/artists.html", "templates/artist_block.html")
	if err != nil {
		log.Println(utils.RED, "Error parsing artist template", utils.NONE)
		ErrorHandler(w, errors.New("500 | Internal server error: Could not parse template"))
		return
	}
	t.Execute(w, data)
}
