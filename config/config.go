package config

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	configuration *Configuration
	configFileExt = ".yml"
	configType    = "yaml"
)

// Configuration ...
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Email    EmailConfiguration
}

// ServerConfiguration is the required parameters to set up a server
type ServerConfiguration struct {
	Env                        string `default:"dev"` // dev, prod
	Port                       string `default:"3625"`
	Domain                     string `default:"API_WEBSITE_URL"`
	Dir                        string `default:"/app/config"`
	Passphrase                 string `default:"passphrase-for-encrypting-passwords-do-not-forget"`
	Secret                     string `default:"JWT_TOKEN`
	Timeout                    int    `default:"24"`
	GeneratedPasswordLength    int    `default:"16"`
	AccessTokenExpireDuration  string `default:"30m"`
	RefreshTokenExpireDuration string `default:"15d"`
	APIKey                     string `default:"SENDGRID_API_KEY"`
}

// DatabaseConfiguration is the required parameters to set up a DB instance
type DatabaseConfiguration struct {
	Name     string `default:"DATABASE_NAME"`
	Username string `default:"DATABASE_USERNAME"`
	Password string `default:"DATABASE_PASSWORD"`
	Host     string `default:"DATABASE_HOST_URL"`
	Port     string `default:"5777"`
	LogMode  bool   `default:"false"`
	SSLMode  string `default:"disable"`
}

// EmailConfiguration is the required parameters to send emails
type EmailConfiguration struct {
	Host     string `default:"smtp.gmail.com"`
	Port     string `default:"25"`
	Username string `default:"EMAIL"`
	From     string `default:"EMAIL"`
	Admin    string `default:"EMAIL"`
}

// Init ...
func Init(configPath, configName string) (*Configuration, error) {
	configFilePath := filepath.Join(configPath, configName) + configFileExt

	// initialize viper configuration
	initializeConfig(configPath, configName)

	// Bind environment variables
	bindEnvs()

	// Set default values
	setDefaults()

	// Read or create configuration file
	if err := readConfiguration(configFilePath); err != nil {
		return nil, err
	}

	// Auto read env variables
	viper.AutomaticEnv()

	// Unmarshal config file to struct
	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

// read configuration from file
func readConfiguration(configFilePath string) error {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		// if file does not exist, simply create one
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			os.Create(configFilePath)
		} else {
			return err
		}

		// let's write defaults
		if err := viper.WriteConfig(); err != nil {
			return err
		}
	}

	return nil
}

// initialize the configuration manager
func initializeConfig(configPath, configName string) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
}

func bindEnvs() {
	viper.BindEnv("server.env", "SM_ENV")
	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("server.domain", "DOMAIN")
	viper.BindEnv("server.passphrase", "SM_SERVER_PASSPHRASE")
	viper.BindEnv("server.secret", "SM_SERVER_SECRET")
	viper.BindEnv("server.timeout", "SM_SERVER_TIMEOUT")

	viper.BindEnv("server.generatedPasswordLength", "SM_SERVER_GENERATED_PASSWORD_LENGTH")
	viper.BindEnv("server.accessTokenExpireDuration", "SM_SERVER_ACCESS_TOKEN_EXPIRE_DURATION")
	viper.BindEnv("server.refreshTokenExpireDuration", "SM_SERVER_REFRESH_TOKEN_EXPIRE_DURATION")

	viper.BindEnv("server.apiKey", "SM_SERVER_API_KEY")
	viper.BindEnv("server.recaptcha", "SM_SERVER_RECAPTCHA")

	viper.BindEnv("database.name", "SM_DB_NAME")
	viper.BindEnv("database.username", "SM_DB_USERNAME")
	viper.BindEnv("database.password", "SM_DB_PASSWORD")
	viper.BindEnv("database.host", "SM_DB_HOST")
	viper.BindEnv("database.port", "SM_DB_PORT")
	viper.BindEnv("database.logmode", "SM_DB_LOG_MODE")

	viper.BindEnv("email.host", "SM_EMAIL_HOST")
	viper.BindEnv("email.port", "SM_EMAIL_PORT")
	viper.BindEnv("email.username", "SM_EMAIL_USERNAME")
	viper.BindEnv("email.password", "SM_EMAIL_PASSWORD")
	viper.BindEnv("email.fromEmail", "SM_EMAIL_FROM_EMAIL")
	viper.BindEnv("email.fromName", "SM_EMAIL_FROM_NAME")
	viper.BindEnv("email.apiKey", "SM_EMAIL_API_KEY")
}

func setDefaults() {
	// Server defaults
	viper.SetDefault("server.env", "production")
	viper.SetDefault("server.port", "3625")
	viper.SetDefault("server.domain", "API_WEBSITE_URL")
	viper.SetDefault("server.passphrase", generateKey())
	viper.SetDefault("server.secret", generateKey())
	viper.SetDefault("server.timeout", 24)
	viper.SetDefault("server.generatedPasswordLength", 16)
	viper.SetDefault("server.accessTokenExpireDuration", "30m")
	viper.SetDefault("server.refreshTokenExpireDuration", "15d")
	viper.SetDefault("server.apiKey", generateKey())
	viper.SetDefault("server.recaptcha", "GoogleRecaptchaSecret")

	// Database defaults
	viper.SetDefault("database.name", "DATABASE_NAME")
	viper.SetDefault("database.username", "DATABASE_USERNAME")
	viper.SetDefault("database.password", "DATABASE_PASSWORD")
	viper.SetDefault("database.host", "DATABASE_HOST_URL")
	viper.SetDefault("database.port", "5777")
	viper.SetDefault("database.logmode", false)

	// Email defaults
	viper.SetDefault("email.host", "smtp.gmail.com")
	viper.SetDefault("email.port", "25")
	viper.SetDefault("email.username", "EMAIL")
	viper.SetDefault("email.fromName", "$PASSWORD_MANAGER_NAME")
	viper.SetDefault("email.fromEmail", "EMAIL")
	viper.SetDefault("email.apiKey", "SENDGRID_API_KEY")
}

func generateKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)

	if err != nil {
		return "SENDGRID_API_KEY"
	}

	keyEnc := base64.StdEncoding.EncodeToString(key)
	return keyEnc
}
