package spotifyintegration

import (
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	gincors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"
)

// define relevant system environment variable names
const (
	CLIENTID    = "CLIENTID"
	SECRETID    = "SECRETID"
	REDIRECTURI = "REDIRECTURI"
)

type SpotifyServer struct {
	router    *gin.Engine
	laststate string
}

func NewSpotifyServer() *SpotifyServer {
	svr := &SpotifyServer{}
	svr.router = gin.Default()
	corsspec := gincors.New(gincors.Config{
		AllowOrigins: []string{"http://localhost:8080", "https://localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Authorization", "content-type"},
	})
	svr.router.Use(corsspec)
	svr.router.GET("/lastalbums", svr.getLastAlbums)
	svr.router.GET("/login", svr.login)
	svr.router.POST("/token", svr.requestTokens)
	svr.router.POST("/refreshtoken", svr.refreshTokens)
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
	authlocation := url.URL{Path: pathprefix, RawQuery: authquery.Encode()}
	ctx.Redirect(http.StatusFound, authlocation.RequestURI())
}

func (svr *SpotifyServer) requestTokens(ctx *gin.Context) {
	const granttype string = "authorization_code"
	const contenttype string = "application/x-www-form-urlencoded"
	const pathprefix string = "https://accounts.spotify.com/api/token"

	var request SpotifyRequestTokensRequestBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// if returned state != state at login
	if request.State != svr.laststate {
		emessage := "Request State does not match Login State"
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New(emessage)))
		return
	}

	// build the query
	authquery := url.Values{}
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
	ctx.IndentedJSON(http.StatusOK, string(body))
}

func (svr *SpotifyServer) refreshTokens(ctx *gin.Context) {
	const granttype string = "refresh_token"
	const contenttype string = "application/x-www-form-urlencoded"
	const pathprefix string = "https://accounts.spotify.com/api/token"

	var request SpotifyRefreshTokensRequestBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// build the query
	authquery := url.Values{}
	authquery.Set("grant_type", granttype)
	authquery.Set("refresh_token", request.RefreshToken)
	authlocation := url.URL{Path: pathprefix, RawQuery: authquery.Encode()}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", authlocation.RequestURI(), nil)
	// set the required headers
	req.Header.Set("Content-Type", contenttype)
	authb64 := buildAuthString(os.Getenv(CLIENTID), os.Getenv(SECRETID))
	req.Header.Set("Authorization", authb64)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	ctx.IndentedJSON(http.StatusOK, string(body))
}

func (svr *SpotifyServer) getLastAlbums(ctx *gin.Context) {
	const granttype string = "authorization_code"
	const contenttype string = "application/x-www-form-urlencoded"
	const pathprefix string = "https://api.spotify.com/v1/me/player/recently-played"

	// make repeated calls to the endpoint until we have
	// filled our map
	albums := make(map[string]SpotifyAlbumTrackInfo)
	var iter int
	var request SpotifyGetAlbumsRequestQuery
	// make multiple requests until 10 albums have been gathered
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	before := strconv.Itoa(request.Before)
	lim, _ := strconv.Atoi(request.Limit)
	nonemptyresponse := true
	for (len(albums) < lim) && (iter < 5) && nonemptyresponse {
		// build the query
		loc := pathprefix + "?before=" + before
		client := &http.Client{}
		req, _ := http.NewRequest("GET", loc, nil)
		// set the required headers
		req.Header.Set("Content-Type", contenttype)
		req.Header.Set("Authorization", ctx.Request.Header["Authorization"][0])
		fmt.Println(loc)
		// submit the request
		resp, _ := client.Do(req)
		// collect the body data
		body, _ := io.ReadAll(resp.Body)
		var albumdata SpotifyRecentlyPlayedBody
		json.Unmarshal(body, &albumdata)
		playedalbums := albumdata.GetUniqueAlbums()
		if len(playedalbums) == 0 {
			nonemptyresponse = false
		} else {
			// collect the results
			for id, album := range playedalbums {
				albums[id] = album
			}
		}
		fmt.Println(playedalbums)
		before = albumdata.Cursors.Before
		iter++
		resp.Body.Close()
		// wait 1 second before next request
		time.Sleep(1 * time.Second)
	}
	// TODO trim to length if above lim.
	// as validation, reutrn the unmarshaled data
	entries := make([]SpotifyAlbumTrackInfo, 0, len(albums))
	for _, album := range albums {
		entries = append(entries, album)
	}
	// sort by album name
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})
	marshalledalbums, _ := json.Marshal(entries)
	ctx.IndentedJSON(http.StatusOK, string(marshalledalbums))
}

// format the authorization string
func buildAuthString(clientid, secretid string) string {
	const prefix string = "Basic "
	clientinfo := fmt.Sprintf("%s:%s", clientid, secretid)
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
