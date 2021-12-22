# AlbumTracker
#### Simple Web Interface for tracking Albums
#### Forked from LiamABailey/LearnDocker

I listen to a variety of albums on a lot of different platforms (Spotify, YouTube, BandCamp, Soundcloud, and physical media), and often struggle to keep track of what I've listened to, especially when I can remember *aspects* of a song or album, but not the title. Consequently, I'm building an album tracker, leveraging a manual web interface to input information, and presenting the data in a query-able manner for later search on a separate web interface.

To ease data entry, a Spotify integration is being built to present data on recently listened albums (leveraging the song-level `recently_played` endpoint & aggregating to the album level). Given that 75%+ of my music listening happens on Spotify, this represents a dramatic efficiency - I don't have to remember what I was listening to last, how to spell an artist or album name, or look up the release year.


## Internals
### Go-based API
Core API functionality is managed via Go, which supports POSTing a new album, DELETE-ing an existing album by ID, and GETting search results (for display). Search supports variable parameters and a limit on returned results. Additionally, Spotify integration (login + recently played albums) is managed. As of 10/02/2021, PATCH is being considered to support functionality such as rating an album (and changing the rating), or updating a last-listened-to date.

### Mongo backend
Data is stored via MongoDB, and is persisted across container lifetimes.

## Externals
A simple UI, built in HTML/CSS, supports submitting new records & querying existing records. JavaScript is used to interact with the API and manage some authorization components.

## TODO
- Clean up UI
- Recently played:
  - Button to move row of album data into form
  - Expose existing functionality to get albums played before a given date
  - Sort by `played_at` (currently by album name)
- Improvement to genre system (Read from DB, allow user to extend)
