package main

import (
	"flag"
	"fmt"
	"github.com/fteem/go-geoip/api"
	rp "github.com/fteem/go-geoip/response_parser"
)

var target = flag.String("target", "", "Target (IP or hostname)")

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		usage()
	} else {
		data := api.Fetch(*target)
		response := rp.Parse(data)
		fmt.Println(response)
	}
}

func usage() {
	fmt.Println("Search the geolocation of an IP or a hostname.\n Usage:")
	flag.PrintDefaults()
}
