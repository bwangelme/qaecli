.PHONY: build

build:
	protoc --go_out=. --go-grpc_out=. app.proto

clean:
	rm -rv gen/