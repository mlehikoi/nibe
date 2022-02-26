package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
func (oauth *Oauth) Refresh(uri string) ([]byte, error) {
	params := url.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("client_id", oauth.ID)
	params.Add("client_secret", oauth.Secret)
	params.Add("refresh_token", oauth.RefreshToken)
	resp, err := http.PostForm(uri, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("%s returned %d", uri, resp.StatusCode)
	}

	token := Token{}
	if err = json.Unmarshal(body, &token); err != nil {
		return nil, err
	}
	oauth.AccessToken = token.AccessToken
	oauth.RefreshToken = token.RefreshToken
	return body, nil
}
