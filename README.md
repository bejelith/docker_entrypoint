# docker_entrypoint
Templating engine for docker images


## Usage
### Template file
Create a template file in golang template format and save it as file.template  
```a={{ .VAR }}```

### Run

`$ export file_VAR=value`  
`$ cmd -template file.template /bin/cat file`
