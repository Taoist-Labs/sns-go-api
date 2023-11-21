package sns_go_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func IsSafe(name, safeHost string) bool {
	return checkReq(fmt.Sprintf("%s/v1/is_sensitive?word=%s", safeHost, name))
}

func Safe(name []string, safeHost string) []string {
	return replaceReq(fmt.Sprintf("%s/v1/sensitive", safeHost), name)
}

func IsAvailable(name, safeHost string) bool {
	return checkReq(fmt.Sprintf("%s/v1/is_unavailable?word=%s", safeHost, name))
}

func Available(name []string, safeHost string) []string {
	return replaceReq(fmt.Sprintf("%s/v1/unavailable", safeHost), name)
}

func checkReq(url string) (result bool) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	result = string(body) == "0"
	return
}

func replaceReq(url string, param []string) (result []string) {
	body, err := json.Marshal(param)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	return result
}
