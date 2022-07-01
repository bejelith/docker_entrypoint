install: bin
	go build -o bin/ ./...
	mv bin/cmd bin/docker_entrypoint
bin:
	mkdir -p bin

docker: install
	docker build --no-cache -t docker_entrypoint .


.PHONY: install docker
