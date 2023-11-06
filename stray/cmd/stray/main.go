package main

import (
	"context"
	"log"
	"net"
	"net/http"

	hello "stray/generated/api/proto/v1"
	// swaggerdata "stray/generated/swagger/data"
	// swaggerui "stray/generated/swagger/ui"
	// assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: in.Name + " test123 "}, nil
}

func localTest() {
	var x int
	x = 1

	defer func() {
		log.Println("x (defer) =", x)
	}()

	log.Println("x =", x)
	x = 101
}

func main() {
	localTest()

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(listen))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = hello.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	// datafs := &assetfs.AssetFS{
	// 	Asset:    swaggerdata.Asset,
	// 	AssetDir: swaggerdata.AssetDir,
	// }
	// serveFileFn := func(w http.ResponseWriter, r *http.Request, root *assetfs.AssetFS, path string) error {
	// 	file, err := root.Open(path)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer file.Close()
	// 	http.ServeContent(w, r, filepath.Base(path), time.Time{}, file)
	// 	return nil
	// }

	// uiServer := http.FileServer(&assetfs.AssetFS{
	// 	Asset:    swaggerui.Asset,
	// 	AssetDir: swaggerui.AssetDir,
	// 	Prefix:   "swagger-ui",
	// 	Fallback: "index.html",
	// })
	// gwmux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", uiServer))

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
