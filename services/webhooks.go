package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func sendWebhookMessage(text string, url string, token string) (bool, error) {

	body, _ := json.Marshal(map[string]string{
		"text": text,
	})

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("token", token)

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	return true, nil
}
