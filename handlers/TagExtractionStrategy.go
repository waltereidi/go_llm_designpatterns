package handlers

import (
	"encoding/json"
	"go_llm_designpatterns/llm"
	"log"
	"strings"
)

type TagExtractionStrategy struct{}

func (c *TagExtractionStrategy) Start(body []byte) (string, error) {
	var data struct {
		Data string `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	formatedString := strings.ReplaceAll(data.Data, "\r", "")
	result, err := llm.Ask("Extract key words from this phrase: " + formatedString)
	if err != nil {
		return "", err
	}
	log.Printf(result)

	log.Printf("👤 Criando usuário: %s", data.Data)
	return result, nil
}
