package utils

import (
	"groupie-tracker/datatypes"
	"strconv"
	"strings"
)

func Filter(search string) []datatypes.Artist {
	filtered := []datatypes.Artist{}

	for _, artist := range ArtistsJson {
		// Artist name search
		if Match(search, artist.Name) {
			filtered = append(filtered, artist)
		}

		// Member search
		for _, member := range artist.Members {
			if Match(search, member) {
				filtered = append(filtered, artist)
			}
		}

		locations := GetLocations(artist.Id)
		for _, location := range locations {
			city := strings.Split(location, "-")[0]
			country := strings.Split(location, "-")[1]
			if Match(city, search) || Match(country, search) {
				filtered = append(filtered, artist)
			}
		}

		// First album date search
		if strings.Contains(artist.FirstAlbum, search) {
			filtered = append(filtered, artist)
		}

		// Creation date search
		if strconv.Itoa(artist.CreationDate) == search {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}
