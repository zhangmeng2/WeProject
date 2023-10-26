package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var JSON2 = `{
	"name":"Gopher",
	"title":"programmer",
	"contact":{
		"home":"444.333.444",
		"cell":"333.222.111"
}
}`

func main55() {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON2), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}
