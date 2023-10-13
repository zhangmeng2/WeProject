package main

import (
	_ "WeProject/go-action/chapterTwo/matchers"
	"WeProject/go-action/chapterTwo/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
