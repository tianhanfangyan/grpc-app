package grpc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	api "github.com/tianhanfangyan/grpc-app/pkg/api"
	"github.com/tianhanfangyan/grpc-app/pkg/api/impl"
	"google.golang.org/grpc"
	"log"
	"net"
)

//curl -i http://127.0.0.1:8080/v1/student/user_1
//curl -i http://127.0.0.1:8080/v1/student
//curl -i  -X POST -d '{"stu":{"id":2,"name":"sara","age":21,"sex":"female"}}' http://127.0.0.1:8080/v1/student
//curl -i  -X PUT -d '{"stu":{"id":1,"name":"sam","age":30,"sex":"man"}}' http://127.0.0.1:8080/v1/student/user1
//curl -i -X DELETE http://127.0.0.1:8080/v1/student/user1

type Option struct {
	Port int
}

// grpc server
func Run(ctx context.Context, opt *Option) error {
	svc := impl.NewServer()

	server := grpc.NewServer()
	api.RegisterStudentServiceServer(server, svc)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opt.Port))
	if err != nil {
		return errors.Wrapf(err, "establish listener failed")
	}
	defer lis.Close()

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	log.Printf("gRPC server is listening on port %d", opt.Port)
	return server.Serve(lis)
}
