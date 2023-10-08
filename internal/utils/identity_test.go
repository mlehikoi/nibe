package utils

import (
	"encoding/base64"
	"testing"
)

func TestStringify(t *testing.T) {
	id := identity{"id", "secret"}
	if id.stringify() != "aWQgc2VjcmV0" {
		t.Fatal("Unexpected stringify: " + id.stringify())
	}
}

func TestParseSuccess(t *testing.T) {
	id, err := parseIdentity("aWQgc2VjcmV0")
	if err != nil {
		t.Fatal("Unexpected parse error")
	}
	if id.ID != "id" || id.Secret != "secret" {
		t.Fatalf("Unexpected parse error: %s:%s", id.ID, id.Secret)
	}
}

func TestParseBadEncoding(t *testing.T) {
	_, err := parseIdentity("äWQgc2VjcmV0")
	if err == nil {
		t.Fatal("Expected an error")
	}
}

func TestParseBadData(t *testing.T) {
	encoded := base64.StdEncoding.EncodeToString([]byte("id_secret"))
	_, err := parseIdentity(encoded)
	if err == nil {
		t.Fatal("Expected an error")
	}
}
