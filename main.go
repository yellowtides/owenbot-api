package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/yellowtides/owenbot-api/utils"
)

func main() {
	text, err := ioutil.ReadFile("assets/nietzsche.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(utils.Owoify(string(text)))
}