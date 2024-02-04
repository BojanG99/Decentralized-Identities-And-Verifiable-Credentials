package resolver

import (
	"errors"
	"fmt"

	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/cache"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/did"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/diddocument"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/didipfs"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/networks"
)

type EthipfsResolver struct {
	MemCache cache.MemoryCache
}

func (e EthipfsResolver) Method() string {
	return ETHIPFS
}

func (e EthipfsResolver) ResolveDID(didstring string) (*diddocument.Document, error) {
	d, err := did.Parse(didstring)
	if err != nil {
		return nil, err
	}
	if d.Method != e.Method() {
		return nil, errors.New("This resolver is only suporting" + e.Method() + "method")
	}

	dd, ok := e.MemCache.Get(d)
	if ok {
		return dd, nil
	}
	//TODO
	// call js method to get cid
	c, err := GetCIDFromRegistryB(d.IDStrings[0], d.IDStrings[1], networks.SEPOLIA, "" /*not yet supported*/)

	if err != nil {
		return nil, err
	}

	if c.IsRevoked {
		return nil, errors.New("Did document i revoked")
	}

	//return nil, nil

	//with cid get file from ipfs
	var ppIpfs didipfs.PublicPinataIpfs

	dd, err = ppIpfs.DownloadDIDDocumentFromIPFS(c.CidString)

	if err != nil {
		return nil, err
	}

	e.MemCache.Put(d, dd)

	return dd, nil
}

func (e EthipfsResolver) ResolveDIDKey(didurl string) (*diddocument.VerificationMethod, error) {
	d, err := did.Parse(didurl)
	if err != nil {
		return nil, err
	}
	if !d.IsKeyURL() {
		return nil, errors.New("Is not key url")
	}
	dd, err := e.ResolveDID(didurl)
	if err != nil {
		return nil, err
	}
	fmt.Println(didurl)
	for _, pubKey := range dd.VerificationMethod {
		fmt.Println(pubKey.ID + " key")
		if pubKey.ID == didurl {
			return &pubKey, nil
		}
	}

	return nil, errors.New("No key with id " + d.Fragment)
}

func MakeEthIpfsResolver() *EthipfsResolver {
	var ethResolver EthipfsResolver
	ethResolver.MemCache = make(map[string]*diddocument.Document)

	return &ethResolver
}
