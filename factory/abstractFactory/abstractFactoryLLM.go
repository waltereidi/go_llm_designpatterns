package abstractFactory

import (
	"fmt"
	"go_llm_designpatterns/factory"
	"go_llm_designpatterns/interfaces"
	"go_llm_designpatterns/models"
)

type AbstractFactoryLLM struct{}

func (af *AbstractFactoryLLM) GetLLMFactory(msg *[]byte) (interfaces.IPromptFactory, error) {
	var message models.Message
	message.MessageToModel(*msg)
	result, err := af.CommandGetFactory(&message.Type)

	return result, err
}

func (cgf *AbstractFactoryLLM) CommandGetFactory(op *string) (interfaces.IPromptFactory, error) {

	switch *op {
	case "text_formatting":
		return &factory.TextFormattingFactory{}, nil
	case "tag_extraction":
		return &factory.TagExtractionFactory{}, nil
	case "syntax_validation":
		return &factory.SintaxValidationFactory{}, nil
	default:
		return nil, fmt.Errorf("invalid operation: %s", *op)
	}

}
