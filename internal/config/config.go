package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	App struct {
		Name string `json:"name"`
	} `json:"app"`
	Web struct {
		Prefork bool `json:"prefork"`
		Port    int  `json:"port"`
	} `json:"web"`
	Log struct {
		Level int `json:"level"`
	} `json:"log"`
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		SSLMode  string `json:"sslmode"`
		TimeZone string `json:"timezone"`
	} `json:"database"`
}

var AppConfig Config

func LoadConfig() {
	// Log the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Cannot get current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", wd)

	// Open the config file
	file, err := os.Open("../../config.json")
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Cannot close config file: %v", err)
		}
	}()

	// Decode the config file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Cannot decode config JSON: %v", err)
	}

	log.Println("Configuration loaded successfully")
}
