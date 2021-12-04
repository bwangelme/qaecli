.PHONY: cpgen build

cpgen:
	cp -rv ../qae/pb .

build:
	go build -o bin/qae .