package main

import (
	"bufio"
	"fmt"
	"os"
)

type Card struct {
	gaui int
	baui int
	bo   int
}

type GameCondition struct {
	setnum    int
	setcoount int
}

func main() {

	var f *os.File
	f = os.Stdin
	defer f.Close()

	// 초기화
	myplayer := Card{3, 3, 3}
	complayer := Card{3, 3, 3}
	gamecount := GameCondition{0, 0}
	// 메뉴
	// 게임 시작
	fmt.Println("경기 시작")

	for gamecount.setnum < 5 {
		for gamecount.setcoount < 5 {
			fmt.Printf("%d 세트 %d 시합입니다.\n", gamecount.setnum+1, gamecount.setcoount+1)
			fmt.Printf("당신은 가위 %d개 바위 %d개 보 %d개 가 있습니다.\n", myplayer.gaui, myplayer.baui, myplayer.bo)
			fmt.Printf("상대는 가위 %d개 바위 %d개 보 %d개 가 있습니다.\n", complayer.gaui, complayer.baui, complayer.bo)
			fmt.Println("무엇을 낼까요? 1.가위 2.바위 3.보")
			fmt.Print(">")
			scanner := bufio.NewScanner(f)
			scanner.Scan()
			if scanner.Text() == "가위" {
				if myplayer.gaui <= 0 {
					fmt.Println("가위를 낼수 없습니다.")
				} else {
					myplayer.gaui = myplayer.gaui - 1
					fmt.Println(">", scanner.Text())
					gamecount.setcoount++
				}
			} else if scanner.Text() == "바위" {
				if myplayer.baui <= 0 {
					fmt.Println("바위를 낼수 없습니다.")
				} else {
					myplayer.baui = myplayer.baui - 1
					fmt.Println(">", scanner.Text())
					gamecount.setcoount++
				}
			} else if scanner.Text() == "보" {
				if myplayer.bo <= 0 {
					fmt.Println("보를 낼수 없습니다.")
				} else {
					myplayer.bo = myplayer.bo - 1
					fmt.Println(">", scanner.Text())
					gamecount.setcoount++
				}
			} else if scanner.Text() == "끝" {
				return
			} else {
				fmt.Println("잘 못 된 입 력 입 니 다.")
			}
		}
		gamecount.setnum++
		gamecount.setcoount = 0
		myplayer.gaui = myplayer.gaui + 5
	}
}
