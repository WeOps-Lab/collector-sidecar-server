prebuild:
	swag init
	wire
test:
	go test -v  -cover -coverprofile=coverage.out ./...|go-test-report