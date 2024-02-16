package main

import (
	"context"
	"fmt"
	"log"

	websocket_simulator "github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/protobuff/stream"

	"google.golang.org/grpc"
)

func runClient(client websocket_simulator.WebSocketSimulatorClient) {
	stream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	messages := []string{"Hello", "How are you?", "Goodbye"}
	for _, message := range messages {
		request := &websocket_simulator.Request{
			Message: message,
		}

		if err := stream.Send(request); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}

		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}
		fmt.Printf("Send message: %v \n", message)
		fmt.Printf("Received response: %v\n", response.Message)
	}

	stream.CloseSend()
}

func main() {
	conn, err := grpc.Dial("0.tcp.eu.ngrok.io:13552", grpc.WithInsecure())
	//	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := websocket_simulator.NewWebSocketSimulatorClient(conn)
	runClient(client)
}
