# registration-relayer

## Description

[//]: # (TODO: add description)

## Install

  ```
  git clone github.com/rarimo/registration-relayer
  cd registration-relayer
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.

Note: if you are using Gitlab for building project `docs/spec/paths` folder must not be
empty, otherwise only `Build and Publish` job will be passed.  

## Config
```yaml
  network:
    rpc: "" # (url) RPC API endpoint
    contract_address: "" # (hex) target contract address
    vault_address: "http://127.0.0.1:8200" # (url) vault address
    vault_mount_path: "secret_data" # (string)
```
ENV
```
  VAULT_TOKEN (will be cleared after start service)
```

## Running from docker 
  
Make sure that docker installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

  ```
  docker build -t github.com/rarimo/registration-relayer .
  docker run -e KV_VIPER_FILE=/config.yaml github.com/rarimo/registration-relayer
  ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `run service` command



### Third-party services


## Contact

Responsible 
The primary contact for this project is  [//]: # (TODO: place link to your telegram and email)
