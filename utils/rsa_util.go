package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
)

func ValidateSignature(challenge string, signatureBase64 string, pemPublicKey string) bool {
	// Parse PEM block
	rsaPubKey, err := ParseRSAPublicKeyFromPEM(pemPublicKey)
	if err != nil {
		log.Println("Error parsing PKIX public key:", err)
	}

	// Decode Base64 signature
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		log.Println("Error decoding signature:", err)
	}

	// Hash the message
	hashed := sha256.Sum256([]byte(challenge))

	// Verify the signature
	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("Error verifying signature:", err)
		return false
	}

	return true
}

// parseRSAPublicKeyFromPEM parses RSA public key from PEM data
func ParseRSAPublicKeyFromPEM(pemPublicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemPublicKey))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	// Attempt to parse as PKCS1 public key
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		return pubKey, nil
	}

	// If parsing as PKCS1 fails, try parsing as PKIX public key
	pubKeyPKIX, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPubKey, ok := pubKeyPKIX.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parsed key is not RSA public key")
	}

	return rsaPubKey, nil
}
