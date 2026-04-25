package handlers

import (
	"encoding/json"
	"log"
	"go_llm_designpatterns/llm"
)

type SintaxValidationCommand struct{}

func (c *SintaxValidationCommand) Execute(body []byte) error {
	var data struct {
		Data string `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	llm.Ask("Verifique se esta frase está sintaticamente correta: " + data.Data)

	log.Printf("👤 Criando usuário: %s", data.Data)
	return nil
}
