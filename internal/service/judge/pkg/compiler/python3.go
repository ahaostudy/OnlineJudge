package compiler

import (
	"main/config"
	"os"
)

type Python3 struct {
	bin      string
	compiled bool
}

func (c *Python3) Build(codePath string) (msg string, err error) {
	c.bin = codePath
	c.compiled = true

	return "", nil
}

func (c *Python3) Executable() (*Executable, error) {
	exe := &Executable{path: config.ConfJudge.Exe.Python, args: []string{c.bin}}
	return exe, nil
}

func (c *Python3) Destroy(removeCode bool) error {
	if removeCode {
		return os.Remove(c.bin)
	}
	return nil
}
