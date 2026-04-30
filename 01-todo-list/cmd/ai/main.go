package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("AI > ")
	input, _ := reader.ReadString('\n')

	response := callLLM(input)

	fmt.Println("→", response)
}

func callLLM(prompt string) string {

	payload := map[string]any{
		"model": "local-model", // config trong LLM Studio
		"messages": []Message{
			{Role: "user", Content: prompt},
		},
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:1234/v1/chat/completions", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var result map[string]any
	json.NewDecoder(resp.Body).Decode(&result)

	return fmt.Sprintf("%v", result)
}
