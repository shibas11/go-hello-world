package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const MSG = "kof xi is the one of the best battle game in the world" // 문자만 54자, EOF포함하면 55

func BufIoTest() {
	//
	// 파일 만들기
	//
	myFile, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
	}
	defer myFile.Close()

	//
	// 파일에다가 쓰기(파일이 data stream)
	//
	// w := bufio.NewWriter(myFile)
	w := bufio.NewWriterSize(myFile, 50)

	// n, err := w.Write([]byte(MSG))
	n, err := w.WriteString(MSG)
	if err != nil {
		panic(err)
	}
	fmt.Println("Write:", n)

	w.Flush()
	// seek 위치 조정(위에서 write & flush 했으므로,터 seek 포인터가 파일 끝에 가 있음)
	myFile.Seek(0, io.SeekStart)

	//
	// 파일로부터 읽어오기(파일이 data stream)
	//
	r := bufio.NewReader(myFile)
	myBuf := getBuffer(myFile)
	r.Read(myBuf) // 읽어온 내용을 myBuf로 저장하기

	fmt.Println("bytes :", myBuf)
	fmt.Println("string:", string(myBuf))
}

func getBuffer(myFile *os.File) []byte {
	var size int64

	// 1) 54글자만 읽어봐
	//size = 54

	// 2) MSG 길이만큼 읽어봐
	//size = int64(len(MSG))

	// 3) 파일 stat.Size()만큼 읽어봐
	fi, _ := myFile.Stat()
	size = fi.Size()

	return make([]byte, size)
}
