package fluentio

import "github.com/fluent/fluent-logger-golang/fluent"

type Config struct {
	fluentConfig *fluent.Config
	basicConfig  *BasicConfig
	tag          string
}

type BasicConfig struct {
	FluentHost string
	FluentPort int
}

func WithBasicConfig(host string, port int) func(*Config) {
	return func(c *Config) {
		c.basicConfig = &BasicConfig{
			FluentHost: host,
			FluentPort: port,
		}
	}
}

func WithFluentConfig(config *fluent.Config) func(*Config) {
	return func(c *Config) {
		c.fluentConfig = config
	}
}

func WithTag(tag string) func(*Config) {
	return func(c *Config) {
		c.tag = tag
	}
}
