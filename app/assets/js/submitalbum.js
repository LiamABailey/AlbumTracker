function asubmit() {
  console.log(getBody())
}

function getBody() {
  // check for all values populated
  if ((albumName.value != '') && (bandName.value != '') && (genre.value != '') && (year.value != '')) {
    postdata = {
      Name: albumName.value,
      Band: bandName.value,
      Genre: genre.value,
      Year: year.value
    };
    return JSON.stringify(postdata);
  } else {
    throw 'All input fields must be populated!';
  }
}
