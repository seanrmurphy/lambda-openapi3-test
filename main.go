package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	api "github.com/seanrmurphy/lambda-openapi3-test/api"
)

var e *echo.Echo
var echoAdapter *echoadapter.EchoLambda
var nullHandler = false

// Lambda mode determines whether this is run locally or remotely
var lambdaMode = false

func init() {

	// Create an instance of our handler which satisfies the generated interface
	a := api.NewApi()

	e = echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api.RegisterHandlers(e, a)

	if lambdaMode {
		echoAdapter = echoadapter.New(e)
	}
}

// Handler handles API requests
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if nullHandler {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       `{"message": "Null handler operational - always returns this message"}`,
		}, nil
	} else {
		fmt.Printf("Path = %v, Resource = %v, req = %v\n", req.Path, req.Resource, req)
		req.Path = req.Path[(len(req.RequestContext.Stage) + 2):]
		return echoAdapter.Proxy(req)
	}
}

func main() {
	if lambdaMode {
		lambda.Start(Handler)
	} else {
		var port = flag.Int("port", 8080, "Port for test HTTP server")
		flag.Parse()
		e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
	}
}
