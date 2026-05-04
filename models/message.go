package models

import "encoding/json"

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func (af *Message) MessageToModel(body []byte) error {

	if err := json.Unmarshal(body, &af); err != nil {
		return err
	}
	return nil
}
