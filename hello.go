package main

import (
	"crypto/aes"
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
