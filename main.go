package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2/core"
	"github.com/AlecAivazis/survey/v2/terminal"
	goexpect "github.com/Netflix/go-expect"
	"github.com/hinshun/vt10x"
)

func init() {
	// disable color output for all prompts to simplify testing
	core.DisableColor = true
}

func main() {
	// buf := new(bytes.Buffer)
	//c, _, err := vt10x.NewVT10XConsole(goexpect.WithStdout(buf))

	c, _, err := vt10x.NewVT10XConsole(goexpect.WithStdout(os.Stdout))
	// c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	tmpdir, _ := ioutil.TempDir("", "")
	os.Chdir(tmpdir)
	fmt.Println(tmpdir)
	cmd := exec.Command("main", "init")

	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	go func() {
		c.ExpectEOF()
	}()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 2)
	time.Sleep(time.Second * 2)

	//lang := regexp.MustCompile("[0-9a-zA-Z]Select language:[0-9a-zA-Z]")

	// c.Expect(goexpect.RegexpPattern("Select", "language:"))
	_, err = c.Expect(goexpect.String("Select language:"))
	fmt.Print(err)
	c.Send(string("go"))
	time.Sleep(time.Second * 2)
	c.Send(string(terminal.KeyEnter))
	time.Sleep(time.Second * 3)
	c.Send(string(terminal.KeyEnter))
	time.Sleep(time.Second * 3)
	c.Send(string(terminal.KeyEnter))
	time.Sleep(time.Second * 3)
	c.SendLine("mygolang" + string(terminal.KeyEnter))
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	_ = c.Tty().Close()

	//fmt.Pri4t(buf.String())
	os.RemoveAll(tmpdir)
}
