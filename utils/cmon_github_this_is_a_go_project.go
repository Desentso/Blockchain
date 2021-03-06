package utils

import (
	"fmt"
	"crypto/sha256"
	"crypto/x509"
	"crypto/rsa"
	"encoding/hex"
	"encoding/pem"
)

/*

This is copy-paste from the utils file
Only because Github calculates the language
distiributions in a weird way. 
I want this project to show up as Go project not 
JavaScript, because in it's core it's a Go project.

*/

func ABCCalculateHash(blockString string) string {
	strAsBytes := []byte(blockString)
	sum := sha256.Sum256(strAsBytes)

	return hex.EncodeToString(sum[:])
}

func ABCPublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		fmt.Println("Error while converting public key to bytes")
	}
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

func ABCPrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

func ABCBytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Println(err)
	}
	return key
}

func ABCBytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		fmt.Println(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		fmt.Println("not ok")
	}
	return key
}


/*

This is copy-paste from the utils file
Only because Github calculates the language
distiributions in a weird way. 
I want this project to show up as Go project not 
JavaScript, because in it's core it's a Go project.

*/
