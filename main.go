package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/Netflix/go-expect"
	"github.com/hinshun/vt10x"
	"github.com/kr/pty"
)

func main() {

	tmpdir, _ := ioutil.TempDir("", "")
	os.Chdir(tmpdir)
	fmt.Println(tmpdir)
	defer os.RemoveAll(tmpdir)

	ptm, pts, err := pty.Open()
	if err != nil {
		log.Fatal(err)
	}

	term := vt10x.New(vt10x.WithWriter(pts))

	c, err := expect.NewConsole(expect.WithStdin(ptm), expect.WithStdout(term), expect.WithCloser(pts, ptm))
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	cmd := exec.Command("odo", "init")
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	res, err := c.ExpectString("Select language")
	fmt.Fprintln(buf, res)
	c.SendLine("go")
	res, err = c.ExpectString("Select project type")
	fmt.Fprintln(buf, res)
	c.SendLine("Go Runtime")
	res, err = c.ExpectString("Which starter project do you want to use")
	fmt.Fprintln(buf, res)
	c.SendLine("go-starter")
	res, err = c.ExpectString("Enter component name")
	fmt.Fprintln(buf, res)
	c.SendLine("mytestapp")
	res, err = c.ExpectString("Your new component \"mytestapp\" is ready in the current directory.")
	fmt.Fprintln(buf, res)

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	// Close the slave end of the pty, and read the remaining bytes from the master end.
	c.Tty().Close()

	fmt.Println(buf)

}
