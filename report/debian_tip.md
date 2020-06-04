#### residual config 삭제
- 필요없는 프로그램의 설정파일을 다 지움
```shell
apt purge $(dpkg -l | grep '^rc' | awk '{print $2}')
```

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

#### 터미널 환경에서 유용하게 쓰이는 mc
- ctrl + o : 평소 화면을 볼 수 있다. 한번 더 누르면 mc로 돌아온다.

#### godoc 사용위해 apt-get install golang-golang-x-tools
 
#### debian 패키지들은 파일 위치가 기본적으로 다른 리눅스와 다름 항상 확인 필요함