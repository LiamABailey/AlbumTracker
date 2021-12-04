function asubmit() {
  const posturl = "http://localhost:8080/albums";
  let body = getBody();
  request = new XMLHttpRequest();
  request.open('POST',posturl, true);
  request.send(body);
  //reset the form
  document.getElementById("input-form").reset()
}

function getBody() {
  // check for all values populated
  if ((albumName.value != '') && (bandName.value != '') && (genre.value != '') && (year.value != '')) {
    let postdata = {
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

function acquireAccessTokens() {
  const HOME = "http://localhost:8080/"
  const loc = window.location.href;
  if (loc != HOME){
    let params = new URLSearchParams(loc.split("?")[1]);
    history.pushState("","","/")
    const tokenurl = "http://localhost:8081/token";
    let code = params.get('code');
    let state = params.get('state');
    let tokendata = {
      Code: code,
      State: state
    };
    let body = JSON.stringify(tokendata);
    let request = new XMLHttpRequest();
    request.open('POST',tokenurl, true);
    request.setRequestHeader("Content-Type", "application/json");
    request.send(body);
    console.log(request.response)
  }

}

function setAuthCookie(authcode, authstate) {
  const d =  new Date();
  // cookie is retained for 24 hours
  d.setTime(d.getTime() + (24 * 60 * 60 * 1000));
  let expiry = "expires=" + d.toUTCString();
  document.cookie = "SpotifyAuthCode="+authcode+"; "+expiry+"; path=/";
  document.cookie = "SpotifyAuthState="+authstate+"; "+expiry+"; path=/";
}
