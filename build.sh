#!/bin/bash

## TODO: Replace me with a very small Makefile

cd 1.5
docker build -t technosophos/goglide:1.5-0 .
cd onbuild
docker build -t technosophos/goglide:1.5-0-onbuild .
cd ../../example
docker build -t goglide-example .
