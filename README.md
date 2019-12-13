# docker_entrypoint
Templating engine for docker images

## Dependencies
1. make 
2. go 1.11 or newer

## Build
to build the base docker_entrypoing image run  
`$ make docker`

## Usage
### Build a configuration file 
Create a template configuration file `server.properties` like the following  
```
server_address={{ .LISTEN_ADDRESS }}
server_port={{ .LISTEN_PORT }}
id={{ .ID }}
```
### Example dockerfile
The build your image using docker_entrypoint as first stage of a multi-stage build:  
```
FROM docker_entrypoint as entrypoint
FROM alpine #OR YOU FAVOURITE BASE IMAGE
COPY --from=entrypoint /docker_entrypoint /docker_entrypoint
RUN ...
MKDIR ...
ENV SERVER_PROPERTIES_LISTEN_ADDRESS="localhost"
ENV SERVER_PROPERTIES_LISTEN_PORT=8080
ENV SERVER_PROPERTIES_ID=1

ENTRYPOINT ["/docker_entrypoint", "-template", "/etc/server/server.properties", "bin/yoursever.bin", "-server_log_dir /var/log/yourserver"]
```
Remember to pass the `-template` parameter to the `/docker_template` executable then pass your executable as first argument after the template list. All following arguments will be passed to your binary of choice.

Now build your dockerfile and run it!
