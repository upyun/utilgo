#!/bin/sh

goimports -w .
golint ./...
