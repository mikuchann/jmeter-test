NAME=sample

build: clean
	GOOS=linux go build -o bin/sample

all: build

clean:
	rm -rfv bin
	
	
image:
	docker build -t sample .