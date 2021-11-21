package spotifyintegration

import (
  "crypto/rand"
  "encoding/hex"
  "github.com/gin-gonic/gin"
  "net/http"
  "net/url"
  "os"
)

// define relevant system environment variable names
const (
  CLIENTID = "CLIENTID"
  SECRETID = "SECRETID"
  REDIRECTURI = "REDIRECTURI"
)

type SpotifyServer struct {
  router      *gin.Engine
  laststate   string
}

func NewSpotifyServer() *SpotifyServer {
  svr := &SpotifyServer{}
  svr.router = gin.Default()
  svr.router.GET("/login", svr.login)
  return svr
}

func (svr *SpotifyServer) Run(address string) error {
  return svr.router.Run(address)
}

// first component of Spotify API auth flow:
// request authorization to access data
func (svr *SpotifyServer) login(ctx *gin.Context) {
  const responsetype string = "code"
  const scopes string = "user-read-recently-played"
  const pathprefix string = "https://accounts.spotify.com/authorize"
  state, stateerr := generateRandomState()
  // panic if we're unable to correctly create state
  if stateerr != nil {
    panic(stateerr)
  }
  // track last state generated at login
  svr.laststate = state
  // build the query
  authquery := url.Values{}
  authquery.Set("response_type", responsetype)
  authquery.Set("client_id", os.Getenv(CLIENTID))
  authquery.Set("scopes", scopes)
  authquery.Set("state", state)
  authquery.Set("redirect_uri", os.Getenv(REDIRECTURI))
  authlocation := url.URL{Path: pathprefix, RawQuery: authquery.Encode()}
  ctx.Redirect(http.StatusFound, authlocation.RequestURI())
}

func requestTokens(ctx *gin.Context) {
  var request SpotifyRequestTokens
  if err:= ctx.ShouldBindJSON(&request); err != nil {
    ctx.JSON(http.StatusBadRequest,errorResponse(err))
    return
  }
  // build the query
  authquery := url.Values{}
  authquery.Set("client_id", os.Getenv(CLIENTID))
  resp, err := http.Get("")
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

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
