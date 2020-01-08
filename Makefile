
RELEASE?=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "latest")
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD?=$(shell date +"%Y-%m-%dT%H:%M:%S%:z")
PROJECT=github.com/tianhanfangyan/grpc-app

GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

.PHONY: vendor
vendor:
	go mod vendor

grpc-gen:
	protoc --proto_path=api/proto \
           --proto_path=third_party \
           --go_out=plugins=grpc:pkg/api \
           api/proto/*.proto

grpc-gateway:
	protoc --proto_path=api/proto \
           --proto_path=third_party \
           --grpc-gateway_out=logtostderr=true:pkg/api \
           api/proto/*.proto

grpc-swagger:
	protoc --proto_path=api/proto \
		    --proto_path=third_party \
		    --swagger_out=logtostderr=true:api/swagger \
		    api/proto/*.proto

grpc: grpc-gen grpc-gateway grpc-swagger

fmt:
	@gofmt -w ${GOFILES}
