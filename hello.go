package main

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/shibas11/go-hello-world/encryption"
	"github.com/shibas11/go-hello-world/stringutil"
)

func main() {
	s := "Hello, world."
	fmt.Println(s)
	fmt.Println(stringutil.Reverse(s))

	fmt.Println()
	symmetricTest(s)

	fmt.Println()
	rsaEncryptionTest()
}

func symmetricTest(s string) {
	key := "Hello, Key 12345" // 16바이트
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("originalText: %s\n", s)
	cipherText := encryption.Encrypt(block, []byte(s))
	fmt.Printf("  cipherText  : %x\n", cipherText)
	plainText := encryption.Decrypt(block, cipherText)
	fmt.Printf("  plainText   : %s\n\n", string(plainText))

	fmt.Printf("originalText: %s\n", s)
	cipherText = encryption.EncryptWithIV(block, []byte(s))
	fmt.Printf("  cipherTextWithIV  : %x\n", cipherText)
	plainText = encryption.DecryptWithIV(block, cipherText)
	fmt.Printf("  plainTextWithIV   : %s\n", string(plainText))
}

func rsaEncryptionTest() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := &privateKey.PublicKey

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	cipherText := encryption.EncryptRSA(publicKey, s)
	fmt.Printf("%x\n", cipherText)

	plainText := encryption.DecryptRSA(privateKey, cipherText)
	fmt.Println(string(plainText))
}
