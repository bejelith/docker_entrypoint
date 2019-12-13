bin/docker_entrypoint: bin
	go build -o $@ ./go/docker_entrypoint

bin:
	mkdir bin

install:
	go install ./...

docker:
	docker build --no-cache -t docker_entrypoint .


.PHONY: install docker
