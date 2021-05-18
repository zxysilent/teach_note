#!/bin/bash

name="noteapi"
go build -tags=prod -o $name.exe main.go