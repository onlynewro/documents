package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 맵을 JSON 문자열로 마샬링한다.
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println("에러:", err)
		return
	}

	fmt.Println(string(data))
}
