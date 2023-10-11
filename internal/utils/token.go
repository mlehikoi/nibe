package utils

import (
	"encoding/json"
	"os"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (token *Token) String() string {
	jsonData, err := json.MarshalIndent(token, "", "    ")
	if err != nil {
		return "{}"
	}
	return string(jsonData)
}

func (token *Token) Dump(filename string) error {
	jsonData, err := json.Marshal(token)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	return err
}
