package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type configType string

func (c configType) String() string {
	return string(c)
}

const (
	yaml        configType = "yaml"
	json        configType = "json"
	toml        configType = "toml"
	env         configType = "env"
	unsupported configType = "unsupported"
)

const defaultEnvConfig = ".env"

var configTypes = map[string]configType{
	".yaml": yaml,
	".yml":  yaml,
	".json": json,
	".toml": toml,
	".env":  env,
}

type ConfigLoader struct {
	configPath string
	configType configType
	envPrefix  string

	defaultEnvConfig string
	logger           *slog.Logger
}

func NewLoader(logger *slog.Logger) *ConfigLoader {
	loader := &ConfigLoader{
		defaultEnvConfig: defaultEnvConfig,
		configPath:       defaultEnvConfig,
		configType:       env,
		envPrefix:        "",
	}
	if logger == nil {
		loader.logger = slog.Default()
	} else {
		loader.logger = logger
	}
	return loader
}

func (c *ConfigLoader) SetConfigPath(configPath string) *ConfigLoader {
	ext := filepath.Ext(configPath)
	cfgType, valid := configTypes[ext]
	if !valid {
		c.configType = unsupported
		return c
	}
	c.configType = cfgType
	c.configPath = configPath
	return c
}

func (c *ConfigLoader) SetEnvPrefix(prefix string) *ConfigLoader {
	c.envPrefix = prefix
	return c
}

func (c *ConfigLoader) SetDefaultEnv(defaultEnv string) *ConfigLoader {
	c.defaultEnvConfig = defaultEnv
	return c
}

func (c *ConfigLoader) Load(cfg any) error {
	if c.configType == unsupported {
		return ErrUnsupportedConfigType
	}

	if c.configType == env {
		err := envconfig.Process(c.envPrefix, cfg)
		if err != nil {
			return err
		}
		return nil
	}

	return c.configureViperConfig(cfg)
}

func (c *ConfigLoader) setDefaultEnvConfig(path string) {
	if _, err := os.Stat(path); err == nil {
		err := godotenv.Load(path)
		if err != nil {
			c.logger.
				With("err", err, "path", path).
				Error("failed to load env variables from default env config path")
		}
	} else {
		c.logger.
			With("path", path).
			Debug("default env config path not found, skipping load")
	}
}

func (c *ConfigLoader) configureViperConfig(cfg any) error {
	c.setDefaultEnvConfig(c.defaultEnvConfig)
	if c.envPrefix != "" {
		viper.SetEnvPrefix(c.envPrefix)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigType(c.configType.String())
	viper.SetConfigFile(c.configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(cfg)
}
