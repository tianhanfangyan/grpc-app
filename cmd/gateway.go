package cmd

import (
	"context"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/tianhanfangyan/grpc-app/pkg/api/impl"
	cmdgateway "github.com/tianhanfangyan/grpc-app/pkg/cmd/gateway"
)

var gatewayOption cmdgateway.Option

var gatewayCmd = &cobra.Command{
	Use:              "gateway",
	Short:            "gRPC HTTP gateway",
	PersistentPreRun: persistentPreRun,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := impl.WithSignals(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
		return cmdgateway.Run(ctx, &gatewayOption)
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)

	flags := gatewayCmd.Flags()
	flags.IntVar(&gatewayOption.Port, "port", 33235, "Gateway HTTP server listening port")
	flags.StringVar(&gatewayOption.Endpoint.Host, "endpoint-host", "127.0.0.1", "gRPC server host")
	flags.IntVar(&gatewayOption.Endpoint.Port, "endpoint-port", 33234, "gRPC server port")
	flags.BoolVar(&gatewayOption.Endpoint.Insecure, "endpoint-insecure", false, "insecure")
	flags.BoolVar(&gatewayOption.Endpoint.SkipVerify, "endpoint-skipverify", false, "skip certificate verify")
	flags.StringVar(&gatewayOption.Endpoint.CAFile, "endpoint-ca-file", "", "ca file")
	flags.StringVar(&gatewayOption.Endpoint.CertFile, "endpoint-cert-file", "", "cert file")
	flags.StringVar(&gatewayOption.Endpoint.KeyFile, "endpoint-key-file", "", "key file")
}
