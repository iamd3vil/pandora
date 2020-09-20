package main

import (
	"io"
	"os"
	"path"
	"runtime"

	"github.com/iamd3vil/pandora/internal/admin"
)

func getHomeDir() string {
	var (
		home string
	)
	if runtime.GOOS == "linux" {
		home = os.Getenv("XDG_CONFIG_HOME")
		if home == "" {
			home = os.Getenv("HOME")
		}
	}

	return path.Clean(home)
}

func getAdmins() (admin.Admins, error) {
	f, err := os.Open(path.Join(DefaultBox, "admins.json"))
	if err != nil {
		return admin.Admins{}, err
	}
	defer f.Close()

	admins := admin.NewAdmins()
	err = admins.Read(f)
	if err != nil {
		return admin.Admins{}, err
	}
	return admins, nil
}

func getConfig() (Config, error) {
	f, err := os.Open(path.Join(DefaultBox, "config.json"))
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	cfg := Config{
		Files: make([]string, 0),
	}
	err = cfg.Read(f)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func writeConfig(cfg Config) error {
	f, err := os.Create(path.Join(DefaultBox, "config.json"))
	if err != nil {
		return err
	}
	defer f.Close()

	err = cfg.Save(f)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}
