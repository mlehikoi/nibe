package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"text/template"
)

const help = `Authorize the application

In order to be able to collect the heat pump data, the application has to be
authorized. First, expose this web server to the outside world.

ngrok http %s:%d

Then go to the forwarding address provided by ngrok with your web browser and
follow the instructions the air. The forwarding address looks like:
https://0123-4567-89ab-cdef-0123-4567-89ab-cdef-0123.ngrok.app
`

func Authorize(port int) {
	listener, err := net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		panic(err)
	}

	addr := listener.Addr().(*net.TCPAddr)
	fmt.Printf(help, addr.IP.String(), addr.Port)

	http.HandleFunc("/", serve)
	panic(http.Serve(listener, nil))
}

const nibeURL string = "https://api.nibeuplink.com"

// This function serves three types of requests:
// 1. the initial form for asking identifier and secret to start the process
// 2. process the form and redirect to NIBE authentication
// 3. the callback URL for NIBE to complete authorization
// The process goes as follows:
// * sign-up.html is served
// * user fills in identifier and secret and submits
// * => requestAuthorization
// * => redirected to api.nibeuplink.com
// * user authenticates
// * => requestAccessToken get called by Nibe Uplink

func serve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serve")
	r.ParseForm()
	id := r.FormValue("Identifier")
	secret := r.FormValue("Secret")
	if id != "" && secret != "" {
		requestAuthorization(w, r, id, secret)
		return
	}

	code := r.FormValue("code")
	state := r.FormValue("state")
	if code != "" && state != "" {
		fmt.Printf("request '%s' '%s'\n", code, state)
		requestAccessToken(w, r, code, state)
		return
	}

	fmt.Println("Render page as default")
	http.ServeFile(w, r, "templates/sign-up.html")
}

// requestAuthorization sends authorization request to NIBE Auth server
func requestAuthorization(w http.ResponseWriter, r *http.Request, id, secret string) {
	fmt.Println("requestAuthorization")

	authorizeURL := fmt.Sprintf(
		"%s/oauth/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
		nibeURL,
		id,
		"READSYSTEM",
		"https://"+r.Host,
		identity{id, secret})
	fmt.Println(authorizeURL)
	http.Redirect(w, r, authorizeURL, http.StatusSeeOther)
}

// requestAccessToken gets called by NIBE Uplink after authorization has
// completed. requestAccessToken will post a request back to the NIBE site from
// where it will get the final token.
func requestAccessToken(w http.ResponseWriter, r *http.Request, code, state string) {

	fmt.Println("requestAccessToken")
	ident, err := parseIdentity(state)
	if err != nil {
		renderPage(w, err.Error())
	}

	if state == "" || code == "" || ident.ID == "" || ident.Secret == "" {
		http.ServeFile(w, r, "templates/sign-up.html")
		return
	}
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", ident.ID)
	params.Add("client_secret", ident.Secret)
	params.Add("code", code)
	params.Add("redirect_uri", "https://"+r.Host)
	params.Add("scope", "READSYSTEM")

	resp, err := http.PostForm(fmt.Sprintf("%s/oauth/token", nibeURL), params)
	if err != nil {
		renderPage(w, err.Error())
		return
	}
	defer resp.Body.Close()
	var token Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		renderPage(w, err.Error())
		return
	}

	if err = token.dump("nibe-uplink-token.json"); err != nil {
		renderPage(w, err.Error())
	}

	renderPage(w, "Application authorized. You may start collecting data.\n"+
		"<pre>"+token.String()+"</pre>")
}

func renderPage(w http.ResponseWriter, extra string) {
	t, err := template.ParseFiles("templates/form.html")
	if err != nil {
		panic(err)
	}
	_ = t.ExecuteTemplate(w, "Error", "<p>\n"+extra+"\n</p>")
}
