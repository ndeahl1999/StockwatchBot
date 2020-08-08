# go base commands for build, clean, test, and get
GOCMD=go
GOGET=${GOCMD} get
GORUN=${GOCMD} run
GOBUILD=${GOCMD} build
GOCLEAN=${GOCMD} clean
GOTEST=${GOCMD} test
GOINSTALL=${GOCMD} install
GODOC=${GOCMD} doc
GOCOVER=${GOCMD} tool cover
GOMOD=${GOCMD} mod

MAIN_FILE=cmd/main.go

lint:
	@echo "******************* Starting linting execution *******************"

	# Run lint
	${GOLINTCMD}

	@echo "******************* Ending linting execution **********************"

run-dev: lint
	@echo "******************* Starting development server *******************"

	${GORUN} ${MAIN_FILE}	
	
	@echo "******************* Stopping development server ********************"