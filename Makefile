setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/google/wire/cmd/wire@latest

prepare:
	swag init
	wire

test:
	go test -v  -cover -coverprofile=coverage.out ./...|go-test-report

push:
	git add . && codegpt commit . && git push