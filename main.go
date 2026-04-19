package main

import (
	"fmt"
	"go_llm_designpatterns/llm"
)

func main() {
	resp, err := llm.Ask("diga olá mundo em ingles")
	if err != nil {
		panic(err)
	}

	fmt.Println("Resposta:", resp)
}
