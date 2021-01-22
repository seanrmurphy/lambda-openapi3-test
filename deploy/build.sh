#! /usr/bin/env bash

cd ..

echo
echo "Building application..."
go build


echo
echo "Moving application to this directory and zipping for Lambda upload..."
mv lambda-openapi3-test deploy
cd deploy
zip lambda-openapi3-test.zip ./lambda-openapi3-test
rm lambda-openapi3-test


