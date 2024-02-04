package resolver

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"os/exec"
	"strings"

	addresbook "github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/gen"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/networks"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CIDData struct {
	CidString string `json:"cid"`
	Version   int    `json:"version"`
	IsRevoked bool   `json:"isRevoked"`
}

func GetCIDFromRegistry(contractAddress string, prefixName string, networkName networks.NetworkName, name string) (*CIDData, error) {

	hhDir := "../blockchain/"

	// Run the JavaScript file using Node.js
	cmd := exec.Command("npx", "hardhat", "get_cid", "--contractaddress", contractAddress, "--prefixname", prefixName /* "--suffixname", name,*/, "--network", string(networkName))
	cmd.Dir = hhDir
	// Capture and print the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(string(output), "Error") {
		return nil, errors.New(string(output))
	}
	var cid_data CIDData
	err = json.Unmarshal(output, &cid_data)

	if err != nil {
		return nil, err
	}
	return &cid_data, nil

}

// *TODO add network name directly from argument and hide private key and infura id
func GetCIDFromRegistryB(contractAddress string, prefixName string, networkName networks.NetworkName, name string) (*CIDData, error) {

	client, err := ethclient.Dial(`https://sepolia.infura.io/v3/cb91be5166984b73ac6e94685fd74987`)

	privateKey, err := crypto.HexToECDSA("163c67333bc9ed7e8bedffb16ee89efaaf2a97cfd2a96d885a096b9f01a5398a")
	if err != nil {
		return nil, err
		//log.Fatal(err)
	}
	add := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		return nil, err
		//log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		return nil, err
		//log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		return nil, err
		//log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
		//log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	cAdd := common.HexToAddress(contractAddress)

	didreg, err := addresbook.NewDIDRegistry(cAdd, client)

	if err != nil {
		return nil, err
		//log.Fatal(err)
	}
	//auth for transaction

	_, cid, version, isRevoked, _, err := didreg.GetIdentity(&bind.CallOpts{
		From: crypto.PubkeyToAddress(privateKey.PublicKey),
	}, prefixName)

	if err != nil {
		return nil, err
		//log.Fatal(err)
	}
	// fmt.Println(ime)
	// fmt.Println(cid)
	// fmt.Println(verzija)
	// fmt.Println(revoked)
	// fmt.Println(data)
	var cid_data CIDData
	cid_data.CidString = cid
	cid_data.IsRevoked = isRevoked
	cid_data.Version = int(version.Int64())

	return &cid_data, nil

}
