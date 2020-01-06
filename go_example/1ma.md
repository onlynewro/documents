# main package
- go로 구현하는 모든 실행 가능한 프로그램들은 반드시 main이라는 이름의 패키지를 가지고 있어야 한다.
- main 패키지는 main() 함수를 선언해야 프로그램 진입점 역활를 수행할 수 있다.
- 컴파일된 실행 바이너리의 이름은 main 패키지가 선언된 디렉토리의 이름을 따온다.

# import 
- 패키지내에 작성된 코드에 접근하게 해준다.
<pre>
    import (
        "fmt"
        "string"
        "github.com/test/tseing"
    }
</pre>
- 컴파일러는 go 설치 디렉토리를 먼저 탐색 후 GOPATH 환경 변수에 나열된 순서대로 탐색한다.
- import 경로가 URL을 포함하고 있다면 go get 명령을 이용해 지정된 URL에서 코드를 가져올 수 있다.
- 다른 사람들이 생성한 패키지는 go get명령을 통해 GOPATH 환경 변수에 지정된 경로에 다운로드 및 설치할 수 있다.
- go get 명령은 재귀적으로 실행되기 때문에 패키지의 소스틀를 모두 탐색 후 의존성 코드들의 대한 가져오기를 수행한다.

# 명명된 import 
<pre>
    package main
    
    import (
        "fmt",
        myfmt "mylib/fmt"
    )

    func main() {
        fmt.Println("표준라이브러리")
        myfmt.Println("mylib/fmt")
    }
</pre>
- 만약 패키지 이름이 동일하지만, 서로 다른 버젼 혹은 서로 다른 위치에서 로딩하고자 할 때는, 패키지 alias를 사용해서 구분할 수 있다.
- 명명된 패키지가 사용되지 않으면 오류를 발생시킴
- 경우에 따라 패키지를 import 하면서 단지 그 패키지 안의 init() 함수만을 호출하고자 하는 케이스가 있다. 이런 경우는 패키지 import 시 _ 라는 alias 를 지정한다. 아래는 other/xlib 패키지를 호출하면서 _ alias를 지정한 예이다.
# init
<pre>
    func init() {
        sql.Register("postgres", new (PostgresDriver))
    }
</pre>
- 패키지는 최초 실행 시점에 init 함수를 필요한 만큼 정의 할 수 있는 기능을 제공한다.
- 모든 init 함수는 main 함수가 실행되기 전에 미리 호출 되도록 예약한다.
- init 함수는 패키지를 설정하거나 변수의 초기화, 또는 프로그램이 실행되기 전에 수행하기 위한 최적의 장소이다.

# go build
- go build
- go bulid path
- go build path/...
- go build test.go
- go bulid .
- go run test.go

# 개발자 도구

## go vet
- 코드상 일반적으로 반생할 수 있는 에러를 검사해 준다.
    * Printf 스타일의 함수 호출 시 잘못된 매개변수 지정
    * 매서드 정의 시 시그너처 관련 에러
    * 잘못 구성된 태그
    * 조합 리터럴 사용 시 누락된 키

## go fmt
- 지정된 소스 코드 파일의 형식을 재조정한 뒤 다시 저장해준다.
- 코드 저장소에 커밋하기 전에 go fmt 사용한다.

## go 문서화 도구
<pre>
    go doc tar
</pre>
- 명령줄에서 문서 탐색하기
<pre>
    godoc -hhtp=:8001
</pre>
- 로컬서버 웹브라우져에서 문서 탐색하기
- godoc 커맨드가 안될때는 golang-golang-godoc 인스톨해야할 때가 있다.

## 다른 go 개발자와 협력하기
- 소스 코드 공유를 위한 저장소 생성
- 패키지는 반드시 저장소의 루트에 저장
- 패키지 크키는 작게 유지
- 코드에 go fmt 명령을 실행하자
- 소스 코드를 문서화 하자

## 의존성 관리
1. 벤더링(godep 또는 vendor)은 import 경로 재작성
2. gb 도구
<pre>
    $PROJECT/src
    $PROJECT/vendor/src
</pre>
* go 는 import 구문 때문에 재사용 가능한 빌드를 만들 수 없다는 점에 기초
* go get 명령이 패키지의 버젼을 가져와야 하는지 명확한 정보를 제공하지 못함
* import 경로를 수정 할 필요 없음.

## go 시작하기
* 코드 작성 기초 방법
* https://github.com/golang-kr/golang-doc/wiki/Go-%EC%BD%94%EB%93%9C%EB%A5%BC-%EC%9E%91%EC%84%B1%ED%95%98%EB%8A%94-%EB%B0%A9%EB%B2%95