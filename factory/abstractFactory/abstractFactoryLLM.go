package abstractFactory

import (
	"fmt"
	"go_llm_designpatterns/factory"
	"go_llm_designpatterns/interfaces"
)

func GetLLMFactory(operation string) (interfaces.IPromptFactory, error) {

	switch operation {
	case "text_formatting":
		return &factory.TextFormattingFactory{}, nil
	case "tag_extraction":
		return &factory.TagExtractionFactory{}, nil
	case "syntax_validation":
		return &factory.SintaxValidationFactory{}, nil
	}

	return nil, fmt.Errorf("invalid operation: %s", operation)
}
