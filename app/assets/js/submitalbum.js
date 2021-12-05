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


function getTokensIfRedirected() {
  const HOME = "http://localhost:8080/"
  const loc = window.location.href;
  if (loc != HOME){
    let params = new URLSearchParams(loc.split("?")[1]);
    history.pushState("","","/")
    let code = params.get('code');
    let state = params.get('state');
    setAccessTokens(code, state)
  }
}


// assign the access tokens to
function setAccessTokens(code, state) {
  const tokenurl = "http://localhost:8081/token";
  let tokendata = {
    Code: code,
    State: state
  };
  let body = JSON.stringify(tokendata);
  let request = new XMLHttpRequest();
  request.open('POST',tokenurl, true);
  request.setRequestHeader("Content-Type", "application/json");
  request.onreadystatechange=function() {
      if(request.readyState==4) {
        const token_resp = JSON.parse(JSON.parse(request.response));
        setTokenCookie(token_resp.access_token, token_resp.refresh_token, token_resp.expires_in);
      }
  }
  request.send(body);
}

function refreshAccessTokens() {
  const tokenurl = "http://localhost:8081/refreshtoken";
  let tokendata = {
    refresh_token: document.cookie.split('SpotifyRefreshToken=')[1].split(';')[0]
  };
  let body = JSON.stringify(tokendata);
  let request = new XMLHttpRequest();
  request.open('POST',tokenurl, true);
  request.setRequestHeader("Content-Type", "application/json");
  request.onreadystatechange=function() {
      if(request.readyState==4) {
        const token_resp = JSON.parse(JSON.parse(request.response));
        setTokenCookie(token_resp.access_token, token_resp.refresh_token, token_resp.expires_in);
      }
  }
  request.send(body);

}


function setTokenCookie(token, refresh_token, expires_in) {
  const d =  new Date();
  // cookie is retained for 24 hours
  d.setTime(d.getTime() + (expires_in * 1000));
  let expiry = " expires=" + d.toUTCString() +";";
  document.cookie = "SpotifyAccessToken="+token +";";
  document.cookie = expiry;
  document.cookie = "SpotifyRefreshToken="+refresh_token+";";
}
