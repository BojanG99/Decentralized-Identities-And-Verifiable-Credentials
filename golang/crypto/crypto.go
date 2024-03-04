package kripto

import (
	"errors"
	"fmt"
	"log"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
)

func SignMessage(message string, key jwk.Key) ([]byte, error) {

	crv, ok := key.Get("crv")
	if !ok {
		return nil, errors.New("No crv field in key")
	}
	stringFromCrv := fmt.Sprintf("%s", crv)

	buff, err := jws.Sign([]byte(message), determineSigningAlgorithm(stringFromCrv), key)

	if err != nil {
		log.Printf("Faild to sign message: %s\n", err)
		return nil, err
	}

	return buff, nil

}

func VerifyMessage(payload string, expectedMessage string, key jwk.Key) bool {

	crv, ok := key.Get("crv")
	if !ok {
		return false
	}
	stringFromCrv := fmt.Sprintf("%s", crv)

	verified, err := jws.Verify([]byte(payload), determineSigningAlgorithm(stringFromCrv), key)
	if err != nil {
		log.Printf("failed to verify message: %s", err)
		return false
	}
	if string(verified) == expectedMessage {
		return true
	}
	return false
}

func determineSigningAlgorithm(crv string) jwa.SignatureAlgorithm {
	switch crv {
	case "P-256":
		return jwa.ES256 // ECDSA with NIST P-256 curve
	case "P-384":
		return jwa.ES384 // ECDSA with NIST P-384 curve
	case "P-521":
		return jwa.ES512 // ECDSA with NIST P-521 curve
	case "Ed25519":
		return jwa.EdDSA
	default:
		return "Unknown"
	}
}

// func main() {
// 	test := `
// 	{
// 		"kty": "EC",
// 		"d": "7xP2kh581uUaTvhlXqTPsHYfQQ2ZiiiOZBQf4wOvNo0",
// 		"use": "sig",
// 		"crv": "P-256",
// 		"kid": "key-9",
// 		"x": "dU1s0XRU2EGrsiNZipSnDlB-SWfQfJGJkbpsYTA7kio",
// 		"y": "LNzwKbI6dEa9W7-ON_Qkl2ZWtODMWIee83i54aBg9gM"
// 	}
// `

// 	mkey, err := jwk.ParseKey([]byte(test))

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	proof := "eyJhbGciOiJFZERTQSIsImtpZCI6ImtleS0xMyJ9.SGVsbG8sIHdvcmxkIQ.8QCSft0wie_WbQqjYWeh0-x5NFJxNlzeIK-tY6bqR6kCVm_I0KeC-XToEn_DO-VpZ0i7rbWB4erAbtAiY4d_Cw"
// 	q, _ := mkey.PublicKey()
// 	crv, ok := mkey.Get("crv")
// 	if !ok {
// 		return
// 	}
// 	stringFromCrv := fmt.Sprintf("%s", crv)

// 	verified, err := jws.Verify([]byte(proof), jwa.SignatureAlgorithm(stringFromCrv), q)
// 	fmt.Println(err)
// 	fmt.Println(string(verified))
// }
