package sns_go_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiUrl = "https://test-sns-api.seedao.tech"

func IsSafe(name string) bool {
	return checkReq(fmt.Sprintf("%s/v1/is_sensitive", apiUrl), name)
}

func Safe(name []string) []string {
	return replaceReq(fmt.Sprintf("%s/v1/sensitive", apiUrl), name)
}

func IsAvailable(name string) bool {
	return checkReq(fmt.Sprintf("%s/v1/is_unavailable", apiUrl), name)
}

func Available(name []string) []string {
	return replaceReq(fmt.Sprintf("%s/v1/unavailable", apiUrl), name)
}

func checkReq(url, word string) (result bool) {
	resp, err := http.Get(fmt.Sprintf("%s?word=%s", url, word))
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
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
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
