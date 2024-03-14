package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the configuration settings
type Config struct {
	ThresholdX float64 `mapstructure:"thresholdX"`
	// Add stellar xdr buying and selling asset specific settings here
	SellingAssetCode string `mapstructure:"sellingAssetCode"`
	SellingAssetIssuer string `mapstructure:"sellingAssetIssuer"`
	BuyingAssetCode string `mapstructure:"buyingAssetCode"`
	BuyingAssetIssuer string `mapstructure:"buyingAssetIssuer"`
	Amount int64 `mapstructure:"amount"`
}

// LoadConfig loads the configuration from the specified file
func LoadConfig(configFile string) error {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// Set default values
	viper.SetDefault("sellingAssetCode", "") // Example default value for thresholdX
	viper.SetDefault("sellingAssetIssuer", "") // Example default value for thresholdX
	viper.SetDefault("buyingAssetCode", "") // Example default value for thresholdX
	viper.SetDefault("buyingAssetIssuer", "") // Example default value for thresholdX

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Config file not found, using defaults: %v", err)
	}

	// Unmarshal the configuration into a struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	// TODO: validate the configuration

	// Set the global configuration
	GlobalConfig = config

	return nil
}

// GlobalConfig holds the global configuration settings
var GlobalConfig Config