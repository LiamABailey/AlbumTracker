package spotifyintegration

// struct for recieving reuqest for access, refresh tokens
// from the application
type SpotifyRequestTokensRequestBody struct {
  Code    string  `json:"Code"`
  State   string  `json:"State"`
}

type SpotifyRefreshTokensRequestBody struct {
  RefreshToken string `json:"refresh_token"`
}

type SpotifyGetAlbumsRequestQuery struct {
  Before  int     `form:"before"`
  Limit   string  `form:"limit"`
}

type SpotifyRecentlyPlayedBody struct {
  Href  string                        `json:"href"`
  Items []SpotifyRecentlyPlayedTrack   `json:"items"`
  Limit int                           `json:"limit"`
  Next  string                        `json:"next"`
  Cursors  struct {
    Before  string  `json:"before"`
    After   string  `json:"after"`
  }                                   `json:"cursors"`
}

func (body SpotifyRecentlyPlayedBody) GetUniqueAlbums() map[string]SpotifyAlbum {
  // create a dict of ID: SpotifyAlbum
  albums := make(map[string]SpotifyAlbum)
  for _, track := range body.Albums {
    // because IDs are unique, we can re-assign
    albums[track.Album.ID] = track.Album
  }
  return albums
}

type SpotifyRecentlyPlayedTrack struct {
  Track struct {
    Album SpotifyAlbum         `json:"album"`
  }                            `json:"track"`
}

type SpotifyAlbum struct {
  Artists       []SpotifyRecentlyPlayedArtist `json:"artists"`
  Name          string                        `json:"name"`
  ReleaseDate   string                        `json:"release_date"`
  ReleasePercision  string                    `json:"release_date_percision"`
  ID            string                        `json:"id"`
}

type SpotifyRecentlyPlayedArtist struct {
  Name  string    `json:"name"`
  ID    string    `json:"id"`
}
