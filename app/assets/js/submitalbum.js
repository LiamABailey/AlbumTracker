function asubmit() {
  const posturl = "http://localhost:8080/albums";
  body = getBody();
  console.log(getBody());
  request = new XMLHttpRequest();
  request.open('POST',posturl, true);
  request.send(body);
  //reset the form
  document.getElementById("input-form").reset()
}

function getBody() {
  // check for all values populated
  if ((albumName.value != '') && (bandName.value != '') && (genre.value != '') && (year.value != '')) {
    postdata = {
      Name: albumName.value,
      Band: bandName.value,
      Genre: genre.value,
      Year: parseInt(year.value)
    };
    return JSON.stringify(postdata);
  } else {
    throw 'All input fields must be populated!';
  }
}
