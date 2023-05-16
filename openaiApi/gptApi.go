package openaiApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonBody map[string]interface{}

// FetchGPT returns a jsonBody object with the GPT-3 response to the given prompt
func FetchGPT(prompt string) jsonBody {
	apiKey := "KEY"
	model := "davinci"
	maxTokens := 50

	data := jsonBody{
		"model":      model,
		"prompt":     prompt,
		"max_tokens": maxTokens,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var response jsonBody
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	return response
}
