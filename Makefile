all: vet test testrace

build: deps
	go build github.com/e421083458/grpc-proxy/...

clean:
	go clean -i github.com/e421083458/grpc-proxy/...

deps:
	go get -d -v github.com/e421083458/grpc-proxy/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/e421083458/grpc-proxy/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/e421083458/grpc-proxy/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/e421083458/grpc-proxy/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/e421083458/grpc-proxy/...

testdeps:
	go get -d -v -t github.com/e421083458/grpc-proxy/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/e421083458/grpc-proxy/...

updatedeps:
	go get -d -v -u -f github.com/e421083458/grpc-proxy/...

updatetestdeps:
	go get -d -v -t -u -f github.com/e421083458/grpc-proxy/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
