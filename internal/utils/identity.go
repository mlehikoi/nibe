package utils

import (
	"encoding/base64"
	"errors"
	"strings"
)

type identity struct {
	ID     string
	Secret string
}

func (id identity) stringify() string {
	return base64.StdEncoding.EncodeToString(
		[]byte(strings.Join([]string{id.ID, id.Secret}, " ")))
}

func (id identity) String() string {
	return base64.StdEncoding.EncodeToString(
		[]byte(strings.Join([]string{id.ID, id.Secret}, " ")))
}

func parseIdentity(str string) (identity, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return identity{}, err
	}
	result := strings.Split(string(decoded), " ")
	if len(result) != 2 {
		return identity{}, errors.New("bad format")
	}
	return identity{result[0], result[1]}, nil
}
