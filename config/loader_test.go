package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dhanar-kusuma/go-spark/config"
	"github.com/stretchr/testify/assert"
)

type SampleConfig struct {
	AppName string `envconfig:"APP_NAME" mapstructure:"app_name"`
	Port    int    `envconfig:"PORT"     mapstructure:"port"`
}

func TestLoadConfig(t *testing.T) {
	t.Run("load config with default config (.env)", func(t *testing.T) {
		_ = os.Setenv("APP_NAME", "TestApp")
		_ = os.Setenv("PORT", "8080")

		loader := config.NewLoader(nil)
		var cfg SampleConfig

		envPath := filepath.Join("testdata", ".env")

		err := loader.SetConfigPath(envPath).Load(&cfg)
		assert.NoError(t, err)
		assert.Equal(t, "TestApp", cfg.AppName)
		assert.Equal(t, 8080, cfg.Port)

		_ = os.Unsetenv("APP_NAME")
		_ = os.Unsetenv("PORT")
	})

	t.Run("load config from .yaml config", func(t *testing.T) {
		yamlConfig := filepath.Join("testdata", "config.yaml")

		var cfg SampleConfig
		err := config.NewLoader(nil).SetConfigPath(yamlConfig).Load(&cfg)
		assert.NoError(t, err)
		assert.Equal(t, "ExampleApp", cfg.AppName)
		assert.Equal(t, 9090, cfg.Port)
	})

	t.Run("load config from json config", func(t *testing.T) {
		jsonConfig := filepath.Join("testdata", "config.json")

		var cfg SampleConfig
		err := config.NewLoader(nil).SetConfigPath(jsonConfig).Load(&cfg)
		assert.NoError(t, err)
		assert.Equal(t, "example_app", cfg.AppName)
		assert.Equal(t, 7777, cfg.Port)
	})

	t.Run("load config from toml config", func(t *testing.T) {
		tomlConfig := filepath.Join("testdata", "config.toml")

		var cfg SampleConfig
		err := config.NewLoader(nil).SetConfigPath(tomlConfig).Load(&cfg)
		assert.NoError(t, err)
		assert.Equal(t, "example_app_toml", cfg.AppName)
		assert.Equal(t, 8989, cfg.Port)
	})

	t.Run("load config and override config using environment variable", func(t *testing.T) {
		_ = os.Setenv("APP_NAME", "app_name_override_envvar")
		_ = os.Setenv("PORT", "1111")

		yamlConfig := filepath.Join("testdata", "config.yaml")

		var cfg SampleConfig
		err := config.NewLoader(nil).SetConfigPath(yamlConfig).Load(&cfg)
		assert.NoError(t, err)
		assert.Equal(t, "app_name_override_envvar", cfg.AppName)
		assert.Equal(t, 1111, cfg.Port)

		_ = os.Unsetenv("APP_NAME")
		_ = os.Unsetenv("PORT")
	})

	t.Run("load config non valid format then return err", func(t *testing.T) {
		csvConfig := filepath.Join("testdata", "config.csv")

		var cfg SampleConfig
		err := config.NewLoader(nil).SetConfigPath(csvConfig).Load(&cfg)
		assert.Error(t, err)
		assert.EqualError(t, err, config.ErrUnsupportedConfigType.Error())
	})
}
