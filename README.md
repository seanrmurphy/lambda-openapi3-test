# What

A small application to demonstrate how an OpenAPI v3.0 API can be implemented
in Go and deployed to AWS using the HTTP API.

Check out [this medium post FIXME]() for more information.

# Requirements

- Modern [`go`](https://golang.org/) toolchain - I used go 1.14.6
- [`oapi-codegen`](https://github.com/deepmap/oapi-codegen)
- AWS [SAM](https://aws.amazon.com/serverless/sam/) application management tool
- Basic utilities: `zip` for function upload to AWS, `curl` and `jq` for testing

# Install

Clone this repo somewhere on your filesystem.

```
git clone github.com/seanrmurphy/lambda-openapi3-test
```

# Build, deploy, test, clean up

- Generate the types and server-side stub code

```
cd api
go generate
```

This generates two files `api-server.gen.go` and `api-types.gen.go` which contain
the server code and the types defined in the OAS3 specification respectively.

- Build the application

This application supports 2 modes: one in which the application runs locally and
one in which the application runs in AWS. `main.go` provides a boolean called `lambdaMode`
which determines whether the resulting binary should be run locally or as a Lamdba
function. Set `lambdaMode` to false to build a binary which can offer the API
locally.

```
go build
```

This builds the application in the top level directory of the repo.

- Run the application locally

If the application was built with `lambdaMode` set to `false`, you can run the
binary directly

```
./lambda-openapi3-test
```

You can then run the `curl` based tests in the `tests` directory.

- Deploy the application to AWS

If the binary is compiled with `lambdaMode` set to true, then it can be deployed
to AWS. Assuming your AWS credentials are configured and SAM is set up

```
cd deploy
./deploy.sh
```

This zips the executable and uploads it to AWS using the `sam` tool and
associated configuration file (`sam-deploy.toml`). On successful completion, it
prints the identifier for the Lambda function and the endpoint of the HTTP API.

- Run the tests

```
cd ../tests
export RESTAPI=<ENDPOINT OUTPUT FROM DEPLOYMENT PROCESS>
./tests.sh
```

The tests demonstrate both successful and unsuccessful calling of all of the
endpoints defined in the API.

- Removing the API

The easiest way to remove the application is via the AWS CloudFormations
dashboard.


# Improved Request Validation

As noted in the medium post, oapi-codegen does not perform validation of the
body of requests. This should be evident from the test suite provided.


One way of dealing with that is to apply [this pull request](https://github.com/deepmap/oapi-codegen/pull/245). This
makes a new capability available to the oapi-codegen tool, that of validating
the unmarshall of objects. It operates by creating an `UnmarshalJSON()` function
for schemas which are used as input parameters. In this case, you can achieve
the same by copying the code below to the `api-types.gen.go` file and rebuilding.

Now, if you run the API locally, the tests generate an error when and incorrect
JSON object is supplied.


```go
func (t *InputObject) UnmarshalJSON(data []byte) error {
	inner := struct {
		Descriptor *string `json:"descriptor,omitempty"`
		IntVal     *int    `json:"int_val,omitempty"`
		String     *string `json:"string,omitempty"`
	}{}
	err := json.Unmarshal(data, &inner)
	if err != nil {
		return err
	}
	if inner.Descriptor == nil {
		return errors.New("descriptor is required")
	}
	if inner.IntVal == nil {
		return errors.New("int_val is required")
	}
	if inner.String == nil {
		return errors.New("string is required")
	}
	*t = InputObject{
		Descriptor: *inner.Descriptor,
		IntVal:     *inner.IntVal,
		String:     *inner.String,
	}
	return nil
}
```

