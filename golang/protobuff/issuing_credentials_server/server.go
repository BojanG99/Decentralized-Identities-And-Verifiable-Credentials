package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	kripto "github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/crypto"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/protobuff/helpers"
	verifiable_credentials "github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/protobuff/issuing_credentials"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/resolver"
	"github.com/lestrrat-go/jwx/jwk"
	"google.golang.org/grpc"
)

type verifiableCredentialServer struct {
	verifiable_credentials.VerifiableCredentialServiceServer
}

type RpcRequestData struct {
	rpcKey          string
	clientChallange string
	vcData          *VCData
	clientDidString string
	serverChallange string
}

func (s *verifiableCredentialServer) IssueCredential(stream verifiable_credentials.VerifiableCredentialService_IssueCredentialServer) error {

	var grpcRequestData RpcRequestData
	//Step1
	clientMessage, err := stream.Recv()
	if err != nil {
		log.Printf("Error receiving client message: %v", err)
		return err
	}

	switch m := clientMessage.Message.(type) {
	case *verifiable_credentials.ClientMessage_Ch:
		err := handleClientHello(stream, m.Ch, &grpcRequestData)
		if err != nil {
			return err
		}
	case *verifiable_credentials.ClientMessage_Err:
		log.Printf("Unsuported client error message: %v\n", m.Err.ErrorMessage)
		return errors.New(m.Err.ErrorMessage)
	default:
		log.Printf("Unsuported server message type: %v\n", m)
		return errors.New("Unsuported client message. 1. korak")
	}

	//Step 2
	clientMessage, err = stream.Recv()
	if err != nil {
		log.Printf("Error receiving client message: %v", err)
		return err
	}

	switch m := clientMessage.Message.(type) {
	case *verifiable_credentials.ClientMessage_Ca:
		//err :=
		handleClientAuthentication(stream, m.Ca, &grpcRequestData)
		//if err != nil {
		//	return err
		//}
	case *verifiable_credentials.ClientMessage_Err:
		log.Printf("Unsuported client error message: %v\n", m.Err.ErrorMessage)
		return errors.New(m.Err.ErrorMessage)
	default:
		log.Printf("Unsuported server message type: %v\n", m)
		return errors.New("Unsuported client message. 2 korak")
	}

	//step3
	clientMessage, err = stream.Recv()
	if err != nil {
		log.Printf("Error receiving client message: %v", err)
		return err
	}
	switch m := clientMessage.Message.(type) {
	case *verifiable_credentials.ClientMessage_Cr:
		handleClientResponse(stream, m.Cr, &grpcRequestData)
	case *verifiable_credentials.ClientMessage_Err:
		log.Printf("Unsuported client error message: %v\n", m.Err.ErrorMessage)
		return errors.New(m.Err.ErrorMessage)
	default:
		log.Printf("Unsuported server message type: %v\n", m)
		return errors.New("Unsuported client message. 3. korak")
	}

	return nil
}

func sendErrorMessage(errorMessage string, stream verifiable_credentials.VerifiableCredentialService_IssueCredentialServer) error {
	serverMessage := &verifiable_credentials.ServerMessage{
		Message: &verifiable_credentials.ServerMessage_Err{
			Err: &verifiable_credentials.Error{
				ErrorMessage: errorMessage,
			},
		},
	}

	if err := stream.Send(serverMessage); err != nil {
		log.Printf("Error sending server message: %v", err)
		return err
	}

	return nil

}

func handleClientHello(stream verifiable_credentials.VerifiableCredentialService_IssueCredentialServer, clientHello *verifiable_credentials.ClientHello, grpcRequestData *RpcRequestData) error {

	didString := clientHello.GetDidString()
	rpcKey := clientHello.GetRpcKey()
	serverChallange := clientHello.GetServerChallange()

	if rpcKey == "" || serverChallange == "" || didString == "" {
		sendErrorMessage("Invalid data. You need to send all ClientHello fields", stream)
		return errors.New("Missig fields")
	}

	vcData, err := read_vc_data(rpcKey)
	pKeys := read_keys()
	ethipfs_resolver := resolver.MakeEthIpfsResolver()
	issuer_diddocument, err := ethipfs_resolver.ResolveDID(vcData.Issuer)

	keyUrl := issuer_diddocument.Authentication[0]

	key, err := ethipfs_resolver.ResolveDIDKey(keyUrl)
	var signKey jwk.Key

	for _, data := range pKeys {
		keyId, ok := data["id"]
		if !ok {
			continue
		}

		if keyId == key.ID {
			jsonString, _ := json.Marshal(data)

			signKey, _ = jwk.ParseKey([]byte(jsonString))
			break
		}
	}

	completed_challange, err := kripto.SignMessage(serverChallange, signKey)

	if err != nil {
		sendErrorMessage("Invalid rpc key", stream)
		return err
	}

	if rpcKey == "" {
		sendErrorMessage("Invalid rpcKey", stream)
	}
	grpcRequestData.vcData = vcData
	grpcRequestData.rpcKey = rpcKey
	grpcRequestData.clientDidString = didString
	grpcRequestData.serverChallange = serverChallange
	clientChallenge := helpers.GenerateRandomString(10)

	grpcRequestData.clientChallange = clientChallenge

	fmt.Println(*clientHello)
	serverMessage := &verifiable_credentials.ServerMessage{
		Message: &verifiable_credentials.ServerMessage_Sh{
			Sh: &verifiable_credentials.ServerHello{
				// Set appropriate fields based on your logic
				DidString:          vcData.Issuer,
				ClientChallange:    clientChallenge,
				CompletedChallange: string(completed_challange),
				KeyUrl:             keyUrl,
			},
		},
	}
	fmt.Println(*serverMessage)
	// Send the response to the client
	if err := stream.Send(serverMessage); err != nil {
		log.Printf("Error sending server message: %v", err)
		return err
	}
	return nil
}

