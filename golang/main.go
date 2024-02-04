package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os/exec"

	addresbook "github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/gen"
	"github.com/BojanG99/Decentralized-Identitise-And-Verifiable-Credentials/resolver"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//runjs()
	// d, err := resolver.GetCIDFromRegistry("0x0f7773CE819dFDF650a6b646E8d34aF63d5E40C4", "testDid1", networks.SEPOLIA, "")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(d.CidString)
	//var memcache resolver.MemoryCache

	ethResolver := resolver.MakeEthIpfsResolver()
	dd, err := ethResolver.ResolveDID("did:ethipfs:0x0f7773CE819dFDF650a6b646E8d34aF63d5E40C4:ETF-Beograd-BUx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dd)

	key, err := ethResolver.ResolveDIDKey("did:ethipfs:0x0f7773CE819dFDF650a6b646E8d34aF63d5E40C4:ETF-Beograd-BU#key1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key)

	//did_blockchain.Tester()
	// fileUrl := `https://aquamarine-big-bison-5.mypinata.cloud/ipfs/Qmc5MXfyXtuw4cerf8E2PpNSYb27getLgomsWds2Dtjdkk`
	// data, err := didipfs.DownloadDIDDocumentFromIPFS(fileUrl)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))
	// didDocument := make(map[string]interface{})
	// err = json.Unmarshal(data, &didDocument)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(didDocument)
	// myDid := "did:eheripfs:S:0x1234567888;service=agent/path1/path2/path3/path4?service=files&relativeRef=/resume.pdf"
	// d, err := did.Parse(myDid)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("Method - %s, ID - %s", d.Method, d.IDStrings[1])
	// fmt.Println(d.Params)

	// diddocument := `{
	// 	"@context": [
	// 	  "https://www.w3.org/ns/did/v1",
	// 	  "https://w3id.org/security/suites/ed25519-2020/v1"
	// 	],
	// 	"id": "did:example:123456789abcdefghi",
	// 	"verificationMethod": [
	// 		{
	// 		"id": "did:example:123#key1",
	// 		"type": "JsonWebKey2020",
	// 		"controller": "did:example:123",
	// 		"publicKeyJwk": {
	// 		  "crv": "Ed25519",
	// 		  "x": "08YtQxbP2-FeiLVmABuIOUxAoTjtyWRX60BExY6hsVw",
	// 		  "kty": "OKP",
	// 		  "kid": "odStaR5pTJ0t-QtuIlGwh_LC8SBqpqpJ8VnsKD-xAPE"
	// 		}
	// 	  }
	// 	],
	// 	"authentication": ["did:example:123#key1"],
	// 	"assertionMethod": ["did:example:123#key1"]

	// }`
	// fmt.Println(diddocument)
	// didObject, err := resolver.DIDDocumentResolver([]byte(diddocument))

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(didObject)

}

func runjs() {
	hhDir := "../blockchain/"

	// Run the JavaScript file using Node.js
	cmd := exec.Command("npx", "hardhat", "get_cid", "--contractaddress", "0xfBfE37229AE38F83E0C718efF6E8B5A1194C9127", "--prefixname", "BojanG", "--suffixname", "BojanG", "--network", "sepolia")
	cmd.Dir = hhDir
	// Capture and print the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Output:\n%s", output)
}

func maint() {

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/cb91be5166984b73ac6e94685fd74987")

	privateKey, err := crypto.HexToECDSA("163c67333bc9ed7e8bedffb16ee89efaaf2a97cfd2a96d885a096b9f01a5398a")
	if err != nil {
		log.Fatal(err)
	}
	add := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	cAdd := common.HexToAddress("0x0f7773CE819dFDF650a6b646E8d34aF63d5E40C4")

	didreg, err := addresbook.NewDIDRegistry(cAdd, client)

	if err != nil {
		log.Fatal(err)
	}
	//auth for transaction

	ime, cid, verzija, revoked, data, err := didreg.GetIdentity(&bind.CallOpts{
		From: crypto.PubkeyToAddress(privateKey.PublicKey),
	}, "BojanGalic2")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ime)
	fmt.Println(cid)
	fmt.Println(verzija)
	fmt.Println(revoked)
	fmt.Println(data)
	// a, tx, _, err := addresbook.DeployAddressBook(auth, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("----------------------------------")
	// fmt.Println(a.Hex())
	// fmt.Println(tx.Hash().Hex())
	// fmt.Println("----------------------------------")
}
