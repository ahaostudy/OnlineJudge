package compiler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"main/config"
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
	exe := &Executable{path: config.ConfJudge.Exe.Python, args: []string{"-OO", c.bin}}
	return exe, nil
}

func (c *Python3) Destroy(removeCode bool) error {
	if removeCode {
		return os.Remove(c.bin)
	}
	return nil
}

func (c *Python3) SaveCode(code []byte) (string, error) {
	codeName := fmt.Sprintf("%s.py", uuid.NewString())
	codePath := filepath.Join(config.ConfJudge.File.CodePath, codeName)
	return codeName, os.WriteFile(codePath, code, 0644)
}
