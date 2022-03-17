
utest:
	go mod tidy; \
	export PROJECT_ENV="unit" && go test -coverpkg=./... -coverprofile=coverage.data ./... -gcflags=-l ;
