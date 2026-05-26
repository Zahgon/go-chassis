package lager

import (
	"github.com/go-chassis/seclog/third_party/forked/cloudfoundry/lager"
)

// constant values for log rotate parameters
const (
	LogRotateDate  = 1
	LogRotateSize  = 10
	LogBackupCount = 7
)

// log level
const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
)

// output type
const (
	Stdout = "stdout"
	File   = "file"
)

// logFilePath log file path
var logFilePath string

// Options is the struct for lager information(lager.yaml)
type Options struct {
	Writers       string `yaml:"logWriters"`
	LoggerLevel   string `yaml:"logLevel"`
	LoggerFile    string `yaml:"logFile"`
	LogFormatText bool   `yaml:"logFormatText"`
	LogColorMode  string `yaml:"logColorMode"`

	LogRotateDisable  bool `yaml:"logRotateDisable"`
	LogRotateCompress bool `yaml:"logRotateCompress"`
	LogRotateAge      int  `yaml:"logRotateAge"`
	LogRotateSize     int  `yaml:"logRotateSize"`
	LogBackupCount    int  `yaml:"logBackupCount"`

	AccessLogFile string `yaml:"accessLogFile"`
}

// Init Build constructs a *Lager.logger with the configured parameters.
func Init(option *Options) { _ = "STUB: not implemented"; return }

// NewLog returns a logger
func NewLog(option *Options) (lager.Logger, error) {
	_ = "STUB: not implemented"
	return *new(lager.Logger), nil
}

// checkPassLagerDefinition check pass lager definition
func checkPassLagerDefinition(option *Options) { _ = "STUB: not implemented"; return }

// createLogFile create log file
func createLogFile(localPath, out string) error { _ = "STUB: not implemented"; return nil }
