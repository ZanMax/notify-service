package services

import "github.com/slack-go/slack"

func sendMessageToSlack(text string, channelID string, token string) (bool, error) {
	slackClient := slack.New(token)
	_, _, err := slackClient.PostMessage(channelID, slack.MsgOptionText(text, false))
	if err != nil {
		return false, err
	}
	return true, nil
}
