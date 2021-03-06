package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// r은 응답 객체이며 r.Body 필드가 io.Reader 인터페이스를 구현한다.
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln("err")
	}

	// 응답을 저장하기 위한 파일을 생성한다.
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln("erre")
	}

	// MultiWriter를 이용하여 한 번의 쓰기 작업으로 표준 출력장치와 파일에 같이 출력
	dest := io.MultiWriter(os.Stdout, file)

	// 응답을 읽어 파일과 표준 출력 장치에 출력하다.
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
