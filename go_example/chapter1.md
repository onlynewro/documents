# 1장 go 와 os
* 사용하지 않는 패키지는 import 하면 컴파일 에러 발생
* import 목록에서 패키지 이름 앞에 _ 를 붙이면 그 패키지를 프로그램에서 사용하지 않도라도 컴파일 에러가 발생하지 않는다.
* stdin, stdout, stderr -> go os.Stdin, os.Stdout, os.Stderr
> fmt.Print(), fmt.Println(), fmt.Fprint(), io.Writer
* 표준출력
> io.WriteStirng(os.Stdout, string)
* 사용자로부터 입력받기
1. 프로그램의 커맨드라인 인수를 읽는다.
<pre>
    os.Args
</pre>

2. 사용자에게 입력값을 물어본다.
<pre>
    var f *os.File
    f = os.Stdin
    defer f.Close()
    
    scanner := bufio.NewScanner(f) 
</pre>

* := 짧은 할당문 = 묵시적 타입의 var 선언문 대신 사용할 수 있다.
* _ 빈 식별자

## 에러출력
1. 로그파일 작성 
> log.Printf(), log.Print(), log.Println(), log.Fatalf(), logFatalln(), log.Panic(), log.Panicln(), log.Panicf()
2. 로그 수준
> debug, info, notice, warning, err, crit, alert, emerg 등
3. 로그 종류
> auth, cron, kern, mail 등등
* log, log/syslog 패키지
>log.Fatal() log.Panic(), 
 ## 에러처리
 나중에 