package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	api "github.com/seanrmurphy/lambda-openapi3-test/api"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	// Create an instance of our handler which satisfies the generated interface
	a := api.NewApi()

	// We now register our petStore above as the handler for the interface
	e := echo.New()

	api.RegisterHandlers(e, a)

	//h := api.Handler(a)

	//s := &http.Server{
	//Handler: h,
	//Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	//}

	// And we serve HTTP until the world ends.
	//log.Fatal(s.ListenAndServe())
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
