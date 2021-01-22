//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ../../petstore-expanded.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ../../petstore-expanded.yaml

package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type Api struct {
	NextId int64
	Lock   sync.Mutex
}

func NewApi() *Api {
	return &Api{
		NextId: 1000,
	}
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendApiError(ctx echo.Context, code int, message string) error {
	apiError := ErrorResponse{
		ErrorNumber: code,
		ErrorString: message,
	}
	err := ctx.JSON(code, apiError)
	return err
}

func (*Api) GetApiIdentifier(ctx echo.Context) error {
	str := "oapi-codegen Lambda integration API (OpenAPI v3 version) - version 1.0"

	r := SimpleMessageResponse{
		Message: str,
	}

	return ctx.JSON(http.StatusOK, r)
}

// An endpoint which returns a HTTP response code only
// (GET /no-params/empty-response)
func (*Api) GetNoParamsEmptyResponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}

// An endpoint which returns a HTTP response code and a JSON encoded message in a body
// (GET /no-params/simple-response)
func (*Api) GetNoParamsSimpleResponse(ctx echo.Context) error {
	str := "Response from GetNoParamsSimpleResponse function"

	r := SimpleMessageResponse{
		Message: str,
	}
	return ctx.JSON(http.StatusOK, r)
}

// An endpoint which returns a HTTP response code and a complex JSON object in a body
// (GET /no-params/complex-response)
func (*Api) GetNoParamsComplexResponse(ctx echo.Context) error {
	stringArray := []string{"param1", "param2"}
	descriptor1 := "A descriptive string (object1)"
	intVal := rand.Int()
	o1 := SimpleObjectOne{
		StringArray: stringArray,
		Descriptor:  descriptor1,
		IntVal:      intVal,
	}

	descriptor2 := "A descriptive string (object2)"
	intArray := []int{rand.Int(), rand.Int(), rand.Int(), rand.Int()}
	o2 := SimpleObjectTwo{
		Descriptor: descriptor2,
		IntArray:   intArray,
	}
	optionalParam := rand.Int()
	r := ComplexObjectResponse{
		Object1:       o1,
		Object2:       o2,
		OptionalParam: &optionalParam,
	}

	return ctx.JSON(http.StatusOK, r)
}

// An endpoint which returns a HTTP response code and an error response
// (GET /no-params/error-response)
func (*Api) GetNoParamsErrorResponse(ctx echo.Context) error {
	str := "Calling this endpoint (GetNoParamsErrorResponse) deliberately generates an error case."
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusTeapot, r)

}

// Endpoint which takes a parameter and returns a HTTP response code only
// (GET /path-param/empty-response/{path-param})
func (*Api) GetPathParamEmptyResponse(ctx echo.Context, pathParam int) error {
	return ctx.JSON(http.StatusOK, nil)
}

// Endpoint which takes a parameter and returns a HTTP response containing a simple message
// (GET /path-param/simple-response/{path-param})
func (*Api) GetPathParamSimpleResponse(ctx echo.Context, pathParam int) error {
	str := fmt.Sprintf("Response from GetPathParamSimpleResponse function - input param = %v", pathParam)

	r := SimpleMessageResponse{
		Message: str,
	}

	return ctx.JSON(http.StatusOK, r)
}

// An endpoint which takes an input parameters and returns a HTTP response code and a complex JSON object in a body
// (GET /path-param/complex-response/{path-param})
func (*Api) GetPathParamComplexResponse(ctx echo.Context, pathParam int) error {
	str := fmt.Sprintf("Input param = %v", pathParam)

	stringArray := []string{"param1", "param2", str}
	descriptor1 := "A descriptive string (object1)"
	intVal := rand.Int()
	o1 := SimpleObjectOne{
		StringArray: stringArray,
		Descriptor:  descriptor1,
		IntVal:      intVal,
	}

	descriptor2 := "A descriptive string (object2)"
	intArray := []int{rand.Int(), rand.Int(), rand.Int(), rand.Int()}
	o2 := SimpleObjectTwo{
		Descriptor: descriptor2,
		IntArray:   intArray,
	}
	optionalParam := rand.Int()
	r := ComplexObjectResponse{
		Object1:       o1,
		Object2:       o2,
		OptionalParam: &optionalParam,
	}
	return ctx.JSON(http.StatusOK, r)
}

// An endpoint which takes an input parameter and returns a HTTP response code and an error response
// (GET /path-param/error-response/{path-param})
func (*Api) GetPathParamErrorResponse(ctx echo.Context, pathParam int) error {
	str := fmt.Sprintf("Calling this endpoint (GetPathParamErrorResponse) deliberately generates an error case - Input param = %v", pathParam)
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusTeapot, r)
}

// Endpoint which takes a parameter and returns a HTTP response code only
// (GET /query-param/empty-response)
func (*Api) GetQueryParamEmptyResponse(ctx echo.Context, params GetQueryParamEmptyResponseParams) error {
	return ctx.JSON(http.StatusOK, nil)
}

// Endpoint which takes a parameter and returns a HTTP response containing a simple message
// (GET /query-param/simple-response)
func (*Api) GetQueryParamSimpleResponse(ctx echo.Context, params GetQueryParamSimpleResponseParams) error {
	str := fmt.Sprintf("Response from GetPathParamSimpleResponse function - input param = %v", params.QueryParam)

	r := SimpleMessageResponse{
		Message: str,
	}

	return ctx.JSON(http.StatusOK, r)
}

