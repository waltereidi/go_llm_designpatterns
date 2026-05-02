package interfaces

type IPromptStrategy interface {
	Start(body []byte) (string, error)
}
