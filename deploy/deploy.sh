#! /usr/bin/env bash

echo
echo "Copying application to this directory and zipping for Lambda upload..."
cp ../lambda-openapi3-test
zip lambda-openapi3-test.zip ./lambda-openapi3-test
rm lambda-openapi3-test

sam deploy --no-confirm-changeset -t ./lambda-openapi3-test-http-api.yaml

