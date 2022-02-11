package main

import (
	"regexp"
	"time"
)

const (
	timeout = 10 * time.Minute
)

var (
	userRE   = regexp.MustCompile("[0-9a-zA-Z]Select language:[0-9a-zA-Z]")
	passRE   = regexp.MustCompile("[0-9a-zA-Z]")
	promptRE = regexp.MustCompile("%")
)

//func main() {
//	flag.Parse()
//	fmt.Println(term.Bluef("Telnet 1 example"))
//
//	e, _, err := expect.Spawn(fmt.Sprintf("main init"), -1)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer e.Close()
//
//	e.Expect(userRE, timeout)
//	e.Send()
//	e.Expect(passRE, timeout)
//	e.Send()
//	e.Expect(promptRE, timeout)
//	e.Send()
//	result, _, _ := e.Expect(promptRE, timeout)
//	e.Send("exit\n")
//
//	fmt.Println(term.Greenf("%s: result: %s\n", result))
//}
