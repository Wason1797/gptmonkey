package ollama

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type ModelResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func GetModelResponse(url string, prompt string, c chan []ModelResponse) {

	body_map := map[string]string{
		"model":  "codellama",
		"prompt": prompt,
	}
	request_body, _ := json.Marshal(body_map)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(request_body))

	if err != nil {
		c <- make([]ModelResponse, 0)
		log.Print(err)
	}

	response_body, _ := io.ReadAll(resp.Body)

	resp.Body.Close()

	response_lines := strings.Split(string(response_body), "\n")
	var response_slice []ModelResponse

	for _, line := range response_lines {
		if len(line) == 0 {
			continue
		}
		parsed_line := ModelResponse{}
		if err := json.Unmarshal([]byte(line), &parsed_line); err != nil {
			log.Print("Could not parse line: ", line)
		}
		response_slice = append(response_slice, parsed_line)
	}

	c <- response_slice
}
