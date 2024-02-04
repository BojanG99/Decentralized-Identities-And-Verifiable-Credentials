package cache

import (
	"errors"
	"fmt"

	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/did"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/diddocument"
)

type MemoryCache map[string]*diddocument.Document

func (mc MemoryCache) Get(parsed *did.DID) (*diddocument.Document, bool) {

	cached, ok := mc[parsed.DIDString()]
	if ok {
		return cached, true
	}

	return nil, false

}

func (mc MemoryCache) Put(parsed *did.DID, ddoc *diddocument.Document) error {

	if ddoc == nil || parsed == nil {
		return errors.New("Arguments are nil")
	}
	_, ok := mc[parsed.DIDString()]
	if ok {
		return errors.New("Document is already in memory")
	}
	fmt.Println(parsed)
	mc[parsed.DIDString()] = ddoc
	return nil

}
