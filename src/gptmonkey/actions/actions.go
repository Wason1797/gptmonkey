package actions

import (
	"log"
	"strings"
	"time"

	"github.com/Wason1797/gptmonkey/configs"
	"github.com/Wason1797/gptmonkey/text"
	"github.com/Wason1797/gptmonkey/utils"
	"github.com/urfave/cli/v2"
)

func initBaseConfigs(config_map configs.ConfigMap) bool {
	if _, ok := config_map[configs.CODELLAMA_URL]; !ok {

		olama_url := utils.ReadInput("Input your olama url")
		if len(olama_url) > 0 {
			config_map[configs.CODELLAMA_URL] = olama_url
			return true
		} else {
			log.Fatal("Invalid url entered")
		}

	}
	return false
}

func MainAction(cCtx *cli.Context) error {
	// 1. Check for configs
	config_map := configs.GetConfigs()

	// 2. Ask for a url to query
	configs_changed := initBaseConfigs(config_map)

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
	animation := text.NewAnimation([]string{"🌎", "🌍", "🌏"})
	animation.Init()
	response_ch := make(chan []ModelResponse)

	// 6. Query Model
	go GetModelResponse(config_map[configs.CODELLAMA_URL], prompt, response_ch)

	for {
		select {
		case rsp := <-response_ch:
			PrintModelResponse(rsp)
			return nil
		case <-time.After(time.Second / 2):
			animation.Animate()
		}
	}

}
