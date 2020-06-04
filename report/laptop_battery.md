## 리눅스에서 배터리 관리(데비안)
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
* gnome shell 프로그램 중에 lenovo ideapad goodies 이용하면 60% 미만일 경우 충전 되며 60% 충전 이후 외부전력만 사용. 
'''shell
    /sys/bus/platform/drivers/ideapad_acpi/VPC2004:00/conservation
'''
* 위 파일 값을 1로 바꾸면 나뭇잎 모양 아이콘이 나오면 완료
* 90% 이상 충전이 필요할 시에는 0으로 바꿈.
* root 권한 필요함

### 4. 배터리 관련 정보
[TLP - Optimize Linux Laptop Battery Life](https://linrunner.de/tlp/index.html)

### 5. 노트북 전원모니터링 앱
```shell
 # apt install powertop
 $ /usr/sbin/powertop 
```
