package compiler

import (
	"bytes"
	"github.com/google/uuid"
	"main/config"
	"main/services/judge/pkg/errs"
	"os"
	"os/exec"
	"path/filepath"
)

type GO struct {
	code     string
	bin      string
	compiled bool
}

func (c *GO) Build(codePath string) (msg string, err error) {
	c.code = codePath
	c.bin = filepath.Join(config.ConfJudge.File.TempPath, uuid.New().String())

	cmd := exec.Command(config.ConfJudge.Exe.Go, "build", "-o", c.bin, codePath)

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

func (c *GO) Executable() (*Executable, error) {
	exe := &Executable{path: c.bin}
	if !c.compiled {
		return exe, errs.ErrCodeNotCompiled
	}
	return exe, nil
}

func (c *GO) Destroy(removeCode bool) error {
	if removeCode {
		err := os.Remove(c.code)
		if err != nil {
			return err
		}
	}
	err := os.Remove(c.bin)
	return err
}
