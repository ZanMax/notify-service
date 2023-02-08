package services

import "github.com/slack-go/slack"

func sendMessageToSlack(text string, channelID string, token string) (bool, error) {
	slackClient := slack.New(token)
	/*
		attachment := slack.Attachment{
			Pretext: "Super Bot Message",
			Text:    "some text",
			// Color Styles the Text, making it possible to have like Warnings etc.
			Color: "#36a64f",
			// Fields are Optional extra data!
			Fields: []slack.AttachmentField{
				{
					Title: "Date",
					Value: time.Now().String(),
				},
			},
		}
		_, timestamp, err := client.PostMessage(
			channelID,
			slack.MsgOptionAttachments(attachment),
		)
	*/
	_, _, err := slackClient.PostMessage(channelID, slack.MsgOptionText(text, false))
	if err != nil {
		return false, err
	}
	return true, nil
}
