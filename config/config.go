package config

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	Host  string
	Port  string
	Nouns []string
}

func NewConfig() Config {
	return Config{
		Host:  getHost(),
		Port:  getPort(),
		Nouns: getNouns(),
	}
}

func getHost() string {
	host := "0.0.0.0:"
	if value, ok := os.LookupEnv("HOST"); ok {
		host = value
	}
	return host
}

func getPort() string {
	port := "8081"
	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}
	return port
}

func getNouns() []string {
	nouns := []string{}

	// load it from file
	content, err := os.ReadFile("russian_nouns.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Split the content by new lines and add each noun to the slice
	for _, noun := range strings.Split(string(content), "\n") {
		if noun != "" { // This check prevents adding empty strings if there are blank lines
			nouns = append(nouns, noun)
		}
	}

	return nouns
}
