package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendTelegramMessage(text string, chatId []string, token string) (bool, error) {
	var err error
	var response *http.Response
	telegramAPIURL := fmt.Sprintf("https://api.telegram.org/bot%s", token)
	defer response.Body.Close()

	for _, id := range chatId {

		url := fmt.Sprintf("%s/sendMessage", telegramAPIURL)
		body, _ := json.Marshal(map[string]string{
			"chat_id": id,
			"text":    text,
		})
		response, err = http.Post(
			url,
			"application/json",
			bytes.NewBuffer(body),
		)
		if err != nil {
			return false, err
		}

		body, err = io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
