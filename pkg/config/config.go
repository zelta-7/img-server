package config

import "github.com/spf13/viper"

const DefaultConfigFilePath = "/home/zelta/img-server/conf/app.yaml"

type Config struct {
	Enabled        bool     `yaml:"enabled"`
	Port           string   `yaml:"port"`
	Url            string   `yaml:"url"`
	ImgRepo        string   `yaml:"imgrepo"`
	DB             Database `yaml:"db"`
	ConfigFilePath string   `yaml:"configfilepath"`
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
	cfg.Port = viper.GetString("img-server.port")
	cfg.Url = viper.GetString("img-server.url")
	cfg.ImgRepo = viper.GetString("img-server.img-repo")
	cfg.DB.Driver = viper.GetString("img-server.database.driver")
	cfg.DB.Name = viper.GetString("img-server.database.name")
	cfg.DB.Port = viper.GetString("img-server.database.port")
	cfg.DB.Username = viper.GetString("img-server.database.username")
	cfg.DB.Password = viper.GetString("img-server.database.password")
}
