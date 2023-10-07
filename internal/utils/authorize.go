package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
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
// * => requestAccessToken

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
	state := base64.StdEncoding.EncodeToString(
		[]byte(strings.Join([]string{id, secret}, " ")))

	authorizeURL := fmt.Sprintf(
		"%s/oauth/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
		nibeURL,
		id,
		"READSYSTEM",
		"https://"+r.Host,
		state)
	fmt.Println(authorizeURL)
	http.Redirect(w, r, authorizeURL, http.StatusSeeOther)
}

// requestAccessToken requests access token with authorization code after
// authorization has completed
func requestAccessToken(w http.ResponseWriter, r *http.Request, code, state string) {

	fmt.Println("requestAccessToken")
	stateDecoded, err := base64.StdEncoding.DecodeString(state)
	if err != nil {
		renderPage(w, err.Error())
	}
	result := strings.Split(string(stateDecoded), " ")
	if len(result) != 2 {
		renderPage(w, "Unexpected format")
		return
	}
	id := result[0]
	secret := result[1]

	if state == "" || code == "" || id == "" || secret == "" {
		http.ServeFile(w, r, "templates/sign-up.html")
		return
	}
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", id)
	params.Add("client_secret", secret)
	params.Add("code", code)
	params.Add("redirect_uri", "https://"+r.Host)
	params.Add("scope", "READSYSTEM")

	resp, err := http.PostForm(fmt.Sprintf("%s/oauth/token", nibeURL), params)
	if err != nil {
		renderPage(w, err.Error())
		return
	}
	defer resp.Body.Close()
	out, err := os.Create("token.json")
	if err != nil {
		renderPage(w, err.Error())
	}

	defer out.Close()
	io.Copy(out, resp.Body)

	renderPage(w, "Application authorized. You may start collecting data.")
}

func renderPage(w http.ResponseWriter, extra string) {
	t, err := template.ParseFiles("templates/form.html")
	if err != nil {
		panic(err)
	}
	_ = t.ExecuteTemplate(w, "Error", "<p>\n"+extra+"\n</p>")
}
