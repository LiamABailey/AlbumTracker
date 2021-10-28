function asearch() {
  const posturlstart = "http://localhost:8080/albums/search";
  var query = getSearchQuery();
  var request = new XMLHttpRequest();
  var qstr = [posturlstart, query].join('')
  console.log(qstr)
  request.open('GET',qstr, true);
  request.onreadystatechange=function() {
      if(request.readyState==4) {
        console.log(request.response)
      }
  }
  request.send({});

}

// Returns the query string, starting with "?"
function getSearchQuery() {
  var qstr = [];
  var params = {'Name': albumName.value,
                'Band': bandName.value,
                'Genre': getSelectedGenres(genre),
                'YearStart': yearStart.value,
                'YearEnd': yearEnd.value,
                'DateAddedStart': entryStart.value,
                'DateAddedEnd': entryEnd.value};
  // TODO hanlde exact match indicators
  // for all except for genre
  Object.entries(params).map(([key, val]) => {
    if (key != 'Genre') {
      if (val != '') {
        qstr.push([key,'=',val.replaceAll(" ", "%20")].join(''),'&')
      }
    } else {
      for (var i=0; i < val.length; i++) {
        qstr.push(['Genres','=',val[i].replaceAll(" ", "%20")].join(''),'&')
      }
    }
  });
  // will end with an '&', so we pop
  qstr.pop();
  // if not empty, start with ?
  if (qstr != []){
    qstr.unshift('?');
  }
  return qstr.join('');
}

// Get all genres selected in the genre picker
function getSelectedGenres(genres) {
  var selections = [];
  var options = genres && genres.options;
  var genre;
  for (var i=0; i < options.length; i++) {
    genre = options[i]
    if (genre.selected) {
      selections.push(genre.value)
    }
  }
  return selections
}
