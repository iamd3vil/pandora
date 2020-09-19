package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/iamd3vil/pandora/internal/admin"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const (
	DefaultBox = ".pandora"
)

// Hub contains all app context
type Hub struct {
	logger *zap.SugaredLogger
}

func (h *Hub) initBox(ctx *cli.Context) error {
	var (
		keyPath = path.Join(getHomeDir(), ".ssh", "id_rsa.pub")
		args    = ctx.Args()
	)
	if args.Present() {
		keyPath = args.First()
	}
	stat, err := os.Stat(DefaultBox)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(DefaultBox, 0755)
			if err != nil {
				return err
			}
			goto Main
		} else {
			return err
		}
	}
	if stat.IsDir() {
		return errors.New("error while initializing pandora: already initialized")
	}

Main:
	// Read SSH Key and add to admin
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("error while initializing pandora: %v", err)
	}

	admins := admin.NewAdmins()

	admin := admin.Admin{
		Key:  strings.TrimSuffix(string(key), "\n"),
		Type: SSHKey,
	}

	admins.Add(admin)

	f, err := os.Create(path.Join(DefaultBox, "admins.json"))
	if err != nil {
		return fmt.Errorf("error while initializing pandora: %v", err)
	}
	defer f.Close()

	admins.Save(f)
	h.logger.Info("pandora is initialized")

	return nil
}

// NewHub returns a new instance of Hub
func NewHub(logger *zap.SugaredLogger) (*Hub, error) {
	return &Hub{
		logger: logger,
	}, nil
}
