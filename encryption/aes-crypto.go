package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Encrypt(block cipher.Block, plainText []byte) []byte {
	if mod := len(plainText) % aes.BlockSize; mod != 0 { // 16의 배수가 아니면 맞춰줄 것
		padding := make([]byte, aes.BlockSize-mod)
		plainText = append(plainText, padding...)
	}

	cipherText := make([]byte, len(plainText))
	block.Encrypt(cipherText, []byte(plainText))

	return cipherText
}

func Decrypt(block cipher.Block, cipherText []byte) []byte {
	plainText := make([]byte, len(cipherText))
	block.Decrypt(plainText, cipherText)

	return plainText
}

func EncryptWithIV(block cipher.Block, plainText []byte) []byte {
	if mod := len(plainText) % aes.BlockSize; mod != 0 { // 16의 배수가 아니면 맞춰줄 것
		padding := make([]byte, aes.BlockSize-mod)
		plainText = append(plainText, padding...)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText)) // iv를 위한 공간 추가

	iv := cipherText[:aes.BlockSize] // 초기화 벡터공간(슬라이스 앞쪽)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText
}

func DecryptWithIV(block cipher.Block, cipherText []byte) []byte {
	if len(cipherText)%aes.BlockSize != 0 {
		fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야합니다.")
		return nil
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	plainText := make([]byte, len(cipherText))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)

	return plainText
}