// An endpoint which takes an input parameters and returns a HTTP response code and a complex JSON object in a body
// (GET /query-param/complex-response)
func (*Api) GetQueryParamComplexResponse(ctx echo.Context, params GetQueryParamComplexResponseParams) error {
	str := fmt.Sprintf("Input param = %v", params.QueryParam)

	stringArray := []string{"param1", "param2", str}
	descriptor1 := "A descriptive string (object1)"
	intVal := rand.Int()
	o1 := SimpleObjectOne{
		StringArray: stringArray,
		Descriptor:  descriptor1,
		IntVal:      intVal,
	}

	descriptor2 := "A descriptive string (object2)"
	intArray := []int{rand.Int(), rand.Int(), rand.Int(), rand.Int()}
	o2 := SimpleObjectTwo{
		Descriptor: descriptor2,
		IntArray:   intArray,
	}
	optionalParam := rand.Int()
	r := ComplexObjectResponse{
		Object1:       o1,
		Object2:       o2,
		OptionalParam: &optionalParam,
	}
	return ctx.JSON(http.StatusOK, r)
}

// An endpoint which takes an input parameter and returns a HTTP response code and an error response
// (GET /query-param/error-response)
func (*Api) GetQueryParamErrorResponse(ctx echo.Context, params GetQueryParamErrorResponseParams) error {
	str := fmt.Sprintf("Calling this endpoint (GetPathParamErrorResponse) deliberately generates an error case - Input param = %v", params.QueryParam)
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusTeapot, r)
}

// Get endpoint defined which simply gives response indicating that POST should be used
// (GET /body-param/empty-response)
func (*Api) GetBodyParamEmptyResponse(ctx echo.Context) error {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusNotImplemented, r)

}

// Endpoint which takes a body parameter and returns a HTTP response code only
// (POST /body-param/empty-response)
func (*Api) PostBodyParamEmptyResponse(ctx echo.Context) error {
	// context Binding is necessary to perform data validation
	var inputParams InputObject
	err := ctx.Bind(&inputParams)
	if err != nil {
		errMessage := ErrorResponse{
			ErrorNumber: 400,
			ErrorString: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, errMessage)
	}
	return ctx.JSON(http.StatusOK, nil)
}

// Get endpoint defined which simply gives response indicating that POST should be used
// (GET /body-param/simple-response)
func (*Api) GetBodyParamSimpleResponse(ctx echo.Context) error {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusNotImplemented, r)
}

// Endpoint which takes a body parameter and returns a HTTP response containing a simple message
// (POST /body-param/simple-response)
func (*Api) PostBodyParamSimpleResponse(ctx echo.Context) error {
	var inputParams InputObject
	err := ctx.Bind(&inputParams)
	if err != nil {
		errMessage := ErrorResponse{
			ErrorNumber: 400,
			ErrorString: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, errMessage)
	}
	str := fmt.Sprintf("Response from PostBodyParamSimpleResponse function - called with JSON object {'descriptor': '%v', 'int_val': %v, 'string': '%v'}",
		inputParams.Descriptor, inputParams.IntVal, inputParams.String)

	r := SimpleMessageResponse{
		Message: str,
	}
	return ctx.JSON(http.StatusOK, r)
}

// Get endpoint defined which simply gives response indicating that POST should be used
// (GET /body-param/complex-response)
func (*Api) GetBodyParamComplexResponse(ctx echo.Context) error {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusNotImplemented, r)
}

// An endpoint which takes an input parameters and returns a HTTP response code and a complex JSON object in a body
// (POST /body-param/complex-response)
func (a *Api) PostBodyParamComplexResponse(ctx echo.Context) error {
	var inputParams InputObject
	err := ctx.Bind(&inputParams)
	if err != nil {
		errMessage := ErrorResponse{
			ErrorNumber: 400,
			ErrorString: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, errMessage)
	}

	str := fmt.Sprintf("Response from PostBodyParamSimpleResponse function - called with JSON object {'descriptor': '%v', 'int_val': %v, 'string': '%v'}",
		inputParams.Descriptor, inputParams.IntVal, inputParams.String)

	stringArray := []string{"param1", "param2", str}
	descriptor1 := "A descriptive string (object1)"
	intVal := rand.Int()
	o1 := SimpleObjectOne{
		StringArray: stringArray,
		Descriptor:  descriptor1,
		IntVal:      intVal,
	}

	descriptor2 := "A descriptive string (object2)"
	intArray := []int{rand.Int(), rand.Int(), rand.Int(), rand.Int()}
	o2 := SimpleObjectTwo{
		Descriptor: descriptor2,
		IntArray:   intArray,
	}
	optionalParam := rand.Int()
	r := ComplexObjectResponse{
		Object1:       o1,
		Object2:       o2,
		OptionalParam: &optionalParam,
	}
	return ctx.JSON(http.StatusOK, r)
}

// Get endpoint defined which simply gives response indicating that POST should be used
// (GET /body-param/error-response)
func (*Api) GetBodyParamErrorResponse(ctx echo.Context) error {
	str := "GET not implemented on this endpoint - please use POST instead."
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusNotImplemented, r)
}

// An endpoint which takes an input parameter and returns a HTTP response code and an error response
// (POST /body-param/error-response)
func (*Api) PostBodyParamErrorResponse(ctx echo.Context) error {
	var inputParams InputObject
	err := ctx.Bind(&inputParams)
	if err != nil {
		errMessage := ErrorResponse{
			ErrorNumber: 400,
			ErrorString: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, errMessage)
	}
	str := fmt.Sprintf("Calling this endpoint (PostBodyParamErrorResponse) deliberately generates an error case - called with JSON object {'descriptor': '%v', 'int_val': %v, 'string': '%v'}",
		inputParams.Descriptor, inputParams.IntVal, inputParams.String)
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusTeapot, r)
}
