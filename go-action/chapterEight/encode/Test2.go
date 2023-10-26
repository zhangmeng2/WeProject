package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `"json:name"`
	Title   string `"json:title"`
	Contact struct {
		Home string `"json:home"`
		Cell string `"json:cell"`
	} `"json:contact"`
}

var JSON = `{
	"name":"Gopher",
	"title":"programmer",
	"contact":{
		"home":"444.333.444",
		"cell":"333.222.111"
}
}`

func main44() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(c)
}
