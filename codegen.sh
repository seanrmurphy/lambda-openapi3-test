#! /usr/bin/env bash

cd api

oapi-codegen -config server.cfg.yaml ../api.yaml
oapi-codegen -config types.cfg.yaml ../api.yaml

