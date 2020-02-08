# 플래그 패키지
- 플래그 패키지를 쓰지 않는 명령 인자 사용

## 플래그를 사용하여 프로그램 동작 변경
* 플래그 패키지 사용에는 세 가지 단계가 포함된다. 
1. 먼저 변수를 정의하여 플래그 값을 캡처한 다음
2. Go 애플리케이션에서 사용할 플래그를 정의하고
3. 마지막으로 애플리케이션으로 제공된 플래그를 구문 분석하십시오.
* 플래그 패키지 내의 대부분의 기능은 플래그를 정의하고 사용자가 정의한 변수에 플래그를 바인딩하는 것과 관련이 있다. 구문 분석 단계는 Parse() 함수로 처리한다.

* https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go

* https://flaviocopes.com/go-command-line-flags/

### exam_1
 - flag.Bool("color", false, "display colorized output") *bool
 - "-color" 옵션이 있는지 없는지 확인하는 함수
 - false 는 지정값이 없을 때 초기화 값을 가지는 인수
 - "display colorized output" help 옵션시 표시
 - flag.Parse() 로 구문 분석
 - flag.Boolvar()는 리턴값 대신 지정한 포인터 변수를 지정한다.
     * flag.Bool( &Bool, "color", false, "display colorized output")



