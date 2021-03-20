package main

import (
	"flag"
	"log"
)

var addr = flag.String("addr", ":8080", "server port")

func main() {
	log.Fatal(listenAndServe(*addr))
}
