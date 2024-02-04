package cache

import (
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/did"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/diddocument"
)

type Cache interface {
	// Get checks the cache for a given did, and if found, returns the previous result.
	Get(parsed *did.DID) (*diddocument.Document, bool)
	Put(parsed *did.DID, ddoc *diddocument.Document) error
}
