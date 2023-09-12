package compiler

import (
	"bytes"
	"main/config"
	"main/internal/service/judge/pkg/errs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
)

type GCC struct {
	code     string
	bin      string
	compiled bool
}

func (c *GCC) Build(codePath string) (msg string, err error) {
	c.code = codePath
	c.bin = filepath.Join(config.ConfJudge.File.TempPath, uuid.New().String())

	cmd := exec.Command(config.ConfJudge.Exe.Gcc, codePath, "-Wall", "-std=c11", "-o", c.bin)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	} else if len(stderr.String()) > 0 {
		return stderr.String(), errs.ErrCompilationFailed
	}

	c.compiled = true

	return "", nil
}

func (c *GCC) Executable() (*Executable, error) {
	exe := &Executable{path: c.bin}
	if !c.compiled {
		return exe, errs.ErrCodeNotCompiled
	}
	return exe, nil
}

func (c *GCC) Destroy(removeCode bool) error {
	if removeCode {
		err := os.Remove(c.code)
		if err != nil {
			return err
		}
	}
	err := os.Remove(c.bin)
	return err
}
