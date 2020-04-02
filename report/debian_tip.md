#### bash 단축키
ctrl + r : 히스토리 찾기

#### 파일 찾기
```shell
$ find / -name 파일명
```

#### grub 설정 변경
```shell
$ cat /usr/sbin/update-grub | grep grub-mkconfig
$ exec /usr/sbin/grub-mkconfig -o /boot/grub/grub.cfg "$@"
```

#### 노트북 파워모니터링 앱
```shell
$ /usr/sbin/powertop 실행
```

#### tlp - 노트북 베터리 절약 설정 앱

#### godoc 사용위해 apt-get install golang-golang-x-tools
 
#### debian 패키지들은 파일 위치가 기본적으로 다른 리눅스와 다름 항상 확인 필요함