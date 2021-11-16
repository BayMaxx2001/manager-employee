package config

import (
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var config Config
var mu sync.RWMutex

type Config struct {
	// server
	Server string
	//mysql
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_NAME     string `mapstructure:"DB_NAME"`

	// redis
	RedisServer   string `mapstructure:"RedisServer"`
	RedisPassword string `mapstructure:"RedisPassword"`
	RedisDB       int    `mapstructure:"RedisDB"`

	// mongo
	MongoDbUrl string `mapstructure:"MongoDbUrl"`
	Database   string `mapstructure:"Database"`
	Collection string `mapstructure:"Collection"`

	// log
	LogFile  string `mapstructure:"LogFile"`
	LOGLEVEL int    `mapstructure:"LOGLEVEL"`
}

func Get() Config {
	mu.RLock()
	defer mu.RUnlock()
	return config
}

func Set(c Config) {
	mu.Lock()
	defer mu.Unlock()
	config = c
}

func LoadEnvFromFile(configPrefix, envPath string) (err error) {
	mu.Lock()
	defer mu.Unlock()

	if err := godotenv.Load(envPath); err != nil {
		return err
	}
	return envconfig.Process(configPrefix, &config)
}

func ConfigZap() *zap.SugaredLogger {
	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.Level(Get().LOGLEVEL)),
		OutputPaths: []string{"stderr"},

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
			EncodeLevel:  CustomLevelEncoder,
			EncodeTime:   SyslogTimeEncoder,
		},
	}

	logger, _ := cfg.Build()
	return logger.Sugar()
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
