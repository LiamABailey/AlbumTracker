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

type SpotifyRecentlyPlayedTrack struct {
  Track struct {
    Album struct {
      Artists       []SpotifyRecentlyPlayedArtist `json:"artists"`
      Name          string                        `json:"name"`
      ReleaseDate   string                        `json:"release_date"`
      ReleasePercision  string                    `json:"release_date_percision"`
    }                                                 `json:"album"`
  }                                                       `json:"track"`
}

type SpotifyRecentlyPlayedArtist struct {
  Name  string    `json:"name"`
  ID    string    `json:"id"`
}
