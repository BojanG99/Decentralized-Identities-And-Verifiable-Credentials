package diddocument

import "encoding/json"

type Document struct {
	Context     []string `json:"@context"`
	ID          string   `json:"id"`
	AlsoKnownAs string   `json:"alsoKnownAs,omitempty"`
	//string or []string
	Controller         []string             `json:"controller,omitempty"`
	VerificationMethod []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication     []string             `json:"authentication,omitempty"`
	AssertionMethod    []string             `json:"assertionMethod,omitempty"`
	KeyAgreement       []string             `json:"keyAgreement,omitempty"`
	//	CapabilityInvocation []VerificationMethod
	//	CapabilityDelegation []VerificationMethod
	//	service              []ServiceEndpoint
}

type VerificationMethod struct {
	ID                 string
	Controller         string
	Type               string
	PublicKeyJWK       map[string]interface{}
	publicKeyMultibase string
}

func Parse(didDocument string) (Document, error) {
	var didObject Document
	err := json.Unmarshal([]byte(didDocument), &didObject)
	if err != nil {
		return didObject, err
	}

	return didObject, err
}
