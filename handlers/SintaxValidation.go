package handlers

import (
	"encoding/json"
	"log"
)

type SintaxValidationCommand struct{}

func (c *SintaxValidationCommand) Execute(body []byte) error {
	var data struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	log.Printf("👤 Criando usuário: %s", data.Name)
	return nil
}
