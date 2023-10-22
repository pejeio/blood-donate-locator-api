SHELL = /bin/bash
PACKAGE ?= blood-donate-locator-api

.SILENT:
.ONESHELL:
.NOTPARALLEL: ;

lint:
	golangci-lint run -c ./.golangci.yaml

dev:
	air

local-dev-start:
	$(MAKE) -C ./local-dev-scripts start

local-dev-stop:
	$(MAKE) -C ./local-dev-scripts stop
