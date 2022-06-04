package cfg

import (
	"airdrop-bot/log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	WalletPassword string   `mapstructure:"walletPassword"`
	AccountsPerIp  int      `mapstructure:"accountsPerIp"`
	Owlracle       Owlracle `mapstructure:"owlracle"`
}

type Owlracle struct {
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.SetDefault("walletPassword", "")
	viper.SetDefault("accountsPerIp", 3)

	// optionally look for config in the working directory
	var cfg Config

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("create config file")
			viper.SafeWriteConfig()
		} else {
			// Config file was found but another error was produced
		}

		return nil, errors.Wrap(err, "load config")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "unmarshal file")
	}

	return &cfg, nil
}
