package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	sampleData := `
	{
		"data": {
			"id":"11230495409546",
			"title":"Atari Pac Man",
			"year":1984,
			"point": 7.6,
			"object":"Arcade"
		}
	}
	`

	var m map[string]map[string]interface {
	}

	err := json.Unmarshal([]byte(sampleData), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	fmt.Println("************")
	content, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
