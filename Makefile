.PHONY:
.SILENT:

run:
	go run cmd/main.go

build:
	go build -o ./bin/lms-system cmd/main.go