## 리눅스에서 배터리 관리
 cli tool인 tlp를 통해서 관리

### 1. 설치 방법
```shell
 $ apt install tlp
```

### 2. 유용한 옵션
* 배터리 정보 확인
>```shell
> $ tlp-stat -b
>```
* 시스템 정보 확인
>```shell
> $ tlp-stat -s
>```

### 3. 배터리 과충전 방지
* Thinkpad 만 가능 나머지는 아직 안됨
* 간단한 tlp 스크립트 생성으로 조작 가능

### 4. 관련 정보
[TLP - Optimize Linux Laptop Battery Life](https://linrunner.de/tlp/index.html)


