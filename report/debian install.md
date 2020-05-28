# 데비안(debian) 설치하기
 쉬운 설치 방법 안내

## 데비안 설치 전 준비물 및 알아야 할 사항
 데비안을 설치전 
### 필요한 물품
* usb 메모리
* 설치할 데비안을 usb에 쓰기위한 컴퓨터 
* https://www.debian.org/distrib/ 에서 iso 내려 받기
* iso 이미지를 usb에 설치 
* usb 우선 부팅으로 변경
* 그래픽 설치 및 기본 설치하면 거이 자동으로 다 설치됨
* 온보드가 아닌 nvidia 그래픽카드를 쓴다면 카드의 종류마다 설치과정이 다를 수 있다.

## 데비안에서 한글 입력
* 설치 단계에서 대한민국을 선택하고 그래픽환경을 선택했다면 기본으로 ibus가 설치되어있으나 한글을 쓸수 없다.
    - 설정 -> 지역 및 언어 -> 입력 소스 에서 한국어(hangul) 을 추가 첫번째 입력소스롤 선택
    - 단축키 설정 후 shift + space 로 한영 전환이 가능하다.

### home안에 디렉토리들의 한글명을 영어로 전환
    - terminal 에서 입력
    ```shell
        $lang=C
        $xdg-user-dirs-gdk-update
    ```

### 초기 상태 백업
* timeshift 설치 후 상태 저장하면 에러 및 꼬인 상태를 쉽게 되돌릴 수 있음.(강추)

### fusuma touchpad 설정

<https://github.com/iberianpig/fusuma>
- xdotool은 wayland,x11,xorg 중 지원이 안될 수 있음
- 일부 설정은 플러그인을 설치해야 함

### 윈도우 글꼴 사용하기
윈도우 폰트 폴더에서 필요한 폰트를 복사
```shell
sudo cp *.* /usr/share/fonts 
sudo fc-cache -fc  
```
super + a(모든프로그램보기) -> 기능개선 -> 글꼴 -> 해당폰트 선택

