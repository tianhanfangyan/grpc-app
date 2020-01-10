package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/tianhanfangyan/grpc-app/version"
	"log"
)

var globalOpts struct {
	log struct {
		level      int
		timeformat string
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "grpc-app",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func persistentPreRun(cmd *cobra.Command, args []string) {
	log.Printf("server version: %s@%s, built at %s", version.Release, version.Commit, version.Build)
}

func init() {
	cobra.OnInitialize()
	// Here you will define your flags ae'xnd configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().IntVar(&globalOpts.log.level, "log-level", 0, "log level")
	rootCmd.PersistentFlags().StringVar(&globalOpts.log.timeformat, "log-timeformat", time.RFC3339, "log timeformat")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
