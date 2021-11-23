module albumtrack

go 1.16

replace github.com/LiamABailey/AlbumTracker/pkg/albumdata => ./pkg/albumdata

replace github.com/LiamABailey/AlbumTracker/pkg/spotifyintegration => ./pkg/spotifyintegration

require (
	github.com/LiamABailey/AlbumTracker/pkg/albumdata v0.0.0-00010101000000-000000000000
	github.com/LiamABailey/AlbumTracker/pkg/spotifyintegration v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
)
