# console I/O 정리
* 다양한 입출력
* curl
* json

## 파일 입출력

## 다양한 입출력
 - 모든 unix os 는 file descriptor 를 사용한다. 각각의 넘버는 0,1,2 이다
 - unix system은 /dev/stdin, /dev/stdout, /dev/stderr
 - mac os 는 /dev/fd/0 
 - debian os 는 /dev/fd/0 과 /dev/pts/0 둘다 사용
 - go에서는 os.Stdin, os.Stdout, os.Stderr 접근 가능

### 키보드 입력받기

1. fmt
    - Print(), Println(), Printf()
    - Scanf(), 
    <pre><code>
        # Read integer 
        var i int
        fmt.Scanf("%d", &i)

        # Read string 
        var str string
        fmt.Scanf("%s", &str)
    </code></pre>

    - Scan()
    <pre><code>
    # Read integer 
        var i int
        fmt.Scan(&i)

    # Read string 
        var str string
        fmt.Scan(&str)
    </code></pre>
    
    - Scasnln()
    <pre><code>
        # Read string
        var input string
        fmt.Scanln(&input)
    </code></pre>

2. os.Stdin, os.Stdout
    - 출력
    - 입력
3. io.Reader, io.Writer
    - bufio 출력
    - bufio 입력
    <pre><code>
        import (
            "bufio"
            "fmt"
            "os"
        )
        var f *os.File
        f = os.Stdin
        defer f.Close()
        scanner := bufio.NewScanner(f)
    </code></pre>
        * 텍스트파일 한 단어씩 읽기
        * 텍스트파일 한 문자씩 읽기
    <pre><code>
        # Read using Reader
        reader := bufio.NewReader(os.Stdin)
        text, err := reader.ReadString('\n')

        # Read using Scanner
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }
    </code></pre>


### 커맨드 라인 인수를 읽기
1. os.Args 로 직접 컨트롤 하기
    <pre><code>
    import (
        "os"
    )
    if len(os.Args) == 1 {
        os.Exit(1)
    }
    arguments := os.Args
    </code></pre>

2. flag 표주 패키지 이용하기
3. go-glag 패키지 이용하기     

### /dev/random 읽기
    - 시스템 디바이스로부터 데이터를 읽는 방법


