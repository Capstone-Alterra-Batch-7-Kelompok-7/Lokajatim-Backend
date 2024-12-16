package config

import "os"

type Config struct {
	GeminiAIKey string
}

func LoadConfigGemini() Config {
	return Config{
		GeminiAIKey: os.Getenv("GEMINI_API_KEY"), 
	}
}
