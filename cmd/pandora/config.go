package main

import (
	"encoding/json"
	"io"
)

// Config is the config stored in the config.json inside .pandora folder
type Config struct {
	Files []string `json:"files"`
}

func (c *Config) Save(w io.Writer) error {
	return json.NewEncoder(w).Encode(c)
}

func (c *Config) Read(r io.Reader) error {
	return json.NewDecoder(r).Decode(c)
}
