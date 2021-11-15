package spotifyintegration

import (
  "crypto/rand"
  "encoding/hex"
  "github.com/gin-gonic/gin"
  "net/http"
  "net/url"
  "os"
  "fmt"
)

// define relevant system environment variable names
const (
  CLIENTID = "CLIENTID"
  REDIRECTURI = "REDIRECTURI"
)


type SpotifyServer struct {
  router *gin.Engine
}

func NewSpotifyServer() (*SpotifyServer) {
  svr := &SpotifyServer{}
  svr.router = gin.Default()
  svr.router.GET("/login", login)
  // TODO svr.router.GET("/recently-played", getRecentlyPlayed)
  return svr
}

func (svr *SpotifyServer) Run(address string) error {
  return svr.router.Run(address)
}

func (svr *SpotifyServer) Use(hf gin.HandlerFunc) {
  svr.router.Use(hf)
}

// first component of Spotify API auth flow:
// request authorization to access data
func login(ctx *gin.Context) {
  const responsetype string = "code"
  const scopes string = "user-read-recently-played"
  const pathprefix string = "http://accounts.spotify.com/authorize"
  state, stateerr := generateRandomState()
  // panic if we're unable to correctly create state
  if stateerr != nil {
    panic(stateerr)
  }
  authquery := url.Values{}
  authquery.Set("response_type", responsetype)
  authquery.Set("client_id", os.Getenv(CLIENTID))
  authquery.Set("scopes", scopes)
  authquery.Set("state", state)
  authquery.Set("redirect_uri", os.Getenv(REDIRECTURI))
  authlocation := url.URL{Path: pathprefix, RawQuery: authquery.Encode()}
  fmt.Println(authlocation.RequestURI())
  ctx.Redirect(http.StatusFound, authlocation.RequestURI())

}


// generates a random hex-string (length 16)
// for saftey in the redirect
func generateRandomState() (string, error) {
  const nbytes int = 32
  b32 := make([]byte, nbytes)
  _, err := rand.Read(b32)
  if err != nil {
    return "", err
  }
  return hex.EncodeToString(b32), nil
}
