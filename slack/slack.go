package slack

import (
	"log"
)

type SlackClientInterface interface {
	PostMessage(channelID string, username string, message string) error
}

type SlackClient struct {
	Token string
}

func NewSlackClient(token string) *SlackClient {
	return &SlackClient{Token: token}
}

func (s *SlackClient) PostMessage(channelID string, username string, message string) error {
	// if channelID is set, send to channel
	// if username is set, send to user
	if channelID == "" && username == "" {
		log.Printf("Cannot post message to Slack without a channelID or username")
		return nil
	}
	if channelID == "" {
		log.Printf("Posting message to Slack to user: %s, message: %v", username, message)
		return nil
	}
	if username == "" {
		log.Printf("Posting message to Slack to channelID: %s, message: %v", channelID, message)
		return nil
	}
	return nil

	// the logic here should use the slack-go/slack package from github.com/slack-go/slack
	// client := slack.New(s.Token)
	// channelID := channel // Replace with the desired channel or user ID

	// _, _, err := client.PostMessage(channelID, slack.MsgOptionText(message, false))
	// if err != nil {
	//     log.Printf("Failed to send message to Slack: %v", err)
	//     return err
	// }
}
