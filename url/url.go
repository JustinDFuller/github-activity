package url

import (
	"os"
	"strings"
)

func getTrimmedEnvironmentVariable(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

func FormatWithAuth(url string) string {
	clientId := getTrimmedEnvironmentVariable("client_id")
	clientSecret := getTrimmedEnvironmentVariable("client_secret")
	return url + "?client_id=" + clientId + "&client_secret=" + clientSecret
}
