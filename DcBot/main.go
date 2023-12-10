package main

import (
	"log"
	"os"
)

var (
	outfile, _ = os.Create("RSSParser.log")
	l          = log.New(outfile, "", log.LstdFlags|log.Lshortfile)
)

func main() {

	ConnectToDC()
}
