package group_app_service

import (
	"errors"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
	Probes ProbeConfig
}

type DBConfig struct {
	DBName   string `yaml:"dbname" env:"DBNAME"`
	Username string `yaml:"username" env:"DBUSERNAME"`
	Password string `yaml:"password" env:"DBPASSWORD"`
	Host     string `yaml:"host" env:"DBHOST"`
	Port     string `yaml:"port" env:"DBPORT"`
	SSLMode  string `yaml:"sslmode" env:"SSLMODE"`
}

type ServerConfig struct {
	Port   string `yaml:"port" env:"PORT"`
	Secret string `yaml:"secret_jwt" env:"SECRET_JWT"`
}

type ProbeConfig struct {
	Port int `env:"PROBES_PORT" envDefault:"3013"`
}

func (c *Config) Validate() error {
	if c.Server.Port == "" {
		return errors.New("empty server config")
	}

	if c.DB.Host == "" && c.DB.Port == "" && c.DB.DBName == "" && c.DB.Username == "" && c.DB.Password == "" {
		return errors.New("empty database config")
	}
	return nil
}

func LoadConfig(configPath, configName string) (*Config, error) {

	var cfg Config

	if err := env.Parse(&cfg.Probes); err != nil {
		logrus.Warnf("Input envs of PROBES: %d, use default port 3010", cfg.Probes)
	}

	if err := env.Parse(&cfg.DB); err != nil {
		logrus.Warnf("Input envs of DB: %+v", cfg.DB)
		logrus.Warnf("Failed to parse DB envs: %s", err.Error())
	}

	if err := env.Parse(&cfg.Server); err != nil {
		logrus.Warnf("Input envs of SERVER: %+v", cfg.Server)
		logrus.Warnf("Failed to parse SERVER envs: %s", err.Error())
	}

	if err := cfg.Validate(); err == nil {
		logrus.Infof("Use envs of PROBES: %d", cfg.Probes)
		logrus.Infof("Use envs of DB: %+v", cfg.DB)
		logrus.Infof("Use envs of SERVER: %+v", cfg.Server)
		return &cfg, nil
	}

	if err := initConfig(configPath, configName); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, errors.New("invalid data in config")
	}

	return &cfg, nil

}

func initConfig(configPath string, configName string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
