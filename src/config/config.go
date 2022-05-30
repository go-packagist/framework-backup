package config

import (
	"github.com/spf13/viper"
)

// Config represents the config.
type Config struct {
	Options *Options
}

// New returns a new config instance.
func New(options *Options) *Config {
	options.Prepare()

	config := &Config{
		Options: options,
	}

	config.init()

	return config
}

// init initializes the config.
func (c *Config) init() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix(c.Options.GetPrefix())
	viper.AutomaticEnv()

	// read config
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error config file: %w \n", err))
	// }
}

// GetOptions returns the options.
func (c *Config) GetOptions() *Options {
	return c.Options
}

func (c *Config) Env(key string) interface{} {
	return viper.Get(key)
}

// GetEnvPath returns the env path.
func (c *Config) GetEnvPath() string {
	return c.Options.GetEnvPath()
}

func (c *Config) Add(key string, value interface{}) {
	viper.Set(key, value)
}

func (c *Config) Get(key string) interface{} {
	return viper.Get(key)
}

func (c *Config) GetString(key string) string {
	return viper.GetString(key)
}

func (c *Config) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (c *Config) GetAll() map[string]interface{} {
	return viper.AllSettings()
}

// Options represents the config options.
type Options struct {
	EnvPath string
	Prefix  string
}

// Prepare prepares the config options.
func (o *Options) Prepare() {
	if o.Prefix == "" {
		o.Prefix = "gp"
	}
}

// GetEnvPath returns the env path.
func (o *Options) GetEnvPath() string {
	return o.EnvPath
}

// GetPrefix returns the prefix.
func (o *Options) GetPrefix() string {
	return o.Prefix
}
