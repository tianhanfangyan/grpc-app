# grpc-app
This project is a grpc app demo, it uses grpc+grpc-gateway

## Quick Start

#### Download
    go get github.com/tianhanfangyan/grpc-app

#### build
    go build --mod=vendor -ldflags "-s -w" -o ./grpc-app

#### run grpc
    ./grpc-app grpc

#### run gateway
    ./grpc-app gateway

## swagger

![swagger](swagger.png)


## License
* [Apache](LICENSE)