package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 문자열을 가지고 있는 JSON 문서
var JSON = `{
	"name":		"Gopher",
	"title":	"programmer",
	"contact": {
		"home":		"415.333.3333",
		"cell":		"415.555.5555"
	}
}`

func main() {
	//JSON 문자열을 맵에 언마샬링한다.
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("에러:", err)
		return
	}

	fmt.Println("이름:", c["name"])

}
