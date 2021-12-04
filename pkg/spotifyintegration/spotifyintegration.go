package spotifyintegration

import (
  "crypto/rand"
  "encoding/hex"
  "github.com/gin-gonic/gin"
  gincors "github.com/gin-contrib/cors"
  "net/http"
  "net/url"
  "os"
  "fmt"
  "io"
  b64 "encoding/base64"
  "errors"
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
  // using the cors default
  svr.router.Use(gincors.Default())
  svr.router.GET("/login", svr.login)
  svr.router.POST("/token", svr.requestTokens)
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
  const contenttype string = "application/x-www-form-urlencoded"
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
  authquery.Set("scope", scopes)
  authquery.Set("state", state)
  authquery.Set("redirect_uri", os.Getenv(REDIRECTURI))
  fmt.Println(authquery.Encode())
  fmt.Println(scopes)
  authlocation := url.URL{Path: pathprefix, RawQuery: authquery.Encode()}
  fmt.Println(authlocation.RequestURI())
  ctx.Redirect(http.StatusFound, authlocation.RequestURI())
}

func (svr *SpotifyServer) requestTokens(ctx *gin.Context) {
  const granttype string = "authorization_code"
  const contenttype string = "application/x-www-form-urlencoded"
  const pathprefix string = "https://accounts.spotify.com/api/token"

  var request SpotifyRequestTokens
  if err:= ctx.ShouldBindJSON(&request); err != nil {
    ctx.JSON(http.StatusBadRequest,errorResponse(err))
    return
  }
  // if returned state != state at login
  if request.State != svr.laststate {
    emessage := "Request State does not match Login State"
    ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(emessage)))
  }

  // build the query
  authquery := url.Values{}
  //authquery.Set("client_id", os.Getenv(CLIENTID))
  //authquery.Set("client_secret", os.Getenv(SECRETID))
  authquery.Set("grant_type", granttype)
  authquery.Set("code", request.Code)
  authquery.Set("redirect_uri", os.Getenv(REDIRECTURI))
  authlocation := url.URL{Path: pathprefix, RawQuery: authquery.Encode()}

  client := &http.Client{}
  req, _ := http.NewRequest("POST", authlocation.RequestURI(), nil)
  // set the required headers
  req.Header.Set("Content-Type", contenttype)
  authb64 := buildAuthString(os.Getenv(CLIENTID), os.Getenv(SECRETID))
  req.Header.Set("Authorization", authb64)
  fmt.Println(authlocation.RequestURI())
  resp, _ := client.Do(req)
  defer resp.Body.Close()
  body, _ := io.ReadAll(resp.Body)
  //fmt.Println(body)
  ctx.IndentedJSON(http.StatusOK, string(body))
}

// format the authorization string
func buildAuthString(clientid, secretid string) string {
  const prefix string = "Basic "
  clientinfo := fmt.Sprintf("%s:%s",clientid, secretid)
  clientinfoenc := b64.StdEncoding.EncodeToString([]byte(clientinfo))
  return fmt.Sprintf("%s%s", prefix, clientinfoenc)
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
