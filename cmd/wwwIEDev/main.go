package main

import (
	"flag"
	"wwwIEDev"
)

func main() {
	var renderDIR string
	flag.StringVar(&renderDIR, "path", "public", "render directory path")
	flag.Parse()
	wwwIEDev.Run(renderDIR)
}
