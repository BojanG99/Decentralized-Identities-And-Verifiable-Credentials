package main

import (
	"context"
	"fmt"
	"log"

	verifiable_credentials "github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/protobuff/issuing_credentials"
	"google.golang.org/grpc"
)

func main() {
	// Replace with the actual gRPC server address
	serverAddress := "localhost:50051"

	// Dial the gRPC server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := verifiable_credentials.NewVerifiableCredentialServiceClient(conn)

	// Simulate the stream of client messages
	clientStream, err := client.IssueCredential(context.Background())
	if err != nil {
		log.Fatalf("Error creating client stream: %v", err)
	}

	// Simulate sending a ClientHello message
	clientHello := &verifiable_credentials.ClientMessage{
		Message: &verifiable_credentials.ClientMessage_Ch{
			Ch: &verifiable_credentials.ClientHello{
				DidString:       "bojan did example_did",
				ServerChallange: "example_server_challenge",
				RpcKey:          "example_rpc_key",
			},
		},
	}

	if err := clientStream.Send(clientHello); err != nil {
		log.Fatalf("Error sending ClientHello message: %v", err)
	}

	// Receive and process server messages
	for {
		serverMessage, err := clientStream.Recv()
		if err != nil {
			log.Fatalf("Error receiving server message: %v", err)
		}

		switch m := serverMessage.Message.(type) {
		case *verifiable_credentials.ServerMessage_Sh:
			fmt.Println("Received ServerHello message:", m.Sh)
			clientAuthentication := &verifiable_credentials.ClientMessage{
				Message: &verifiable_credentials.ClientMessage_Ca{
					Ca: &verifiable_credentials.ClientAuthentication{
						CompletedChallange: "Test",
						KeyUrl:             "bojan key url",
						Pin:                "12345",
					},
				},
			}
			if err := clientStream.Send(clientAuthentication); err != nil {
				log.Fatalf("Error sending ClientHello message: %v", err)
			}
			// Process ServerHello message
		case *verifiable_credentials.ServerMessage_Sp:
			fmt.Println("Received ServerClaimsProposal message:", m.Sp)
			clientResponse := &verifiable_credentials.ClientMessage{
				Message: &verifiable_credentials.ClientMessage_Cr{
					Cr: &verifiable_credentials.ClientResponse{
						AcceptCredentials: true,
					},
				},
			}
			if err := clientStream.Send(clientResponse); err != nil {
				log.Fatalf("Error sending ClientHello message: %v", err)
			}
			// Process ServerClaimsProposal message
		case *verifiable_credentials.ServerMessage_Vc:
			fmt.Println("Received VerifiableCredential message:", m.Vc)
			clientStream.CloseSend()
			return
			// Process VerifiableCredential message
		case *verifiable_credentials.ServerMessage_Err:
			fmt.Println("Received Error message:", m.Err)
			// Process Error message
		default:
			log.Println("Unknown message type received")
		}
	}
}
