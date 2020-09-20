package main

import (
	"go.uber.org/zap"
)

const (
	DefaultBox = ".pandora"
)

// Hub contains all app context
type Hub struct {
	logger *zap.SugaredLogger
}

// NewHub returns a new instance of Hub
func NewHub(logger *zap.SugaredLogger) (*Hub, error) {
	return &Hub{
		logger: logger,
	}, nil
}
