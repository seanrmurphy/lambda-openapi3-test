#! /usr/bin/env bash

cd api

oapi-codegen.original -config server.cfg.yaml ../api.yaml
oapi-codegen.original -config types.cfg.yaml ../api.yaml

