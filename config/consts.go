package config

import "os"

var (
	Root string
)

func initConsts() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	Root = root

	return nil
}
