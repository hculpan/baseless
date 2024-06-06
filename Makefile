
clean:
	rm -f baseless

run:
	go run cmd/cli/*.go

build: clean
	go build -o baseless cmd/cli/*.go

test:
	go test tests/*/***

test-verbose:
	go test -v tests/*/***
