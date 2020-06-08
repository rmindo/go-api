package lib

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

/**
 * Private key struct
 */
type privateKey struct {
	*rsa.PrivateKey
}

/**
 * An interface of signer
 */
type Signer interface {
	Sign(data []byte) ([]byte, error)
}

/**
 * Sign data with private key
 */
func Sign(data string) string {

	signed, err := getSigner("private").Sign([]byte(data))
	if err != nil {
		fmt.Errorf("Unable to sign data: %v", err)
	}
	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(signed)))
}

/**
 * Parse the public key
 */
func parse(bytes []byte) (Signer, error) {

	var raw interface{}

	key, _ := pem.Decode(bytes)
	if key == nil {
		return nil, errors.New("Key is not exist")
	}

	switch key.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(key.Bytes)
		if err != nil {
			return nil, err
		}
		raw = rsa
	default:
		return nil, fmt.Errorf("Key type %q is not supported", key.Type)
	}
	return newSigner(raw)
}

/**
 * New signer
 */
func newSigner(k interface{}) (Signer, error) {
	var key Signer
	switch t := k.(type) {
	case *rsa.PrivateKey:
		key = &privateKey{t}
	default:
		return nil, fmt.Errorf("Key type %T is not supported", k)
	}
	return key, nil
}

/**
 * Get signer of the data
 */
func getSigner(name string) Signer {

	content, err := ioutil.ReadFile("certs/" + name + ".pem")
	if err != nil {
		fmt.Print(err)
	}
	signer, err := parse([]byte(content))

	if err != nil {
		fmt.Errorf("Something wrong with the signer: %v", err)
	}
	return signer
}

/**
 * Sign data with SHA256 algo
 */
func (r *privateKey) Sign(data []byte) ([]byte, error) {

	hash := sha256.New()
	hash.Write(data)
	sum := hash.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, sum)

	if err != nil {
		fmt.Errorf("Unable to sign data: %v", err)
	}
	return sign, err
}
