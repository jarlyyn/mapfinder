#!/bin/bash

CGO_ENABLED=0 go build -tags 'netgo' --trimpath -o ../../bin/app ../