.PHONY: run

run:
	@go run main.go

test_creational:
	@go test ./creational $(options)

test_structural:
	@go test ./structural $(options)