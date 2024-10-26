package configs

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strings"

	"github.com/Wason1797/gptmonkey/utils"
)

const DEFAULT_CONFIG_FILE = ".gptmonkey"
const COMMENT = "#"
const EQ = "="
const API_KEY = "API_KEY"
const CODELLAMA_URL = "CODELLAMA_URL"
const OUTPUT_MODE = "OUTPUT_MODE"

type ConfigMap map[string]string
type TokenAction func(ConfigMap, string)

type ConfigToken struct {
	token  string
	action TokenAction
}

func (cm ConfigMap) OutputMode() string {
	return cm[OUTPUT_MODE]
}

func (cm ConfigMap) CodellamaURL() (string, bool) {
	url, ok := cm[CODELLAMA_URL]
	return url, ok
}

func makeToken(token string, action TokenAction) ConfigToken {
	return ConfigToken{token, action}
}

func parseConfigLine(token string, line string) string {
	result, found := strings.CutPrefix(line, token+EQ)

	if !found {
		log.Fatal("Invalid configuration line: ", line)
	}

	return strings.ReplaceAll(result, " ", "")

}

var token_list = []ConfigToken{
	makeToken(COMMENT, func(ConfigMap, string) {}),
	makeToken(API_KEY, func(m ConfigMap, line string) {
		m[API_KEY] = parseConfigLine(API_KEY, line)
	}),
	makeToken(CODELLAMA_URL, func(m ConfigMap, line string) {
		m[CODELLAMA_URL] = parseConfigLine(CODELLAMA_URL, line)
	}),
	makeToken(OUTPUT_MODE, func(m ConfigMap, line string) {
		m[OUTPUT_MODE] = parseConfigLine(OUTPUT_MODE, line)
	}),
}

func getConfigFilePath() string {
	home_dir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return path.Join(home_dir, DEFAULT_CONFIG_FILE)

}

/*
Function to get all configuration values from the default file
*/
func GetConfigs() ConfigMap {

	config_path := getConfigFilePath()

	var config_file *os.File = nil
	if _, err := os.Stat(config_path); errors.Is(err, os.ErrNotExist) {
		config_file, err = os.Create(config_path)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		config_file, err = os.Open(config_path)

		if err != nil {
			log.Fatal(err)
		}
	}

	if config_file == nil {
		log.Fatal("Could not open config file in ", config_path)
	}

	defer config_file.Close()

	configs := make(ConfigMap)

	scanner := bufio.NewScanner(config_file)

	for scanner.Scan() {
		line := scanner.Text()

		for _, token := range token_list {
			if strings.HasPrefix(line, token.token) {
				token.action(configs, line)
			}
		}

	}

	return configs
}

func SaveConfigs(configs ConfigMap) {
	config_file, err := os.OpenFile(getConfigFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer config_file.Close()
	for config_key, config_value := range configs {
		config_file.WriteString(config_key + EQ + config_value + "\n")
	}
}

func InitBaseConfigs(config_map ConfigMap) bool {
	if _, ok := config_map.CodellamaURL(); !ok {

		olama_url := utils.ReadInput("Input your olama url")
		if len(olama_url) > 0 {
			config_map[CODELLAMA_URL] = olama_url
			return true
		} else {
			log.Fatal("Invalid url entered")
		}

	}
	return false
}
