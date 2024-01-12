/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"stray/internal/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	infrav1 "stray/gen/api/infra/v1"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tool",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Tools tests")
		testTool()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.core.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".core" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".core")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func unaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("Interceptor\n")
		start := time.Now()
		cos := runtime.GOOS
		ctx = metadata.AppendToOutgoingContext(ctx, "client-os", cos)
		err := invoker(ctx, method, req, reply, cc, opts...)
		end := time.Now()
		fmt.Printf("RPC: %s,,client-OS: '%v' req:%v start time: %s, end time: %s, err: %v", method, cos, req,
			start.Format(time.RFC3339), end.Format(time.RFC3339), err)
		return err
	}
}

func testTool() {
	cancelCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	dialOpts := []grpc.DialOption{grpc.WithBlock()}
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(unaryClientInterceptor()))

	rpcSrvAddr := net.JoinHostPort("127.0.0.1", "8081")
	conn, err := grpc.DialContext(cancelCtx, rpcSrvAddr, dialOpts...)
	if err != nil {
		fmt.Printf("dial error %s", err.Error())
		return
	}

	rpcClient := infrav1.NewGreeterClient(conn)

	rep, err := rpcClient.SayHello(context.Background(), &infrav1.HelloRequest{
		Name: "333",
	})
	if err != nil {
		fmt.Printf("rpc error %s", err.Error())
	} else {
		fmt.Printf("msg %s", rep.GetMessage())
	}

	rep, err = rpcClient.SayHello(context.Background(), &infrav1.HelloRequest{
		Name: "333",
	})
	if err != nil {
		fmt.Printf("2: rpc error %s", err.Error())
	} else {
		fmt.Printf("2: msg %s", rep.GetMessage())
	}

	go func() {
		<-cancelCtx.Done()
		fmt.Printf("closing gRPC client connection: %v", cancelCtx.Err())
		if cerr := conn.Close(); cerr != nil {
			fmt.Printf("Failed to close gRPC conn to %s: %v", rpcSrvAddr, cerr)
		}
	}()

}
