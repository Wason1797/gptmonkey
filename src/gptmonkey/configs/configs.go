package configs

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path"
	"strings"
)

const DEFAULT_CONFIG_FILE = ".gptmonkey"
const COMMENT = "#"
const EQ = "="
const API_KEY = "API_KEY"
const CODELLAMA_URL = "CODELLAMA_URL"
const ONLY_CODE = "ONLY_CODE"

type ConfigMap map[string]string
type TokenAction func(ConfigMap, string)

type ConfigToken struct {
	token  string
	action TokenAction
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
	makeToken(ONLY_CODE, func(m ConfigMap, line string) {
		m[ONLY_CODE] = parseConfigLine(ONLY_CODE, line)
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
