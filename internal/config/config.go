package config

import (
	"os"
	"time"
)

const defaultKey = "RgkH0m_iWfVsoSuLheI46BAYy1Ig1Ab3vvElYAESZ34"

// Config holds application configuration
type Config struct {
	// Unsplash API access key
	UnsplashAccessKey string

	// API request timeout
	RequestTimeout time.Duration
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	accessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
	if accessKey == "" {
		accessKey = defaultKey
	}
	return &Config{UnsplashAccessKey: accessKey, RequestTimeout: 10 * time.Second}, nil
}
