package compiler

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"main/config"
	"main/services/judge/pkg/errs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Java struct {
	code     string
	binPath  string
	binName  string
	compiled bool
}

func (c *Java) Build(codePath string) (msg string, err error) {
	c.code = codePath
	c.binPath = filepath.Join(config.ConfJudge.File.TempPath, uuid.New().String())
	c.binName = strings.TrimSuffix(filepath.Base(codePath), filepath.Ext(codePath))
	if err := os.MkdirAll(c.binPath, os.ModePerm); err != nil {
		return "", err
	}

	cmd := exec.Command(config.ConfJudge.Exe.Javac, "-d", c.binPath, "-encoding", "UTF8", codePath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		return stderr.String(), err
	} else if len(stderr.String()) > 0 {
		return stderr.String(), errs.ErrCompilationFailed
	}

	c.compiled = true
	return "", nil
}

func (c *Java) Executable() (*Executable, error) {
	exe := &Executable{
		path:   config.ConfJudge.Exe.Java,
		args:   []string{"-cp", c.binPath, "-XX:MaxRAM={MaxMemory}", "-Djava.security.manager", "-Dfile.encoding=UTF-8", "-Djava.awt.headless=true", c.binName}, // args中支持使用{参数}，在调用时会自动获取
		kwargs: map[string]interface{}{"max_memory": nil, "memory_limit_check_only": 1},
	}
	if !c.compiled {
		return exe, errs.ErrCodeNotCompiled
	}
	return exe, nil
}

func (c *Java) Destroy(removeCode bool) error {
	if removeCode {
		err := os.Remove(c.code)
		if err != nil {
			return err
		}
	}
	err := os.RemoveAll(c.binPath)
	return err
}
