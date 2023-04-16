package utils

import (
	"os"
	"sync"

	"github.com/sashabaranov/go-openai"
)

const IMAGE_SIZE = openai.CreateImageSize512x512
const PAGE_LIMIT = 48

type DBConfig struct {
	URI    string
	DBName string
}

type OpenAIConfig struct {
	ApiKey string
}

type Config struct {
	OpenAI *OpenAIConfig
}

var config *Config
var once sync.Once

func initConfig() {
	config = &Config{
		OpenAI: &OpenAIConfig{
			ApiKey: os.Getenv("OPENAI_API_KEY"),
		},
	}
}

func GetConfig() *Config {
	once.Do(initConfig)
	return config
}
