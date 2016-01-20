// main
package main

import (
	"flag"
	"fmt"
)

var appName, name, file, cfg *string
var ver *int

func init() {
	ver = flag.Int("ver", 0, "Resource version")
	cfg = flag.String("cfg", "", "node login config")
	name = flag.String("name", "", "Resource name")
	file = flag.String("file", "", "Resource file")
	appName = flag.String("appname", "", "app name")
	flag.Parse()
}

func main() {
	var err error
	cmd := flag.Arg(0)
	switch cmd {
	case "push": //资源发布
		err = Push()
	case "release":
		err = Release()
	case "show":
		err = Show()
	case "clear":
		err = Clear()
	default:
		err = errCmd.Format(cmd)
	}
	if err != nil {
		fmt.Printf("An error occurred during program execution.\r\n%s\r\n", err.Error())
	}
}
