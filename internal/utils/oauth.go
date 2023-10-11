package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Oauth struct {
	AccessToken  string
	ID           string
	Secret       string
	RefreshToken string
}

func (oauth Oauth) Get(uri string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet,
		uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+oauth.AccessToken)
	return http.DefaultClient.Do(req)
}

// Refresh tries to refresh the access token. Returns servers response
func (oauth *Oauth) Refresh(uri string) (Token, error) {
	params := url.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("client_id", oauth.ID)
	params.Add("client_secret", oauth.Secret)
	params.Add("refresh_token", oauth.RefreshToken)
	resp, err := http.PostForm(uri, params)

	fmt.Println("Connect to " + uri)
	if err != nil {
		fmt.Println("Error:", err)
		return Token{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad status", resp.StatusCode)
		return Token{}, fmt.Errorf("%s returned %d", uri, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return Token{}, err
	}

	token := Token{}
	if err = json.Unmarshal(body, &token); err != nil {
		fmt.Println("Error:", err)
		return Token{}, err
	}
	fmt.Println("Great success!")
	oauth.AccessToken = token.AccessToken
	oauth.RefreshToken = token.RefreshToken
	return token, nil
}
