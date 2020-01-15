
RELEASE?=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "latest")
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD?=$(shell date +"%Y-%m-%dT%H:%M:%S%:z")
PROJECT=github.com/tianhanfangyan/grpc-app

GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

.PHONY: vendor
vendor:
	go mod vendor

fmt:
	@gofmt -w ${GOFILES}

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

app:
	go build \
           -mod=vendor \
           -ldflags "-s -w \
           -X ${PROJECT}/version.Release=${RELEASE} \
           -X ${PROJECT}/version.Commit=${COMMIT} \
           -X ${PROJECT}/version.Build=${BUILD}" \
           -o ./grpc-app \
           ./main.go

image: grpc
	docker build -t grpc-app \
           --build-arg RELEASE=${RELEASE} \
           --build-arg COMMIT=${COMMIT} \
           --build-arg BUILD=${BUILD} \
           --build-arg GOPROXY=https://goproxy.cn \
           .

doc: grpc-swagger
	docker run -d --name grpc-app-doc \
           -p 8080:8080 \
           -v ${PWD}/api/swagger:/data \
           -e "SWAGGER_JSON=/data/student.swagger.json" \
           -e "VALIDATOR_URL=null" \
           -e 'DEFAULT_MODELS_EXPAND_DEPTH=-1' \
           swaggerapi/swagger-ui:v3.22.0
