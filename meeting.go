package meeting

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/leo-andrei/meeting-service/slack"
)

type Payload struct {
	MeetingSummary string   `json:"meetingSummary"`
	Highlights     []string `json:"highlights"`
	Username       string   `json:"username"`
	Channel        string   `json:"channel"`
}

// define a interface for meeting service
type MeetingService interface {
	SendMeetingInfo(w http.ResponseWriter, r *http.Request)
}

// define a struct for meeting service
type meetingSlackService struct {
	// define a slack client
	client slack.SlackClientInterface
}

// define a function to create a new meeting service
func NewMeetingService(client slack.SlackClientInterface) MeetingService {
	return &meetingSlackService{client: client}
}

// define a function to send meeting info
func (m *meetingSlackService) SendMeetingInfo(w http.ResponseWriter, r *http.Request) {
	// send meeting info to slack api
	// transform payload to slack message
	var payload Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	message := composeMessage(payload)
	err := m.client.PostMessage(payload.Channel, payload.Username, message)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Message sent successfully")
}

func composeMessage(payload Payload) string {
	if payload.MeetingSummary != "" && len(payload.Highlights) > 0 {
		return fmt.Sprintf("Summary: %s\nHighlights: %v", payload.MeetingSummary, payload.Highlights)
	} else if payload.MeetingSummary != "" {
		return fmt.Sprintf("Summary: %s", payload.MeetingSummary)
	} else if len(payload.Highlights) > 0 {
		return fmt.Sprintf("Highlights: %v", payload.Highlights)
	}
	return "No content available"
}
