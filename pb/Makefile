.PHONY: build clean

build: clean
	protoc --go_out=. --go-grpc_out=. app.proto commit.proto

clean:
	rm -rv gen/