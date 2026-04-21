package interfaces

type CommandPattern interface {
	Execute(body []byte) error
}
