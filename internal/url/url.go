package url

import (
	"net/url"
	"os"
	"strings"
)

func getTrimmedEnvironmentVariable(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

func FormatWithAuth(rawurl string) string {
	clientId := getTrimmedEnvironmentVariable("client_id")
	clientSecret := getTrimmedEnvironmentVariable("client_secret")
	u, _ := url.ParseRequestURI(rawurl)
	originalQuery := u.RawQuery

	queryMap, _ := url.ParseQuery(originalQuery)
	queryMap.Set("client_id", clientId)
	queryMap.Set("client_secret", clientSecret)
	newQuery := queryMap.Encode()

	if originalQuery != "" {
		return strings.Replace(rawurl, originalQuery, newQuery, 1)
	}

	return rawurl + "?" + newQuery
}
