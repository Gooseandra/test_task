package main

type (
	Config struct {
		Database DatabaseSettings `yaml:"database"`
	}

	DatabaseSettings struct {
		Arguments string `yaml:"arguments"`
		Type      string `yaml:"type"`
	}
)
