package apigatewaylambdagoexample

import (
	"errors"
	"log"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap/zapcore"
)

// ErrLogLevelConfigParse LogLevelConfig parse error.
var ErrLogLevelConfigParse = errors.New("unable to parse LogLevelConfig")

//Config represents apigatewaylambdagoexample configs
type Config struct {
	APP      string         `envconfig:"APP_NAME" default:"apigatewaylambdagoexample"`
	Port     int            `envconfig:"APP_PORT" default:"8080"`
	LogLevel LogLevelConfig `envconfig:"LOG_LEVEL" default:"info"`

	AWSRegion string `envconfig:"TPRTC_AWS_REGION" default:"us-east-1"`

	SQSEndpoint string `envconfig:"SQS_ENDPOINT"`
}

//NewConfig config constructor
func NewConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}

//Env environment value
type Env struct {
	Name       string
	Type       string
	DefaultVal string
	Required   bool
}

// LogLevelConfig log level config.
type LogLevelConfig struct {
	Value zapcore.Level
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The text is expected to have a valid zap level.
func (l *LogLevelConfig) UnmarshalText(text []byte) error {
	level := new(zapcore.Level)
	err := level.Set(string(text))
	l.Value = *level

	return err
}
