# AlbumTracker - WORK IN PROGRESS
#### Simple Web Interface for tracking Albums
#### Forked from LiamABailey/LearnDocker

I listen to a variety of albums on a lot of different platforms (Spotify, YouTube, Bandcamp, Soundcloud, and physical media), and often struggle to keep track of what I've listened to, especially when I can remember *aspects* of a song or album, but not the title. Consequently, I'm building a last.fm-inspired music (album) tracker, leveraging a manual web interface to input information, and presenting the data in a query-able manner for later search on a separate web interface.

I'm starting with the underlying API, which will then be rigged up to a JavaScript-based web app to receive and display data.

## Internals
### Go-based API
Core API functionality is managed via Go, which supports POSTing a new album, DELETE-ing an existing album by ID, and GETting search results (for display). Search supports variable parameters and a limit on returned results. As of 10/02/2021, PATCH is being considered to support functionality such as rating an album (and changing the rating), or updating a last-listened-to date.

### Mongo backend
Data is stored via MongoDB, and is persisted across container lifetimes.


## Externals
Currently researching a JavaScript (or TypeScript?)/HTML/CSS solution for building a front-end.
