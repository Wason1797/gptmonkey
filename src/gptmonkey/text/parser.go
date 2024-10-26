package text

import (
	"bytes"

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
	return ""
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
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}
