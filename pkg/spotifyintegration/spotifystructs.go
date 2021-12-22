package spotifyintegration

import (
	"strconv"
	"strings"
)

// struct for recieving reuqest for access, refresh tokens
// from the application
type SpotifyRequestTokensRequestBody struct {
	Code  string `json:"Code"`
	State string `json:"State"`
}

type SpotifyRefreshTokensRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

type SpotifyGetAlbumsRequestQuery struct {
	Before int    `form:"before"`
	Limit  string `form:"limit"`
}

type SpotifyRecentlyPlayedBody struct {
	Href    string                       `json:"href"`
	Items   []SpotifyRecentlyPlayedTrack `json:"items"`
	Limit   int                          `json:"limit"`
	Next    string                       `json:"next"`
	Cursors struct {
		Before string `json:"before"`
		After  string `json:"after"`
	} `json:"cursors"`
}

func (body SpotifyRecentlyPlayedBody) GetUniqueAlbums() map[string]SpotifyAlbumTrackInfo {
	// create a dict of ID: SpotifyAlbum
	albums := make(map[string]SpotifyAlbumTrackInfo)
	for _, track := range body.Items {
		// because IDs are unique, we can re-assign
		artists := make([]string, 0)
		for _, artist := range track.Track.Album.Artists {
			artists = append(artists, artist.Name)
		}
		//gather release year, defaulting to zero
		year := 0
		if track.Track.Album.ReleasePrecision == "day" {
			// take the first four chars from YYYY-MM-DD string
			year, _ = strconv.Atoi(track.Track.Album.ReleaseDate[0:4])
		} else if track.Track.Album.ReleasePrecision == "year" {
			year, _ = strconv.Atoi(track.Track.Album.ReleaseDate)
		}
		barSepArtists := strings.Join(artists, "|")
		albums[track.Track.Album.ID] = SpotifyAlbumTrackInfo{barSepArtists, track.Track.Album.Name, year}
	}
	return albums
}

type SpotifyRecentlyPlayedTrack struct {
	Track struct {
		Album SpotifyAlbum `json:"album"`
	} `json:"track"`
}

type SpotifyAlbum struct {
	Artists          []SpotifyRecentlyPlayedArtist `json:"artists"`
	Name             string                        `json:"name"`
	ReleaseDate      string                        `json:"release_date"`
	ReleasePrecision string                        `json:"release_date_precision"`
	ID               string                        `json:"id"`
}

// Album information formatted for output by getLastAlbums
type SpotifyAlbumTrackInfo struct {
	Artists     string `json:"artists"`
	Name        string `json:"name"`
	ReleaseYear int    `json:"release_date"`
}

type SpotifyRecentlyPlayedArtist struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
