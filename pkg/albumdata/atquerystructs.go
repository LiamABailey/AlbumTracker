package albumdata

import (
	"time"
)


// master query struct
type AlbumQuery struct {
  AlbumName string `json:"AlbumName"`
	NameExactMatch bool `json:"NameExactMatch"`
	BandName string `json:"BandName"`
	BandNameExactMatch bool `json:"BandNameExactMatch"`
	Genres []string `json:"Genres"`
	YearStart int `json:"YearStart"`
	YearEnd int `json:"YearEnd"`
	DateAddedStart time.Time `json:"DateAddedStart"`
	DateAddedEnd time.Time `json:"DateAddedStart"`
	MaxResults int `json:"MaxResults"`
}
