package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", Serve)
	http.ListenAndServe(":8080", nil)
}

type Identity struct {
	ID     string
	Secret string
}

var States = map[string]Identity{}

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func RedirectToNIBE(w http.ResponseWriter, r *http.Request, id string, secret string) {
	state := uuid.New().String()
	States[state] = Identity{id, secret}
	authorizeURL := fmt.Sprintf(
		"https://api.nibeuplink.com/oauth/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
		id,
		"READSYSTEM",
		"https://"+r.Host,
		state)
	fmt.Println(authorizeURL)
	http.Redirect(w, r, authorizeURL, http.StatusSeeOther)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parsing request")

	r.ParseForm()
	id := r.FormValue("Identifier")
	secret := r.FormValue("Secret")
	if id != "" && secret != "" {
		RedirectToNIBE(w, r, id, secret)
		return
	}
	code := r.FormValue("code")
	state := r.FormValue("state")
	id = States[state].ID
	secret = States[state].Secret
	if state == "" || code == "" || id == "" || secret == "" {
		http.ServeFile(w, r, "./SignUp.html")
		return
	}

	fmt.Println(state)
	fmt.Println(code)
	fmt.Println("state:" + state + "\ncode:" + code)
	fmt.Fprintln(w, code)

	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", id)
	params.Add("client_secret", secret)
	params.Add("code", code)
	params.Add("redirect_uri", "https://"+r.Host)
	params.Add("scope", "READSYSTEM")

	resp, err := http.PostForm("https://api.nibeuplink.com/oauth/token", params)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Log the response body
	bodyString := string(body)
	log.Print("resp: " + bodyString)

	// Unmarshal result
	token := Token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	log.Printf("Post added with ID %s", token.AccessToken)

	// reqURL := fmt.Sprintf("https://api.nibeuplink.com/oauth/token")
	// req, err := http.NewRequest(http.MethodPost, reqURL,
	// 	strings.NewReader(form.Encode()))
	// if err != nil {
	// 	fmt.Fprintf(w, "could not create HTTP request: %v", err)
	// 	return
	// }
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// // We set this header since we want the response
	// // as JSON
	// req.Header.Set("accept", "application/json")

	// return

	// // Send out the HTTP request
	// httpClient := http.Client{}
	// res, err := httpClient.Do(req)
	// if err != nil {
	// 	fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
	// defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	//var t OAuthAccessResponse
	// if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
	// 	fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	// Finally, send a response to redirect the user to the "welcome" page
	// with the access token
	// w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
	// w.WriteHeader(http.StatusFound)

	// code := r.URL.Query()["code"][0]
	// state := r.URL.Query()["state"][0]
	// fmt.Fprintln(w, code)
	// fmt.Fprintln(w, state)

	// // We will be using `httpClient` to make external HTTP requests later in our code
	// httpClient := http.Client{}

	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Fprint(w, "Post\n")
	// fmt.Fprint(w, string(body))
	// fmt.Fprint(w, "\nPost printed\n")

}
