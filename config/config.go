package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

	"gopkg.in/yaml.v2"
)

type LimiterConfig struct {
	Enabled           bool    `yaml:"enabled"`
	RequestsPerSecond float64 `yaml:"requests_per_second"`
}

type LoggerTypeConfig struct {
	File     string `yaml:"file"`
	Live     bool   `yaml:"live"`
	JsonDump bool   `yaml:"json_dump"`
}

type HeaderConfig struct {
	Header  string `yaml:"header"`
	Content string `yaml:"content"`
}

type WorkerConfig struct {
	Limiter     LimiterConfig `yaml:"limiter"`
	WordLists   []string      `yaml:"wordlists"`
	Append      []string      `yaml:"append"`
	IgnoreCodes []int         `yaml:"ignore"`
	Threads     int           `yaml:"max_threads"`
}

type LoggingConfig struct {
	Stats         bool             `yaml:"stats"`
	DebugLogger   LoggerTypeConfig `yaml:"debug_logger"`
	SuccessLogger LoggerTypeConfig `yaml:"success_logger"`
	ErrorLogger   LoggerTypeConfig `yaml:"error_logger"`
}

type RequestConfig struct {
	Cookie  string         `yaml:"cookie"`
	Headers []HeaderConfig `yaml:"headers"`
}

type AppConfig struct {
	WorkerConfig  WorkerConfig  `yaml:"worker"`
	LoggingConfig LoggingConfig `yaml:"logging"`
	RequestConfig RequestConfig `yaml:"requests"`
}

func LoadConfig(profile string) *AppConfig {
	_, err := user.Current()
	if err != nil {
		log.Fatalf("Could not find user home directory")
	}
	// data, err := ReadConfigFile(fmt.Sprintf("%s/.godirb/%s.yaml", user.HomeDir, profile))
	data, err := ReadConfigFile(fmt.Sprintf("./config/%s.yaml", profile))
	if err != nil {
		log.Fatalf("Could not read config file for profile: %s - %s", profile, err)
	}
	appConfig, err := ParseYamlConfig(data)
	if err != nil {
		log.Fatalf("Could not parse config for profile: %s", profile)
	}
	return appConfig
}

func ReadConfigFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseYamlConfig(yamlConfig []byte) (*AppConfig, error) {
	config := &AppConfig{}
	if err := yaml.Unmarshal(yamlConfig, config); err != nil {
		return &AppConfig{}, err
	}
	return config, nil
}
