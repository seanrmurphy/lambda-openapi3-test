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
	ctx.Bind(&inputParams)
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
func (*Api) PostBodyParamComplexResponse(ctx echo.Context) error {
	var inputParams InputObject
	ctx.Bind(&inputParams)
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
	ctx.Bind(&inputParams)
	str := fmt.Sprintf("Calling this endpoint (PostBodyParamErrorResponse) deliberately generates an error case - called with JSON object {'descriptor': '%v', 'int_val': %v, 'string': '%v'}",
		inputParams.Descriptor, inputParams.IntVal, inputParams.String)
	randomInt := rand.Int()

	r := ErrorResponse{
		ErrorNumber: randomInt,
		ErrorString: str,
	}
	return ctx.JSON(http.StatusTeapot, r)
}

//Here, we implement all of the handlers in the ServerInterface
//func (p *PetStore) FindPets(ctx echo.Context, params FindPetsParams) error {
//p.Lock.Lock()
//defer p.Lock.Unlock()

//var result []Pet

//for _, pet := range p.Pets {
//if params.Tags != nil {
//If we have tags,  filter pets by tag
//for _, t := range *params.Tags {
//if pet.Tag != nil && (*pet.Tag == t) {
//result = append(result, pet)
//}
//}
//} else {
//Add all pets if we're not filtering
//result = append(result, pet)
//}

//if params.Limit != nil {
//l := int(*params.Limit)
//if len(result) >= l {
//We're at the limit
//break
//}
//}
//}
//return ctx.JSON(http.StatusOK, result)
//}

//func (p *PetStore) AddPet(ctx echo.Context) error {
//We expect a NewPet object in the request body.
//var newPet NewPet
//err := ctx.Bind(&newPet)
//if err != nil {
//return sendPetstoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
//}
//We now have a pet, let's add it to our "database".

//We're always asynchronous, so lock unsafe operations below
//p.Lock.Lock()
//defer p.Lock.Unlock()

//We handle pets, not NewPets, which have an additional ID field
//var pet Pet
//pet.Name = newPet.Name
//pet.Tag = newPet.Tag
//pet.Id = p.NextId
//p.NextId = p.NextId + 1

//Insert into map
//p.Pets[pet.Id] = pet

//Now, we have to return the NewPet
//err = ctx.JSON(http.StatusCreated, pet)
//if err != nil {
//Something really bad happened, tell Echo that our handler failed
//return err
//}

//Return no error. This refers to the handler. Even if we return an HTTP
//error, but everything else is working properly, tell Echo that we serviced
//the error. We should only return errors from Echo handlers if the actual
//servicing of the error on the infrastructure level failed. Returning an
//HTTP/400 or HTTP/500 from here means Echo/HTTP are still working, so
//return nil.
//return nil
//}

//func (p *PetStore) FindPetById(ctx echo.Context, petId int64) error {
//p.Lock.Lock()
//defer p.Lock.Unlock()

//pet, found := p.Pets[petId]
//if !found {
//return sendPetstoreError(ctx, http.StatusNotFound,
//fmt.Sprintf("Could not find pet with ID %d", petId))
//}
//return ctx.JSON(http.StatusOK, pet)
//}

//func (p *PetStore) DeletePet(ctx echo.Context, id int64) error {
//p.Lock.Lock()
//defer p.Lock.Unlock()

//_, found := p.Pets[id]
//if !found {
//return sendPetstoreError(ctx, http.StatusNotFound,
//fmt.Sprintf("Could not find pet with ID %d", id))
//}
//delete(p.Pets, id)
//return ctx.NoContent(http.StatusNoContent)
//}
