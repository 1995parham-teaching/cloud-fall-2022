package config

import (
	"log"
	"strings"

	"github.com/1995parham-teaching/cloud-fall-2022/internal/metric"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/tracing"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type Config struct {
	Metric  metric.Config  `koanf:"metric"`
	Tracing tracing.Config `koanf:"tracing"`
}

// CLOUD_METRIC__ADDRESS
// metric__address
// metric.address

const (
	prefix    = "CLOUD_"
	seprator  = "__"
	delimeter = "."
)

func New() Config {
	k := koanf.New(".")

	// load default configuration from default function
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Printf("error loading config.toml: %s", err)
	}

	callback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, prefix))

		return strings.ReplaceAll(base, seprator, delimeter)
	}

	// load environment variables
	if err := k.Load(env.Provider(prefix, delimeter, callback), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	var instance Config
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	return instance
}
