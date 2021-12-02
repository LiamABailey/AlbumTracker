package spotifyintegration

// struct for recieving reuqest for access, refresh tokens
// from the application
type SpotifyRequestTokens struct {
  Code    string `json:"Code"`
  State   string `json:"State"`
}
