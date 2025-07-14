package config

import "github.com/spf13/viper"

var config Config

// Config interface associated with reading/saving configuration files.
type Config interface {
	Load() error
	ReadSection(key string, val interface{}) error
}

// Init config from file
func Init(file string) error {
	config = newConfig(file)
	return config.Load()
}

var _ Config = (*viperImpl)(nil)

type viperImpl struct {
	file  string
	viper *viper.Viper
}

func newConfig(file string) Config {
	return &viperImpl{file: file, viper: viper.New()}
}

// Load config from file
func (v *viperImpl) Load() error {
	v.viper.SetConfigFile(v.file)
	return v.viper.ReadInConfig()
}

// ReadSection unmarshals the config into a Struct
func (v *viperImpl) ReadSection(key string, val interface{}) error {
	return v.viper.UnmarshalKey(key, val)
}
