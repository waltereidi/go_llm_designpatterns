package llm

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Response struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func Ask(prompt string) (string, error) {
	client := resty.New()

	reqBody := Request{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY")).
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return "", err
	}

	var result Response
	json.Unmarshal(resp.Body(), &result)

	return result.Choices[0].Message.Content, nil
}
