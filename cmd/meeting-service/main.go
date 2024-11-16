package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	meeting "github.com/leo-andrei/slackapi"
	"github.com/leo-andrei/slackapi/api"
	"github.com/leo-andrei/slackapi/slack"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	token := os.Getenv("SLACK_TOKEN")
	slackClient := slack.NewSlackClient(token)
	slackService := meeting.NewMeetingService(slackClient)
	server := api.NewServer(port, slackService)
	httpServer := server.Start()

	// Set up channel to listen for OS signals for graceful shutdown.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Wait for an interrupt signal
	<-signalChan
	log.Println("Received shutdown signal, shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	} else {
		log.Println("Server shutdown completed gracefully.")
	}
}