func handleClientAuthentication(stream verifiable_credentials.VerifiableCredentialService_IssueCredentialServer, clientAuthentication *verifiable_credentials.ClientAuthentication, grpcRequestData *RpcRequestData) {

	fmt.Println(*clientAuthentication)

	serverMessage := &verifiable_credentials.ServerMessage{
		Message: &verifiable_credentials.ServerMessage_Sp{
			Sp: &verifiable_credentials.ServerClaimsProposal{
				// Set appropriate fields based on your logic
				Claims: grpcRequestData.vcData.Claims,
			},
		},
	}
	fmt.Println(*serverMessage)
	// Send the response to the client
	if err := stream.Send(serverMessage); err != nil {
		log.Printf("Error sending server message: %v", err)
		//	return err
	}
}

func handleClientResponse(stream verifiable_credentials.VerifiableCredentialService_IssueCredentialServer, ClientResponse *verifiable_credentials.ClientResponse, grpcRequestData *RpcRequestData) {
	fmt.Println(*ClientResponse)

	currentTime := time.Now()

	// Add 2 hours and 30 minutes to the current time

	// Format the new time as a string in the desired layout
	dateString := currentTime.Format("2006-01-02 15:04:05")

	serverMessage := &verifiable_credentials.ServerMessage{
		Message: &verifiable_credentials.ServerMessage_Vc{
			Vc: &verifiable_credentials.VerifiableCredential{
				// Set appropriate fields based on your logic
				Id:           "my did",
				Type:         grpcRequestData.vcData.Type,
				Issuer:       grpcRequestData.vcData.Issuer,
				Subject:      grpcRequestData.clientDidString,
				IssuanceDate: dateString,
				Claims:       grpcRequestData.vcData.Claims,
				Proof: map[string]string{
					"one":   "1",
					"two":   "2",
					"three": "3",
				},
			},
		},
	}

	fmt.Println(*serverMessage)
	// Send the response to the client
	if err := stream.Send(serverMessage); err != nil {
		log.Printf("Error sending server message: %v", err)
		//	return err
	}
}

var res *resolver.EthipfsResolver

func main() {

	res := resolver.MakeEthIpfsResolver()
	res.ResolveDID("did:ethipfs:0x0f7773CE819dFDF650a6b646E8d34aF63d5E40C4:ETF-Beograd-BU")
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	verifiable_credentials.RegisterVerifiableCredentialServiceServer(server, &verifiableCredentialServer{})

	fmt.Println("gRPC server is running on port 50051...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func read_keys() []map[string]string {
	//filePath := "../data/vc_data-1234-5678-9abc-def0.json"
	filePath := "./data/private_keys.json"
	// Read the content of the JSON file
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty interface to store the parsed JSON data
	var jsonData map[string][]map[string]string

	// Unmarshal the JSON data into the empty interface
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// Now, you can work with the parsed JSON data dynamically
	return jsonData["keys"]
}

type VCData struct {
	Issuer string            `json:"issuer"`
	Type   []string          `json:"type"`
	Claims map[string]string `json:"claims"`
	// Add more fields as needed
}

func read_vc_data(rcp_id string) (*VCData, error) {
	filePath := "./data/vc_data-" + rcp_id + ".json"
	//	filePath := "./data/vc_data-1234-5678-9abc-def0.json"
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty interface to store the parsed JSON data
	var jsonData VCData

	// Unmarshal the JSON data into the empty interface
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		return nil, err
	}

	return &jsonData, nil

}
