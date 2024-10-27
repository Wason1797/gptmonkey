package text

import (
	"bytes"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/Wason1797/gptmonkey/ollama"
)

const ONLY_CODE = "ONLY_CODE"
const FULL = "FULL"

type FormatterFunction = func(string) string

var markdownFormatter FormatterFunction = func(model_result string) string {
	return string(markdown.Render(model_result, 80, 6))
}

var onlyCodeFormatter FormatterFunction = func(model_result string) string {
	var buffer bytes.Buffer

	first_index := strings.Index(model_result, "```")

	if first_index == -1 {
		return model_result
	}
	for {
		new_index := strings.Index(model_result[first_index+3:], "```")
		if new_index == -1 {
			break
		}
		buffer.WriteString(model_result[first_index : first_index+new_index+6])
		first_index = (first_index + new_index + 3)
		if first_index > len(model_result) {
			break
		}
	}

	return buffer.String()

}

var rawFormatter FormatterFunction = func(model_result string) string {
	return model_result
}

func GetResponseFormatter(config_option string) FormatterFunction {
	switch config_option {

	case ONLY_CODE:
		return onlyCodeFormatter
	case FULL:
		return markdownFormatter
	default:
		return rawFormatter
	}
}

func ModelResponseToText(response_slice []ollama.ModelResponse) string {

	var buffer bytes.Buffer

	for _, response := range response_slice {
		if !response.Done {
			buffer.WriteString(response.Response)
		}
	}
	return buffer.String()
}
