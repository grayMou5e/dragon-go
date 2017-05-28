package config

import "go.uber.org/zap"

//NewLogger creates and returns new logger
func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"data.log",
	}
	cfg.DisableCaller = true

	return cfg.Build()
}
