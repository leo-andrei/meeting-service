package api

import (
	"log"
	"net/http"

	meeting "github.com/leo-andrei/meeting-service"
)

type Server struct {
	Port    string
	Service meeting.MeetingService
}

// define a function to create a new server
func NewServer(port string, service meeting.MeetingService) *Server {
	return &Server{Port: port, Service: service}
}

// define a function to start the server
func (s *Server) Start() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/send-meeting-info", s.Service.SendMeetingInfo)

	httpServer := &http.Server{
		Addr:    ":" + s.Port,
		Handler: mux,
	}

	go func() {
		log.Printf("Server starting on port %s", s.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	return httpServer
}
