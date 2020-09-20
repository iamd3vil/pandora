package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

// AddToBox adds a file to pandora's box
func (h *Hub) AddToBox(ctx *cli.Context) error {
	const errMsg = "error while add file to the box: %v"
	if !ctx.Args().Present() {
		return fmt.Errorf(errMsg, "file name can't be blank")
	}

	fp := ctx.Args().First()
	if filepath.IsAbs(fp) {
		return fmt.Errorf(errMsg, "filepath can't be absolute")
	}

	// HACK: Return errors for files which start with `.` or `/` or `\`
	if strings.HasPrefix(fp, ".") || strings.HasPrefix(fp, "/") || strings.HasPrefix(fp, `\`) {
		return fmt.Errorf(errMsg, "only relative paths from the root directory of the repo is accepted")
	}

	// Check if file exists
	_, err := os.Stat(fp)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}

	// Add file to config
	cfg, err := getConfig()
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}

	cfg.Files = append(cfg.Files, fp)

	err = writeConfig(cfg)
	if err != nil {
		return fmt.Errorf(errMsg, fmt.Sprintf("can't write config to config.json: %v", err))
	}

	return nil
}
