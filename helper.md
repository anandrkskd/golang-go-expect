package main

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	goexpect "github.com/Netflix/go-expect"
	"github.com/hinshun/vt10x"
	"gotest.tools/v3/assert"
)

type promptTest struct {
	procedure func(*goexpect.Console) error
}

func RunPromptTest(t *testing.T, test promptTest) {
	test.runTest(t, test.procedure, func(stdio terminal.Stdio) error {
		//	res.AskOpts = WithStdio(stdio)
		//	err := res.createInteractive()
		cmd := exec.Command("main", "init")
		cmd.Stdin = c.Tty()
		cmd.Stdout = c.Tty()
		cmd.Stderr = c.Tty()
		err := cmd.Start()
		if err != nil {
			if err.Error() == "resource already exist" {
				return nil
			}
			return err
		}

		return err
	})
}

func stdio(c *goexpect.Console) terminal.Stdio {
	return terminal.Stdio{In: c.Tty(), Out: c.Tty(), Err: c.Tty()}
}

func (pt *promptTest) runTest(t *testing.T, procedure func(*goexpect.Console) error, test func(terminal.Stdio) error) {
	t.Parallel()

	// Multiplex output to a buffer as well for the raw bytes.
	buf := new(bytes.Buffer)
	c, state, err := vt10x.NewVT10XConsole(goexpect.WithStdout(buf))
	assert.NilError(t, err)
	defer c.Close()

	donec := make(chan struct{})
	go func() {
		defer close(donec)
		if err := procedure(c); err != nil {
			t.Logf("procedure failed: %v", err)
		}
	}()

	assert.NilError(t, test(stdio(c)))

	// Close the slave end of the pty, and read the remaining bytes from the master end.
	_ = c.Tty().Close()
	<-donec

	t.Logf("Raw output: %q", buf.String())

	// Dump the terminal's screen.
	t.Logf("\n%s", goexpect.StripTrailingEmptyLines(state.String()))
}

// WithStdio helps to test interactive command
// by setting stdio for the ask function
func WithStdio(stdio terminal.Stdio) survey.AskOpt {
	return func(options *survey.AskOptions) error {
		options.Stdio.In = stdio.In
		options.Stdio.Out = stdio.Out
		options.Stdio.Err = stdio.Err
		return nil
	}
}
