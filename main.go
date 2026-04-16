package main

import (
	"fmt"
	"go_llm_designpatterns/llm"
)

func main() {
	resp, err := llm.Ask("Explique Golang em uma frase")
	if err != nil {
		panic(err)
	}

	fmt.Println("Resposta:", resp)
}
