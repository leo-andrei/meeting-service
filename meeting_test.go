package meeting_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	meeting "github.com/leo-andrei/slackapi"
	"github.com/leo-andrei/slackapi/slack/mockslack"
)

func TestSendMeetingInfo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockslack.NewMockSlackClientInterface(ctrl)
	service := meeting.NewMeetingService(mockClient)

	testPayload := meeting.Payload{
		MeetingSummary: "Team meeting summary",
		Highlights:     []string{"Discussed Q4 goals", "Reviewed project timelines"},
		Username:       "test-user",
		Channel:        "test-channel",
	}

	jsonPayload, _ := json.Marshal(testPayload)

	req := httptest.NewRequest("POST", "/sendMeetingInfo", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockClient.EXPECT().PostMessage("test-channel", "test-user", gomock.Any()).Return(nil)

	service.SendMeetingInfo(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestSendMeetingInfo_InvalidPayload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockslack.NewMockSlackClientInterface(ctrl)
	service := meeting.NewMeetingService(mockClient)

	req := httptest.NewRequest("POST", "/sendMeetingInfo", bytes.NewBuffer([]byte("invalid payload")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	service.SendMeetingInfo(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestSendMeetingInfo_FailureToSendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockslack.NewMockSlackClientInterface(ctrl)
	service := meeting.NewMeetingService(mockClient)

	testPayload := meeting.Payload{
		MeetingSummary: "Team meeting summary",
		Highlights:     []string{"Discussed Q4 goals"},
		Username:       "test-user",
		Channel:        "test-channel",
	}

	jsonPayload, _ := json.Marshal(testPayload)

	req := httptest.NewRequest("POST", "/sendMeetingInfo", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockClient.EXPECT().PostMessage("test-channel", "test-user", gomock.Any()).Return(fmt.Errorf("mock failure"))

	service.SendMeetingInfo(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}
