package handlers

import (
	"encoding/json"
	"go_llm_designpatterns/llm"
	"log"
	"strings"
)

type TextFormattingCommand struct{}

func GetTemplateText(str string) string {
	result := "Reestruture o texto fornecido com elegância, clareza e boa organização. Preserve integralmente o significado original, mas melhore a apresentação por meio de uma estrutura refinada, com fluidez, coesão e legibilidade. Ajuste a formatação para destacar ideias principais, criar uma hierarquia visual harmoniosa e tornar a leitura mais agradável. Evite redundâncias, mantenha um tom profissional e entregue o resultado como um texto bem estruturado e esteticamente equilibrado. "
	result += str
	return result
}

func (c *TextFormattingCommand) Execute(body []byte) error {
	var data struct {
		Data string `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	
	formatedString := strings.ReplaceAll(data.Data, "\r", "")
	result, err := llm.Ask(GetTemplateText(formatedString))
	if err != nil {
		return err
	}
	log.Printf(result)

	log.Printf("👤 Criando usuário: %s", data.Data)
	return nil
}
