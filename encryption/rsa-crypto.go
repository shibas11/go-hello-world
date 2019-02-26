package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func DecryptRSA(privateKey *rsa.PrivateKey, cipherText []byte) []byte {
	plainText, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		cipherText,
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return plainText
}

func EncryptRSA(publicKey *rsa.PublicKey, s string) []byte {
	cipherText, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cipherText
}
