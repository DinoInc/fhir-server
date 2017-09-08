package main

import (
	"flag"

	"github.com/intervention-engine/fhir/server"
)

func main() {
	reqLog := flag.Bool("reqlog", false, "Enables request logging -- do NOT use in production")
	flag.Parse()
	s := server.NewServer("localhost")
	if *reqLog {
		s.Engine.Use(server.RequestLoggerHandler)
	}

	s.Run(server.DefaultConfig)
}
