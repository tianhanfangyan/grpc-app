package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	api "github.com/tianhanfangyan/grpc-app/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
	"net/http"
	"time"
)

type Option struct {
	Port     int
	Endpoint struct {
		Host       string
		Port       int
		Insecure   bool
		SkipVerify bool
		CAFile     string
		KeyFile    string
		CertFile   string
	}
}

// http server
func Run(ctx context.Context, opt *Option) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	endpoint := fmt.Sprintf("%s:%d", opt.Endpoint.Host, opt.Endpoint.Port)
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return errors.Wrapf(err, "dial gRPC server failed")
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()

	err = api.RegisterStudentServiceHandlerFromEndpoint(ctx, grpcMux, endpoint, opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	mux.Handle("/healthz", healthzHandler(conn))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", opt.Port),
		Handler: grpcMux,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	log.Printf("gRPC-Gateway is listening on port %d", opt.Port)
	return srv.ListenAndServe()
}

func healthzHandler(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("gRPC server is %s", s), http.StatusBadGateway)
		}
	}
}
