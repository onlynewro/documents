# tmux 구성 요소
+ sesstion:  여러 윈도우로 구성
+ window: 터미널 화면, 세션 내에서 탭처럼 사용가능
+ pane: 하나의 윈도우 내에서의 화면 분활
## sesstion
>tmux new -s (sesstion_name)
>tmux new -s (sesstion_name -n (window_name)
>exit
>tmux ls
>tmux -attach -t sesstion_number
>(ctrl + b) d
// 스크롤 하기
>(ctrl + b + [
tmux kill_sesstion -t sesstion_number

## window
>(ctrl + b) c
>(ctrl + b) ( 숫자)

## pane
>(ctrl + b) % # 좌우로 나누기
>(ctrl + b) " # 위아래로 나누기
>(ctrl + b) 방향키
>(ctrl + b) q
>(ctrl + b) o # 순서대로 이동
>(ctrl + d) # 삭제
>(ctrl + b) : resize_pane -L 10 # L,R,U,D 입력하면 상하좌우로 조절
>(ctrl + b) # 사이즈 조절
>(ctrl + b) ? # 단축키 목록 
