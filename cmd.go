package main

import (
	"checkers/client"
	"checkers/server"
	"flag"
)

func ParseFlags() {
	flag.IntVar(&server.Port, "server", 8080, "--server")
	flag.StringVar(&client.ServeAddr, "connect", "", "--connect <SERVER:PORT>")
	flag.Parse()
}
