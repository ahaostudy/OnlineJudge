package compiler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"

	"main/services/judge/config"
	"main/services/judge/pkg/errs"
)

type GPP struct {
	code     string
	bin      string
	compiled bool
}

func (c *GPP) Build(codePath string) (msg string, err error) {
	c.code = codePath
	c.bin = filepath.Join(config.Config.File.TempPath, uuid.New().String())

	cmd := exec.Command(config.Config.Exe.Gpp, "-std=c++11", codePath, "-o", c.bin)

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

func (c *GPP) Executable() (*Executable, error) {
	exe := &Executable{path: c.bin}
	if !c.compiled {
		return exe, errs.ErrCodeNotCompiled
	}
	return exe, nil
}

func (c *GPP) Destroy(removeCode bool) error {
	if removeCode {
		err := os.Remove(c.code)
		if err != nil {
			return err
		}
	}
	err := os.Remove(c.bin)
	return err
}

func (c *GPP) SaveCode(code []byte) (string, error) {
	codeName := fmt.Sprintf("%s.cpp", uuid.NewString())
	codePath := filepath.Join(config.Config.File.CodePath, codeName)
	return codeName, os.WriteFile(codePath, code, 0644)
}
