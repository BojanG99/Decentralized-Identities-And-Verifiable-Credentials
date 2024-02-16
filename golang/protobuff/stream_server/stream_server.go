package main

import (
	"fmt"
	"log"
	"net"

	websocket_simulator "github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/protobuff/stream"

	"google.golang.org/grpc"
)

type server struct {
	websocket_simulator.UnimplementedWebSocketSimulatorServer
}

func (s *server) SendMessage(stream websocket_simulator.WebSocketSimulator_SendMessageServer) error {
	for {
		request, err := stream.Recv()
		if err != nil {
			break
		}

		response := &websocket_simulator.Response{
			Message: fmt.Sprintf("Server received: %s", request.Message),
		}

		if err := stream.Send(response); err != nil {
			break
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	websocket_simulator.RegisterWebSocketSimulatorServer(s, &server{}) // Use correct registration function

	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
