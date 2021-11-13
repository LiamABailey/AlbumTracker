import (
  "crypto/rand"
  "encoding/hex"
  "github.com/gin-gonic/gin"
  "net/http"
  "os"
)

// define relevant system environment variable names
const (
  CLIENTID = "CLIENTID"
  REDIRECTURI = "REDIRECTURI"
)


type SpotifyServer struct {
  router *gin.engine
}

func NewSpotifyServer() (*SpotifyServer, error) {
  svr := &SpotifyServer{}
  svr.router = gin.Default()
  svr.router.GET("/login", login)
  // TODO svr.router.GET("/recently-played", getRecentlyPlayed)
}

func (srv *SpotifyServer) Run(address string) error {
  return srv.rounter.Run(address)
}

func (srv *SpotifyServer) Use(hf gin.HandlerFunc) {
  svr.router.Use(hf)
}

func login(ctx *gin.Context) {
  const responsetype string = "code"
  const scopes string = "user-read-recently-played"
  const pathprefix string = "https://accounts.spotify.com/authorize?"
  var state, stateerr := generateRandomState()
  // panic if we're unable to correctly create state
  if staterr != nil {
    panic(staterr)
  }
  // todo : submit the login request and redirect

}


// generates a random hex-string (length 16)
// for saftey in the redirect
func generateRandomState() (string, error) {
  const nbytes int = 32
  b32 := make([]byte, nbytes)
  _, err := rand.Read(b32)
  if err != nil {
    return nil, err
  }
  return hex.EncodeToString(b32), nil
}
