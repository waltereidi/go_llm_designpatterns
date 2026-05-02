package factory

import (
	"go_llm_designpatterns/handlers"
	"go_llm_designpatterns/interfaces"
)

type SintaxValidationFactory struct{}

func (svf *SintaxValidationFactory) CreateStrategy() interfaces.IPromptStrategy {
	return &handlers.SintaxValidationStrategy{}
}
