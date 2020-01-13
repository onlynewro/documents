package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Buffer 값을 생성한 후 버퍼에 문자열을 출력한다.
	// 이때 io.Writer 인터페이스를 구현한 Write 메서드를 호출한다.
	var b bytes.Buffer
	b.Write([]byte("안녕하세요"))

	// 버퍼에 문자열을 결합하기 위해 Fprintf 함수를 호출한다.
	fmt.Fprintf(&b, "Golang!")

	// 버퍼의 콘텐츠를 표준 출력 장치에 쓴다.
	b.WriteTo(os.Stdout)
}
