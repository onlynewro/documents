# 데비안(debian) gnome-desktop 팁

## 데비안 설치 전 알아야 할 사항
 * 한국에서는 우분투 사용률이 높아서 데비안에대한 정보가 부족
 * 우분투와는 비슷하지만 조금 다름

## 설치 팁
* 설치 시디 선택시 라이브 시디를 선택해서 라이브 모드로 부팅을 해보는 것이 좋음.
* 와이파이 그래픽카드 등 하드웨어적인 문제를 알아서 처리해줌. 
* 라이브 부팅이 잘된다면 설정을 따라서 할 수 있음

## 데비안에서 한글 입력
* 설치 단계에서 대한민국을 선택하고 그래픽환경을 선택했다면 기본으로 ibus가 설치되어있으나 한글 입력이 안됨
    - 설정 -> 지역 및 언어 -> 입력 소스 에서 한국어(hangul) 을 추가 첫번째 입력소스로 선택
    - 단축키(shift + space)등으로 설정 후 한영 전환이 가능하다.

## home 디렉토리들의 한글명을 영어로 전환 방법
    - terminal 에서 입력
    ```shell
        $lang=C
        $xdg-user-dirs-gdk-update
    ```

## 초기 상태 백업
* timeshift 설치 후 상태 저장하면 에러 및 꼬인 상태를 쉽게 되돌릴 수 있음.(강추)

## fusuma touchpad 설정
[fusuma 사이트](https://github.com/iberianpig/fusuma)
- xdotool은 wayland,x11,xorg 중 지원이 안될 수 있음
- 일부 설정은 플러그인을 설치해야 함

## 윈도우 글꼴 사용하기
윈도우 폰트 폴더에서 필요한 폰트를 복사
```shell
sudo cp *.* /usr/share/fonts 
sudo fc-cache -fc  
```
super + a(모든프로그램보기) -> 기능개선(tweaks) -> 글꼴(font) -> 해당폰트 선택

