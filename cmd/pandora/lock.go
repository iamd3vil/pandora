package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"filippo.io/age"
	"filippo.io/age/agessh"
	"github.com/iamd3vil/pandora/internal/admin"
	"github.com/urfave/cli/v2"
)

func (h *Hub) LockBox(ctx *cli.Context) error {
	cf, err := os.Open(path.Join(DefaultBox, "config.json"))
	if err != nil {
		return fmt.Errorf("error locking pandora's box: %v", err)
	}
	defer cf.Close()
	cfg := &Config{}
	err = cfg.Read(cf)
	if err != nil {
		return fmt.Errorf("error locking pandora's box: %v", err)
	}

	// Read admins
	af, err := os.Open(path.Join(DefaultBox, "admins.json"))
	if err != nil {
		return fmt.Errorf("error locking pandora's box: %v", err)
	}
	defer af.Close()
	admins := admin.NewAdmins()
	err = admins.Read(af)
	if err != nil {
		return fmt.Errorf("error locking pandora's box: %v", err)
	}

	// Loop over each file and encrypt and move it to the crct path
	pathInBox := path.Join(DefaultBox, "files")
	os.Mkdir(pathInBox, 0750)
	for _, fp := range cfg.Files {
		p := path.Join(pathInBox, fp)
		recipients := []age.Recipient{}
		for _, a := range admins.Adms {
			r, err := agessh.ParseRecipient(a.Key)
			if err != nil {
				return errors.New("error locking pandora's box: invalid key in admins")
			}
			recipients = append(recipients, r)
		}
		// Create the directory
		dir := path.Dir(p)
		os.MkdirAll(dir, 0750)

		f, err := os.Create(p)
		if err != nil {
			return fmt.Errorf("error locking pandora's box: %v", err)
		}
		s, err := age.Encrypt(f, recipients...)
		if err != nil {
			return fmt.Errorf("error locking pandora's box: error encrypting file: %v", err)
		}
		s.Close()
	}
	return nil
}
