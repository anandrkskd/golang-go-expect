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
)

var (
	userRE    = regexp.MustCompile("[0-9a-zA-Z]Select language:[0-9a-zA-Z]")
	passRE    = regexp.MustCompile("[0-9a-zA-Z]")
	projectRE = regexp.MustCompile("[0-9a-zA-Z]Select project type:[0-9a-zA-Z]")
	starterRE = regexp.MustCompile("[0-9a-zA-Z]Which starter project do you want to use[0-9a-zA-Z]")
	endRE     = regexp.MustCompile("[0-9a-zA-Z]directly[0-9a-zA-Z]")
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

	e, _, err := expect.Spawn(fmt.Sprintf("odo init"), -1)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	st, _, _ := e.Expect(userRE, timeout)
	fmt.Println(st, err)
	e.Send("go\n")

	st, _, _ = e.Expect(passRE, timeout)
	fmt.Println(st, err)
	e.Send("\n")

	st, _, _ = e.Expect(projectRE, timeout)
	fmt.Println(st, err)
	e.Send("\n")

	st, _, err = e.Expect(starterRE, timeout*3)
	fmt.Println(st, err)
	e.Send("\n")

	// st, _, _ = e.Expect(passRE, timeout)
	// // fmt.Println(st)
	// e.Send("mygoapp\n")

	st, _, _ = e.Expect(endRE, timeout)
	// fmt.Println(st, match)

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		{
			fmt.Println(file)
		}
	}
	//
}
