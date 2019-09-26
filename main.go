package main

import (
	"crypto"
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/shibas11/go-hello-world/encryption"
	"github.com/shibas11/go-hello-world/multiplexer/exam1"
	"github.com/shibas11/go-hello-world/multiplexer/exam2"
	"github.com/shibas11/go-hello-world/multiplexer/exam3"
	"github.com/shibas11/go-hello-world/network/grpcTest"
	"github.com/shibas11/go-hello-world/network/rpc"
	"github.com/shibas11/go-hello-world/network/tcp"
	pb "github.com/shibas11/go-hello-world/network/types"
	"github.com/shibas11/go-hello-world/stringutil"
	"log"
)

func main() {
	// https://www.integralist.co.uk/posts/understanding-golangs-func-type/
	examNumber := 3
	switch examNumber {
	case 1:
		exam1.Exam()
	case 2:
		exam2.Exam()
	case 3:
		exam3.Exam()
	}

	s := "Hello, world."
	fmt.Println(s)
	fmt.Println(stringutil.Reverse(s))

	fmt.Println()
	symmetricTest(s)

	fmt.Println()
	rsaEncryptionTest()

	fmt.Println()
	rsaSignAndVerifyTest()

	//fmt.Println()
	//tcpServerTest()

	//fmt.Println()
	//rpcServerTest()

	fmt.Println("\nprotobuf test")
	protobufTest()

	fmt.Println("\ngrpc test")
	grpcAndProtobufTest()

	fmt.Print("Bye.")
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

func rsaSignAndVerifyTest() {
	message := "안녕하세요. Go 언어"
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15( // 개인키로 서명
		rand.Reader,
		privateKey,
		h1,
		digest,
	)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15( // 공개키로 검증
		publicKey,
		h2,
		digest,
		signature,
	)

	if err != nil {
		fmt.Println("검증 실패:", err)
	} else {
		fmt.Println("검증 성공")
	}

}

func tcpServerTest() {
	server := tcp.NewServer(8080)
	if err := server.Serve(); err != nil {
		fmt.Println(err)
		return
	}
}

func rpcServerTest() {
	server := rpc.NewServer(6000)
	if err := server.Serve(); err != nil {
		fmt.Println(err)
		return
	}
}

func protobufTest() {
	p := pb.Person{
		Name:  "Jake",
		Id:    1,
		Email: "abc@abc.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "1234-5678", Type: pb.Person_MOBILE},
		},
	}

	book := &pb.AddressBook{
		People: []*pb.Person{
			&p,
		},
	}

	// Write the new address book back to disk.
	outBytes, err := proto.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	// Read the existing address book.
	newBook := &pb.AddressBook{}
	proto.Unmarshal(outBytes, newBook)
	fmt.Println(newBook)
}

func grpcAndProtobufTest() {
	server := grpcTest.NewServer(50051)
	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to start server: " + err.Error())
		return
	}
}
