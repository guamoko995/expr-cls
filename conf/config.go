package conf

import (
	buildEnv "github.com/guamoko995/expr-cls/builder/env"
)

var (
	// DefaultMaxNodes represents default maximum allowed AST nodes by the compiler.
	DefaultMaxNodes uint = 1e4
)

type Config struct {
	BuildStageEnvironment *buildEnv.Env
	MaxNodes              uint
}

// New creates new config with default values.
func New() *Config {
	c := &Config{
		BuildStageEnvironment: buildEnv.Global,
		MaxNodes:              DefaultMaxNodes,
	}
	return c
}

func (c *Config) WithEnv(buildStageEnv *buildEnv.Env) {
	c.BuildStageEnvironment = buildStageEnv
}
