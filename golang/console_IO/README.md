# console I/O 
- os는 화면출력과 키보드 입력 등 처리 작업을 모두 파일로 간주하고 처리함
- 파일읽기쓰기가 console I/O의 입출력이다.
- unix os 는 file descriptor 를 사용한다. 각각의 넘버는 0,1,2 이다
- unix system은 /dev/stdin, /dev/stdout, /dev/stderr
- mac os 는 /dev/fd/0 
- debian os 는 /dev/fd/0 과 /dev/pts/0 둘다 사용
- go에서는 각 descriptor에 os.Stdin, os.Stdout, os.Stderr를 이용해 접근 가능

1. fmt 패키지를 이용한 출력 및 입력
    - Print(), Println(), Printf() 생략
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

2. file descriptor 를 이용한 입출력
    - 출력 os.Stdout
    - 입력 os.Stdin

3. io bufio 패키지를  이용한 입출력
    - Standard_Library 의 bufio 패키지 정리 참조
    - io.Reader, io.Writer 
    - bufio 출력
    - bufio 입력
    https://tutorialedge.net/golang/reading-console-input-golang/ 
    * 텍스트파일 한 단어씩 읽기
    * 텍스트파일 한 문자씩 읽기

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

# 커맨드 라인 인수를 읽기

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
    - strconv 패키지로 스트링타입 인수를 산술데이터로 변경해서 사용
2. flag 표준 패키지
- 표준 Go 라리브러리의 일부로 제공되어 충분한 테스트를 거치고 디버깅 되었다고 불 수 있어 장점
- flag 패키지 정리 참조
- Using FlagSet to Implement Sub-commands
    * https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
3. go-flag 패키지  
- go-flag 패키지 정리 참조   


# /dev/random 읽기
    - 시스템 디바이스로부터 데이터를 읽는 방법


