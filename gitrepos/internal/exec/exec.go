package exec

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Command struct {
	cmd    *exec.Cmd
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func NewCommand(argv []string, dir string) *Command {
	cmd := new(Command)

	cmd.stdout = bytes.Buffer{}
	cmd.stderr = bytes.Buffer{}

	cmd.cmd = exec.Command(argv[0], argv[1:]...)
	cmd.cmd.Dir = dir
	cmd.cmd.Stdout = &cmd.stdout
	cmd.cmd.Stderr = &cmd.stderr

	return cmd
}

func (cmd *Command) Start() error {
	return cmd.cmd.Start()
}

func (cmd *Command) Wait() error {
	return cmd.cmd.Wait()
}

func (cmd *Command) Stdout() string {
	return cmd.stdout.String()
}

func (cmd *Command) Stderr() string {
	return cmd.stderr.String()
}

func Run(argv []string, execPath string) error {
	cmd := NewCommand(argv, execPath)
	err := cmd.Start()
	msg := fmt.Sprintf("error: %v", err)
	if err != nil {
		return fmt.Errorf("start failed. the %s", msg)
	}
	err = cmd.Wait()
	msg = fmt.Sprintf("error: %v", err)
	if err != nil {
		return fmt.Errorf("wait failed, message=%s", msg)
	}
	return nil
}

// Exists is exist this command
func Exists(bin string) bool {
	_, err := Find(bin)
	if err != nil {
		return false
	}
	return true
}

// Find find the command
func Find(bin string) (string, error) {
	return exec.LookPath(bin)
}
