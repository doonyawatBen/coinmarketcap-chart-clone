package main

import (
	"errors"
	"log"

	"github.com/lodashventure/nlp/handlers"

	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/modules"
)

func init() {
	infrastructure.Init()
}

func main() {
	switch infrastructure.ConfigGlobal.ServiceName {
	case "webservice":
		rest, route := modules.Webservice()
		handlers.RegisterRoute(route)
		modules.WebserviceEnd(rest)
	default:
		log.Fatalln(errors.New(infrastructure.ConfigGlobal.ServiceName + "is not supported."))
	}
}
