package main

import (
	"fmt"
	"os"

	"github.com/ci2rpi/rpi-websrv/config"
	"github.com/ci2rpi/rpi-websrv/web_server"
)

var (
	server = new(web_server.WebServer)
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %v config.json\n", os.Args[0])
		os.Exit(1)
	}

	conf := config.NewConfigFromFile(os.Args[1])
	server.Run(conf.Port, conf.StaticPagesDirectory)
}
