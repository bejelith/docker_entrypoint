bin/docker_entrypoint: bin
	go build -o bin/ ./...
	mv bin/cmd bin/docker_entrypoint
bin:
	mkdir -p bin

install: bin/docker_entrypoint

docker: install
	docker build --no-cache -t docker_entrypoint .


.PHONY: install docker
