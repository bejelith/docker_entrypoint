bin/docker_entrypoint: bin
	cd go/cmd && go build -o ../../$@

bin:
	mkdir bin

install:
	go install ./...

docker:
	docker build --no-cache -t docker_entrypoint .

clean:
	rm -rf bin \
	docker rmi docker_entrypoint:latest
.PHONY: install docker clean
