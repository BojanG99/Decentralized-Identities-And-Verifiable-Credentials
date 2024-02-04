package didipfs

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/diddocument"
)

type IpfsDownloader interface {
	DownloadDIDDocumentFromIPFS(cidstring string) (*diddocument.Document, error)
}

type PinataIpfs struct {
	Gateway      string
	PinataAPIKey string
	PinataJWT    string
}

const publicPinataGateway = "https://gateway.pinata.cloud/ipfs/"

type PublicPinataIpfs struct {
}

func (pipfs PublicPinataIpfs) DownloadDIDDocumentFromIPFS(cidstring string) (*diddocument.Document, error) {

	response, err := http.Get(publicPinataGateway + cidstring)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	fileData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Extract content type from headers
	contentType := response.Header.Get("Content-Type")
	fmt.Println(contentType)

	// Extract file extension from content type
	fileExtension := strings.Split(contentType, "/")[1]

	if fileExtension != "json" {

		return nil, errors.New("cid is not walid JSON did")
	}
	// Write file to local filesystem

	dd, err := diddocument.Parse(string(fileData))

	if err != nil {
		return nil, err
	}

	return &dd, nil

}
