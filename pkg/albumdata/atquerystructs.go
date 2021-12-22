package albumdata

import (
	"strings"
	"time"
)

// master query struct
type AlbumQuery struct {
	AlbumName           string    `json:"AlbumName"`
	AlbumNameExactMatch bool      `json:"AlbumNameExactMatch"`
	BandName            string    `json:"BandName"`
	BandNameExactMatch  bool      `json:"BandNameExactMatch"`
	Genres              []string  `json:"Genres"`
	YearStart           int       `json:"YearStart"`
	YearEnd             int       `json:"YearEnd"`
	DateAddedStart      time.Time `json:"DateAddedStart"`
	DateAddedEnd        time.Time `json:"DateAddedStart"`
}

// replaces space placeholders '%20' in all string fields
func (a *AlbumQuery) DecodeSpaces() {
	a.AlbumName = strings.Replace(a.AlbumName, "%20", " ", -1)
	a.BandName = strings.Replace(a.BandName, "%20", " ", -1)
	replGenre := make([]string, len(a.Genres))
	for i, g := range a.Genres {
		replGenre[i] = strings.Replace(g, "%20", " ", -1)
	}
	a.Genres = replGenre
}
