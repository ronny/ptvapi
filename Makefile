all: examples

# needs Go 1.13 to be able to use -o with a directory
examples:
	env CGO_ENABLED=0 go build -o ./bin/ ./examples/...

test: vet
	env ENV=test go test -race -cover ./...

vet:
	go vet ./...

regenerate-client:
	# You will need go-swagger to be installed first:
	#
	#		https://goswagger.io/install.html
	#
	swagger generate client --name "PTV Timetable API v3" --additional-initialism PTV

.PHONY: all examples test vet generate regenerate-client
