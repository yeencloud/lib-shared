package config

import (
	"reflect"

	"github.com/fatih/structs"
	log "github.com/sirupsen/logrus"
)

var localCfg *Config

type Config struct {
	sourceInterface SourceConfig

	typeMap map[reflect.Kind]valueHandler
}

func NewConfig(sourceInterface SourceConfig) *Config {
	var secret Secret

	config := &Config{
		sourceInterface: sourceInterface,
		typeMap: map[reflect.Kind]valueHandler{
			reflect.TypeOf(secret).Kind(): handleSecret,
			reflect.String:                handleString,
			reflect.Int:                   handleInt,
			reflect.Bool:                  handleBool,
		},
	}

	localCfg = config

	return config
}

func (cfg *Config) AvailableTypes() []string {
	types := make([]string, 0, len(cfg.typeMap))
	for k := range cfg.typeMap {
		types = append(types, k.String())
	}
	return types
}

func FetchConfig[T any]() (*T, error) {
	if localCfg == nil {
		return nil, &ConfigurationNotInitializedError{}
	}

	var object T
	s := structs.New(&object)

	for _, field := range s.Fields() {
		configKey := field.Tag("config")
		if configKey == "" {
			log.WithField("field", field.Name()).Warn("`config` tag is missing from struct field")
			continue
		}

		value, err := localCfg.sourceInterface.ReadString(configKey)
		if err != nil {
			log.WithField("key", configKey).WithError(err).Warn("Failed to read config")
			return nil, err
		}

		usingDefault := false
		defaultValue := field.Tag("default")
		if value == "" {
			value = defaultValue
			if value != "" {
				usingDefault = true
			}
		}

		handler, found := localCfg.typeMap[field.Kind()]
		if !found {
			log.WithField("config", configKey).Warnf("Unsupported config type: %s for variable %s", field.Kind().String(), field.Name())
			return nil, UnsupportedConfigTypeError{Type: field.Kind().String(), Variable: field.Name()}
		}

		err = handler(field, value)
		log.WithField(configKey, field.Value()).WithField("default", usingDefault).Debugf("Config loaded: %s", configKey)
		if err != nil {
			return nil, err
		}
	}
	return &object, nil
}
