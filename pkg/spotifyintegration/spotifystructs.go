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
