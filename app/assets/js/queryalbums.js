function asearch() {
  const posturlstart = "http://localhost:8080/albums/search";
  var query = getSearchQuery();
  var request = new XMLHttpRequest();
  var qstr = [posturlstart, query].join('')
  console.log(qstr)
  request.open('GET',qstr, true);
  request.onreadystatechange=function() {
      if(request.readyState==4) {
        let albums = JSON.parse(request.response);
        if (albums != null) {
          populateSearchResultTable(albums);
        }
      }
  }
  clearSearchResultTable();
  request.send({});
}

// clears the results in the table
function clearSearchResultTable() {
  var table = document.getElementById("result-table-body");
  while (table.rows.length > 0) {
    table.deleteRow(0);
  }
}

// populates the results in the table
// given an array of json response data
function populateSearchResultTable(albums) {
  var name_order = ["Name","Band","Genre","Year"]
  var table = document.getElementById("result-table-body");
  for (const albumJson of albums) {
    let row = table.insertRow();
    for (const name of name_order) {
      let cell = row.insertCell();
      let text = document.createTextNode(albumJson[name]);
      cell.appendChild(text);
    }
  }
}

// Returns the query string, starting with "?"
function getSearchQuery() {
  var qstr = [];
  var params = {'AlbumName': albumName.value,
                'AlbumNameExactMatch': albumNameExactMatch.checked,
                'BandName': bandName.value,
                'BandNameExactMatch': bandNameExactMatch.checked,
                'Genre': getSelectedGenres(genre),
                'YearStart': yearStart.value,
                'YearEnd': yearEnd.value,
                'DateAddedStart': entryStart.value,
                'DateAddedEnd': entryEnd.value};
  // TODO hanlde exact match indicators
  // for all except for genre
  Object.entries(params).map(([key, val]) => {
    if (key != 'Genre' && key != 'AlbumNameExactMatch' && key != 'BandNameExactMatch') {
      if (val != '') {
        qstr.push([key,'=',val.replaceAll(" ", "%20")].join(''),'&')
      }
    } else if (key == 'AlbumNameExactMatch' || key == 'BandNameExactMatch') {
      qstr.push([key,'=',val].join(''),'&')
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
