package utils

import (
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(url string) ([]byte, int, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, response.StatusCode, err
}

func GenerateURL(endpoint string, params map[string]string) string {
	for k, v := range params {
		endpoint += fmt.Sprintf("%s=%s&", k, v)
	}
	return endpoint
}
