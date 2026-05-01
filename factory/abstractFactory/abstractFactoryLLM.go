package abstractFactory

import (
	"fmt"
	"go_llm_designpatterns/factory"
	"go_llm_designpatterns/interfaces"
)

func GetLLMFactory(operation string) (interfaces.IPromptFactory, error) {
	if operation == "text_formatting" {
		return &factory.TextFormattingFactory{}, nil
	}

	return nil, fmt.Errorf("invalid operation: %s", operation)
}
