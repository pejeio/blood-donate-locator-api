SHELL = /bin/bash
PACKAGE ?= blood-donate-locator-api

.SILENT:
.ONESHELL:
.NOTPARALLEL: ;

lint:
	golangci-lint run -c ./.golangci.yaml
