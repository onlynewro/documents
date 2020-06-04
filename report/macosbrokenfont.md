# MacOs 에서 linux로 파일 옮기면 한글 깨지는 문제
 맥에서 리눅스로 파일을 옮기면 한글이 이상하게 보이는 문제가 발생. 동일한 utf-8 환경이지만 리눅스는 NFC(normalization form C) 맥은 NFD(normalization form D) 방식을 사용 하기 때문임
 
## convmv  설치
```shell
    # apt install convmv
```
## 명령
```shell
    $ convmv -f utf8 -t utf8 -r --nfc --notest *
```
