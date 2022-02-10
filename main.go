package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	expect "github.com/google/goexpect"
)

const (
	timeout = 30 * time.Second
	FIX_DSR = "\x1b[1;"
)

var (
	userRE    = regexp.MustCompile("Select language:")
	projectRE = regexp.MustCompile("Select project type:")
	starterRE = regexp.MustCompile("Which starter project do you want to use")
	nameRE    = regexp.MustCompile("Enter component name")
	endRE     = regexp.MustCompile("directly")
)

func helper() string {
	directory, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(directory)
	return directory
}

func main() {
	flag.Parse()
	dir := helper()
	defer os.RemoveAll(dir)

	e, _, err := expect.Spawn("odo init", -1)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	st, _, _ := e.Expect(userRE, timeout)
	fmt.Println(st, err)
	e.Send("go\n")

	st, _, _ = e.Expect(projectRE, timeout)
	fmt.Println(st, err)
	e.Send("\n")

	st, _, err = e.Expect(starterRE, timeout*3)
	fmt.Println(st, err)
	e.Send("\n")

	st, _, _ = e.Expect(nameRE, timeout)
	fmt.Println(st)
	e.Send(FIX_DSR + "mygoapp\n")

	st, _, _ = e.Expect(endRE, timeout)
	fmt.Println(st)

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		{
			fmt.Println(file)
		}
	}
	//
}
