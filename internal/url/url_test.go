package url

import (
	"os"
	"testing"
)

func TestFormatWithAuth(t *testing.T) {
	expected := "https://github.com/user/me?client_id=client_id_with_spaces&client_secret=client_secret_with_spaces"

	os.Setenv("client_id", "     client_id_with_spaces   ")
	os.Setenv("client_secret", "     client_secret_with_spaces   ")

	url := FormatWithAuth("https://github.com/user/me")

	if url != expected {
		t.Error("Expected URL: " + expected + "But received: " + url)
	}
}

func TestFormatWithAuthWithExistingQuery(t *testing.T) {
	expected := "https://github.com/user/me?client_id=client_id_with_spaces&client_secret=client_secret_with_spaces&test=testValue"

	os.Setenv("client_id", "     client_id_with_spaces   ")
	os.Setenv("client_secret", "     client_secret_with_spaces   ")

	url := FormatWithAuth("https://github.com/user/me?test=testValue")

	if url != expected {
		t.Error("Expected URL: " + expected + "But received: " + url)
	}
}
