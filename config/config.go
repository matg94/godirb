package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

type AppFlags struct {
	URL      string
	Profile  string
	Limiter  float64
	Threads  int
	Wordlist string
	Cookie   string
	JsonPipe bool
	OutFile  string
	Stats    bool
	Silent   bool
}

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
	AppendOnly  bool          `yaml:"append_only"`
	Append      []string      `yaml:"append"`
	IgnoreCodes []int         `yaml:"ignore"`
	Threads     int           `yaml:"threads"`
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
	data, err := ReadConfigFile(fmt.Sprintf("%s/.godirb/%s.yaml", os.Getenv("HOME"), profile))
	if err != nil {
		log.Fatalf("Could not read config file for profile: %s - %s", profile, err)
	}
	appConfig, err := ParseYamlConfig(data)
	if err != nil {
		log.Fatalf("Could not parse config for profile: %s", profile)
	}
	return appConfig
}

func LoadConfigWithFlags(flags AppFlags) *AppConfig {
	profileConfig := LoadConfig(flags.Profile)
	if flags.Limiter != -1 {
		profileConfig.WorkerConfig.Limiter.RequestsPerSecond = float64(flags.Limiter)
		profileConfig.WorkerConfig.Limiter.Enabled = true
	}
	if flags.Threads != -1 {
		profileConfig.WorkerConfig.Threads = flags.Threads
	}
	if flags.Wordlist != "" {
		profileConfig.WorkerConfig.WordLists = []string{flags.Wordlist}
	}
	if flags.Cookie != "" {
		profileConfig.RequestConfig.Cookie = flags.Cookie
	}
	if flags.JsonPipe {
		profileConfig.LoggingConfig.Stats = false
		profileConfig.LoggingConfig.DebugLogger.Live = false
		profileConfig.LoggingConfig.ErrorLogger.Live = false
		profileConfig.LoggingConfig.SuccessLogger.Live = false
		profileConfig.LoggingConfig.SuccessLogger.JsonDump = true
	}
	if flags.OutFile != "" {
		profileConfig.LoggingConfig.SuccessLogger.File = flags.OutFile
	}
	if flags.Stats {
		profileConfig.LoggingConfig.Stats = true
	}
	if flags.Silent {
		profileConfig.LoggingConfig.DebugLogger.Live = false
		profileConfig.LoggingConfig.ErrorLogger.Live = false
		profileConfig.LoggingConfig.SuccessLogger.Live = false
	}
	return LoadRequiredDefaults(profileConfig)
}

func LoadRequiredDefaults(config *AppConfig) *AppConfig {
	if config.WorkerConfig.Threads <= 0 {
		config.WorkerConfig.Threads = 1
	}
	return config
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
