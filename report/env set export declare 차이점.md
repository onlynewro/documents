## env set export declare 차이점

리눅스에서 가장 많이 사용하는 Bash(혹은 sh)의 빌틴 명령어

각 명령어마다 쓰임새가 조금씩 다름

### env : 환경변수를 보여주거나, 설정 혹은 삭제 하는 명령. http://ss64.com/bash/env.html

```shell
$ env
```
현재 세션에 정의된 환경 변수들을 화면에 출력

```shell
$ env NAME=VALUE 
```
NAME이라는 환경변수에 VALUE라는 값을 지정

```shell
$ env -u NAME
```
NAME 환경변수를 삭제

### set : Bash의 쉘 변수를 관리하는 명령 http://ss64.com/bash/set.html
```shell
$set NAME=VALUE 
```

Bash에서는 다음과 같이 생략이 가능
```shell
NAME=VALUE
```
쉘 변수는 Bash라는 쉘 스크립트 언어에서 사용하는 변수이고, 환경 변수는 운영체제에서 사용하는 변수(예: PATH)

해제하는 방법은 unset 명령어를 사용

### export: 쉘 변수를 환경 변수로 변경해주는 명령
```shell
NAME=VALUE
export NAME
```
위는 set과 export를 사용하여 다음 명령과 동일한 결과

```shell
env NAME=VALUE
```
### declare: 이는 쉘 변수의 타입을 지정하는 명령. https://wiki.kldp.org/HOWTO/html/Adv-Bash-Scr-HOWTO/declareref.html

읽기전용, 정수, 배열, 함수 등으로 쉘 변수의 성격을 선언할 때 사용