1.파일 찾기
> find / -name 파일명

2. grub 설정 변경
> cat /usr/sbin/update-grub | grep grub-mkconfig
> exec /usr/sbin/grub-mkconfig -o /boot/grub/grub.cfg "$@"

3. 노트북 파워모니터링 앱
> /usr/sbin/powertop 실행

4. tlp - 노트북 베터리 절약 설정 앱

5. godoc 사용위해 apt-get install golang-golang-x-tools
 
6. 