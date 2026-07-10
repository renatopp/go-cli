package internal

import (
	"context"
	"fmt"
	"io"
	"maps"
	"os/exec"
	"time"
)

// Shell represents a shell command to be executed, along with its arguments.
// It is a convenience wrapper around exec.Cmd that allows setting various
// options like working directory, environment variables and timeout.
type Shell struct {
	command string
	args    []string
	dir     string
	envs    map[string]string
	timeout time.Duration
	stdin   io.Reader
	stdout  io.Writer
	stderr  io.Writer
}

// NewShell creates a new Shell instance with the specified command and
// arguments.
func NewShell(command string, args ...string) *Shell {
	return &Shell{
		command: command,
		args:    args,
		dir:     "",
		envs:    make(map[string]string),
		timeout: 0,
		stdin:   nil,
		stdout:  nil,
		stderr:  nil,
	}
}

// WithDir sets the working directory for the shell command.
func (s *Shell) WithDir(dir string) *Shell {
	s.dir = dir
	return s
}

// WithEnv sets an environment variable for the shell command.
func (s *Shell) WithEnv(key, value string) *Shell {
	s.envs[key] = value
	return s
}

// WithEnvs sets multiple environment variables for the shell command.
func (s *Shell) WithEnvs(envs map[string]string) *Shell {
	maps.Copy(s.envs, envs)
	return s
}

// WithTimeout sets a timeout for the shell command. If the command does not
// complete within the specified duration, it will be killed and an error will
// be returned.
func (s *Shell) WithTimeout(timeout time.Duration) *Shell {
	s.timeout = timeout
	return s
}

// WithStdin sets the standard input for the shell command.
func (s *Shell) WithStdin(stdin io.Reader) *Shell {
	s.stdin = stdin
	return s
}

// WithStdout sets the standard output for the shell command.
func (s *Shell) WithStdout(stdout io.Writer) *Shell {
	s.stdout = stdout
	return s
}

// WithStderr sets the standard error for the shell command.
func (s *Shell) WithStderr(stderr io.Writer) *Shell {
	s.stderr = stderr
	return s
}

// Cmd creates an exec.Cmd instance based on the Shell configuration.
func (s *Shell) Cmd() (*exec.Cmd, context.CancelFunc) {
	ctx := context.Background()
	var cancel context.CancelFunc = func() {}
	if s.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), s.timeout)
	}
	cmd := exec.CommandContext(ctx, s.command, s.args...)
	cmd.Dir = s.dir
	for k, v := range s.envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	if s.stdout != nil {
		cmd.Stdout = s.stdout
	}
	if s.stderr != nil {
		cmd.Stderr = s.stderr
	}
	return cmd, cancel
}

// Run executes the shell command and returns its combined output and any
// error that occurred during execution. If stdout or stderr is set, it
// will return an empty string for the output and the error from cmd.Run()
// instead.
func (s *Shell) Run() (string, error) {
	cmd, cancel := s.Cmd()
	defer cancel()
	if s.stdout == nil && s.stderr == nil {
		output, err := cmd.CombinedOutput()
		return string(output), err
	} else {
		err := cmd.Run()
		return "", err
	}
}
