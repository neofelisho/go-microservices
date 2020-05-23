# Environment Variables

In this project we use OS environment variables to keep configuration setting.

## Dependency

- [envconfig](https://github.com/kelseyhightower/envconfig)

## Usage

### *.env

Keep variables as key/value pairs, e.g.,

```.env
GMS_DATABASE_USER=postgres
GMS_DATABASE_PASSWORD=postgres
``` 

The supported environments for now are:

```shell script
$ env/env.sh                                                                                                                    27.6s  Sat 23 May 2020 10:43:23 PM UTC
Please specify one of the available environment:
dev
docker
```

### Sourcing Helper

Use `env/env.sh` to source specific environment, notice that working directory is $PROJECT_ROOT.

For bash shell, we can use:

```shell script
# To source the development environment
$ source env/env.sh dev 
```

or:

```shell script
$ source <(env/env.sh dev)
```

For fish shell, we can use:

```shell script
$ env/env.sh dev | source
```
