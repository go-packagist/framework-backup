package config

import (
	"github.com/spf13/viper"
)

// Config represents the config.
type Config struct {
	viper   *viper.Viper
	Options *Options
}

// New returns a new config instance.
func New(options *Options) *Config {
	options.Prepare()

	config := &Config{
		viper:   viper.New(),
		Options: options,
	}

	config.init()

	return config
}

// init initializes the config.
func (c *Config) init() {
	c.viper.SetConfigType("env")
	c.viper.AddConfigPath(".")
	c.viper.SetEnvPrefix(c.Options.GetPrefix())
	c.viper.AutomaticEnv()
}

// GetOptions returns the options.
func (c *Config) GetOptions() *Options {
	return c.Options
}

func (c *Config) Env(key string) interface{} {
	return c.viper.Get(key)
}

// GetEnvPath returns the env path.
func (c *Config) GetEnvPath() string {
	return c.Options.GetEnvPath()
}

// Add adds the key and value to the config.
func (c *Config) Add(key string, value interface{}) {
	c.viper.Set(key, value)
}

// Set sets the options.
func (c *Config) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

// Get returns the value by the key.
func (c *Config) Get(key string) interface{} {
	return c.viper.Get(key)
}

// GetString returns the string by the key.
func (c *Config) GetString(key string) string {
	return c.viper.GetString(key)
}

// GetBool returns the bool by the key.
func (c *Config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

// GetAll returns the all config.
func (c *Config) GetAll() map[string]interface{} {
	return c.viper.AllSettings()
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

func (o *Options) SetPrefix(prefix string) *Options {
	o.Prefix = prefix

	return o
}

// GetPrefix returns the prefix.
func (o *Options) GetPrefix() string {
	return o.Prefix
}
