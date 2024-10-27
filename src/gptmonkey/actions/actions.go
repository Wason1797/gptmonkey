package actions

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Wason1797/gptmonkey/animation"
	"github.com/Wason1797/gptmonkey/configs"
	"github.com/Wason1797/gptmonkey/ollama"
	"github.com/Wason1797/gptmonkey/text"
	"github.com/urfave/cli/v2"
)

func MainAction(cCtx *cli.Context) error {
	// 1. Check for configs
	config_map := configs.GetConfigs()

	// 2. Ask for a url to query
	configs_changed := configs.InitBaseConfigs(config_map)

	// 3. Store said url in a config file if something was changed
	if configs_changed {
		configs.SaveConfigs(config_map)
	}

	// 4. Parse Input
	cmd_args := cCtx.Args()

	prompt := ""

	if cmd_args.Len() > 1 {
		prompt = strings.Join(cmd_args.Slice(), " ")
	}

	if cmd_args.Len() == 1 {
		prompt = cmd_args.First()
	}

	if cmd_args.Len() == 0 {
		log.Fatal("Please provide a promt for the model")
	}

	// 5. Add animations while waiting :^)
	monkey := animation.NewAnimation([]string{"🙈", "🙉", "🙊", "🐵", "🐒"})
	monkey.Init()
	response_ch := make(chan []ollama.ModelResponse)

	if config_map.OutputMode() == text.ONLY_CODE {
		prompt += ". only code or commands, no explanation"
	}

	// 6. Query Model
	go ollama.GetModelResponse(config_map[configs.CODELLAMA_URL], prompt, response_ch)

	for {
		select {
		case rsp := <-response_ch:
			monkey.End()
			formatter := text.GetResponseFormatter(config_map.OutputMode())
			result := formatter(text.ModelResponseToText(rsp))
			fmt.Print(result)
			return nil
		case <-time.After(time.Second / 2):
			monkey.Animate()
		}
	}

}
