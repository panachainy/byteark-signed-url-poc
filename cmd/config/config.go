package config

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	ExampleConfig string `mapstructure:"EXAMPLE_CONFIG"`
}

var (
	// default values
	config = Configuration{
		ExampleConfig: "example config value",
	}
	configOnce sync.Once
)

func Provide() *Configuration {
	configOnce.Do(func() {
		// Default read config from `.env`
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("env read in config file failed: %v continue loading from env", err)
		}

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		BindEnvs(config)

		if err := viper.Unmarshal(&config); err != nil {
			log.Panicf("env unable to decode into struct, %v", err)
		}
		log.Print(config)
	})

	return &config
}

func BindEnvs(dest interface{}, parts ...string) {
	ifv := reflect.ValueOf(dest)
	ift := reflect.TypeOf(dest)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)

		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			envKey := strings.Join(append(parts, tv), ".")
			err := viper.BindEnv(envKey)
			if err != nil {
				fmt.Printf("bind env key %s failed: %v\n", envKey, err)
			}
		}
	}
}

func ResetProvide() {
	configOnce = sync.Once{}
}
