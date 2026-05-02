package handlers

import (
	"encoding/json"
	"go_llm_designpatterns/llm"
	"log"
)

type SintaxValidationStrategy struct {
}

func (s *SintaxValidationStrategy) Start(body []byte) (string, error) {
	var payload struct {
		Data string `json:"data"`
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return "", err
	}

	response, err := llm.Ask(
		"Verifique se esta frase está sintaticamente correta: " + payload.Data,
	)
	if err != nil {
		return "", err
	}

	log.Printf("🔎 Validação concluída: %s", response)

	return response, nil
}
