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
        setCookie("SpotifyAccessToken",token_resp.access_token, token_resp.expires_in);
        setCookie("SpotifyRefreshToken", token_resp.refresh_token, null);
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
        setCookie("SpotifyAccessToken",token_resp.access_token, token_resp.expires_in);
      }
  }
  request.send(body);
}

// set a single cookie given name, value, and optional expiration time
// (in seconds in the future)
function setCookie(cookie_name, cookie_value, expires_in) {
  let cookie_str = cookie_name+"="+cookie_value+";";
  if (expires_in != null) {
    cookie_str += "max-age=" + expires_in +";";
  }
  document.cookie = cookie_str
}
