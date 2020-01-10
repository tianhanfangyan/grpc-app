package cmd

import (
	"context"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/tianhanfangyan/grpc-app/pkg/api/impl"

	cmdgrpc "github.com/tianhanfangyan/grpc-app/pkg/cmd/grpc"
)

var grpcOptions cmdgrpc.Option

var grpcCmd = &cobra.Command{
	Use:              "grpc",
	Short:            "gRPC server",
	PersistentPreRun: persistentPreRun,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := impl.WithSignals(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
		return cmdgrpc.Run(ctx, &grpcOptions)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	grpcCmd.Flags().IntVar(&grpcOptions.Port, "port", 33234, "gRPC server listening port")

}
