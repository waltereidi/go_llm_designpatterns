package factory

import (
	"go_llm_designpatterns/handlers"
	"go_llm_designpatterns/interfaces"
)

type TagExtractionFactory struct {
}

func (te *TagExtractionFactory) CreateStrategy() interfaces.IPromptStrategy {
	return &handlers.TagExtractionStrategy{}
}
