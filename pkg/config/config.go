package config

import "github.com/spf13/viper"

const DefaultConfigFilePath = "/home/zelta/img-server/conf/app.yaml"

type Config struct {
	Enabled        bool `yaml:"enabled"`
	Port           string
	Url            string
	ImgRepo        string
	DB             Database
	ConfigFilePath string
}

type Database struct {
	Driver   string
	Name     string
	Port     string
	Username string
	Password string
}

func Default() *Config {
	return &Config{}
}

func (cfg *Config) SetConfigFilePath() *Config {
	viper.SetConfigFile(DefaultConfigFilePath)
	return cfg
}

func (cfg *Config) LoadConfigFile() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) InitImgServer() {
	cfg.Enabled = viper.GetBool("img-server.enabled")
	cfg.Url = viper.GetString("img-server.url")
}
