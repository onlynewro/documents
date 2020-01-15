# 표준 입출력 예제 및 정리
* 다양한 입출력
* curl
* json

## 다양한 입출력
 - 모든 unix os 는 file descriptor 를 사용한다. 각각의 넘버는 0,1,2 이다
 - unix system은 /dev/stdin, /dev/stdout, /dev/stderr
 - mac os 는 /dev/fd/0 
 - debian os 는 /dev/fd/0 과 /dev/pts/0 둘다 사용
 - go에서는 os.Stdin, os.Stdout, os.Stderr 접근 가능
 
1. fmt
    - 화면 출력
2. os.Stdin, os.Stdout
    - 출력
    - 입력
3. io.Reader, io.Writer
    - bufio 출력
    - bufio 입력
        * 텍스트파일 한 단어씩 읽기
        * 텍스트파일 한 문자씩 읽기
4. 사용자로부터 입력 받기
    1. 커맨드 라인 인수를 읽기
    <pre><code>
    import (
        "os"
    )
    if len(os.Args) == 1 {
        os.Exit(1)
    }
    arguments := os.Args
    </code></pre>
    2. Stdin 을 이용해 사용자에게 입력 값을 받기
    <pre><code>
    import (
        "bufio"
        "ftm"
        "os"
    )
    var f *os.File
    f = os.Stdin
    defer f.Close()
    scanner := bufio.NewScanner(f)
    </code></pre>
    3. 외부 파일 읽기
5. /dev/random 읽기
    - 시스템 디바이스로부터 데이터를 읽는 방법

