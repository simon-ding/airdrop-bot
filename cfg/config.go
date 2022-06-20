package cfg

import (
	"airdrop-bot/log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	ChromeModXvfb = "xvfb"
)

type Config struct {
	WalletPassword   string     `mapstructure:"walletPassword"`
	AccountsPerIp    int        `mapstructure:"accountsPerIp"`
	AccountsToGen    int        `mapstructure:"accountsToGen"`
	RunSingleStep    bool       `mapstructure:"runSingleStep"`
	Owlracle         Owlracle   `mapstructure:"owlracle"`
	Dir              string     `mapstructure:"dir"`
	ChromeMode       string     `mapstructure:"chromeMode"`
	GasFeeAcceptable int        `mapstructure:"gasFeeAcceptable"`
	Cloudflare       Cloudflare `mapstructure:"cloudflare"`
	AWS              AWS        `mapstructure:"aws"`
	Binance          Binance    `mapstructure:"binance"`
}

func (c *Config) XvfbMod() bool {
	return c.ChromeMode == ChromeModXvfb
}

type AWS struct {
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	InstanceName    string `mapstructure:"instanceName"`
	Region          string `mapstructure:"region"`
}

type Cloudflare struct {
	ApiKey    string `mapstructure:"apiKey"`
	ZoneId    string `mapstructure:"zoneId"`
	KvId      string `mapstructure:"kvId"`
	AccountId string `mapstructure:"accountId"`
	URL       string `mapstructure:"url"`
}

type Owlracle struct {
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
	Disable   bool   `mapstructure:"disable"`
}

type Binance struct {
	ApiKey    string `mapstructure:"apiKey"`
	SecretKey string `mapstructure:"secretKey"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.SetDefault("walletPassword", "")
	viper.SetDefault("accountsPerIp", 3)
	viper.SetDefault("dir", ".")
	viper.SetDefault("gasFeeAcceptable", 7)

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
