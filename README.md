# docker_entrypoint
Templating engine for docker images

Did you have to spend time in finding out to configure a random docker image you just downloaded from Docker registry?  

Wouldn't be nice having the same ENV variable be the same across images without the need of replacing entirelly the whole congiguration file or having to dig `docker inspect` and just run an help command?

Docker entrypoint manages your configuration file templates rendering them from environemnt variables values. 

For example, setting Kafka's broker ID would be as simple as setting an environment variable  
`docker run -d --name kafka -e SERVER_PROPERPERTIES_BROKER_ID = 1 mykafka-image`

and gives users an easy way to access all configuration variables with a `help` command:  
`docker run -ti --rm mykafka-image --help

## Dependencies
1. make 
2. go 1.11 or newer

## Build
to build the base docker_entrypoing image run  
`$ make docker`

## Usage
### Build a configuration file 
Create a template configuration file `server.properties.template` like the following  
```
server_address={{ .LISTEN_ADDRESS }}
server_port={{ .LISTEN_PORT }}
id={{ .ID }}
```
#### ENVORNMENT VARS Spec
The environment vars hare composed by a PREFIX which identifies the template the variable is referring to and a POSTFIX which represents the varible inside the template eg.

To render a template named testfile.config.template containing a variable named ID the corrent ENV variable will need to look like this:
`testfile_config_ID=1`. This is to support multiple templates for a single run, for exaple for Java applications that depend on mulitple `.properties` files (log4j2.properties, server.properties, ...).

Example:  
`docker_entrypoint -template /config/log4j2.properties.template -template /config/server.properties.template java -jar ...`

After the templates are processed the output of `/config/log4j2.properties.template` will be written in `/config/log4j2.properties` for each invocation

### Example dockerfile
The build your image using docker_entrypoint as first stage of a multi-stage build:  
```
FROM docker_entrypoint as entrypoint
FROM alpine #OR YOU FAVOURITE BASE IMAGE
COPY --from=entrypoint /docker_entrypoint /docker_entrypoint
RUN ...
MKDIR ...
ENV server_properties_LISTEN_ADDRESS="localhost"
ENV server_properties_LISTEN_PORT=8080
ENV server_properties_ID=1

ENTRYPOINT ["/docker_entrypoint", "-template", "/etc/server/server.properties.template", "bin/yoursever.bin", "-server_log_dir /var/log/yourserver"]
```
Each template has to be created with a `.template` extension as the entrypoing will remove it when rendering the file.  
Remember to pass the `-template` parameter to the `/docker_template` executable then pass your executable as first argument after the template list. All following arguments will be passed to your binary of choice.

Now build your dockerfile and run it!

# Better code badge
[![BCH compliance](https://bettercodehub.com/edge/badge/bejelith/docker_entrypoint?branch=master)](https://bettercodehub.com/)
