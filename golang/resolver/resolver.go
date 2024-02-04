package resolver

import (
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/diddocument"
)

const (
	ETHIPFS = "ethipfs"
)

type Resolver interface {
	Method() string
	ResolveDID(didstring string) (*diddocument.Document, error)
	ResolveDIDKey(didstring string) (diddocument.VerificationMethod, error)
}

// func ResolveDid(_did string) (*did.DID, error) {

// 	d, err := did.Parse(_did)
// 	if err != nil {
// 		return d, nil
// 	}
// }
