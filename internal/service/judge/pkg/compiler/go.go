package compiler

import (
	"bytes"
	"fmt"
	"main/config"
	"main/internal/service/judge/pkg/errs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
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

func (c *GO) SaveCode(code []byte) (string, error) {
	codeName := fmt.Sprintf("%s.go", uuid.NewString())
	codePath := filepath.Join(config.ConfJudge.File.CodePath, codeName)
	return codeName, os.WriteFile(codePath, code, 0644)
}
