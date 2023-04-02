package fluentio

import (
	"encoding/json"
	"github.com/fluent/fluent-logger-golang/fluent"
	"io"
)

type Writer struct {
	client *fluent.Fluent
	tag    string
}

var _ io.Writer = (*Writer)(nil)

func New(opts ...func(config *Config)) (*Writer, error) {
	config := new(Config)
	for _, opt := range opts {
		opt(config)
	}

	if config == nil {
		return nil, ErrNoConfigProvided
	}

	var cfg *fluent.Config

	if config.basicConfig != nil {
		cfg = &fluent.Config{
			FluentHost: config.basicConfig.FluentHost,
			FluentPort: config.basicConfig.FluentPort,
		}
	}

	if config.fluentConfig != nil {
		cfg = config.fluentConfig
	}

	if cfg == nil {
		return nil, ErrNoConfigProvided
	}

	client, err := fluent.New(*cfg)
	if err != nil {
		return nil, err
	}

	var tag string
	if config.tag != "" {
		tag = config.tag
	}

	return &Writer{
		client: client,
		tag:    tag,
	}, nil
}

func (f *Writer) Write(p []byte) (n int, err error) {
	var m map[string]interface{}
	err = json.Unmarshal(p, &m)
	if err != nil {
		return 0, err
	}

	err = f.client.Post(f.tag, m)
	if err != nil {
		return 0, err
	}

	return len(p), nil
}
