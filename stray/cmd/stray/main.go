package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"time"

	hello "stray/generated/api/proto/v1"
	swaggerdata "stray/generated/swagger/data"
	swaggerui "stray/generated/swagger/ui"

	assetfs "github.com/elazarl/go-bindata-assetfs"

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
	log.Println("get message ")
	return &hello.HelloReply{Message: in.Name + " test123 "}, nil
}

func clientTest() {
	go func() {
		time.Sleep(8 * time.Second)
		log.Println("sending abc")
		cc, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalln("Failed to dial server:", err)
		}
		client := hello.NewGreeterClient(cc)
		client.SayHello(context.Background(), &hello.HelloRequest{
			Name: "abc",
		})
	}()
}

func main() {
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

	log.Println("testing")
	cc, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = hello.RegisterGreeterHandler(context.Background(), gwmux, cc)
	if err != nil {
		log.Fatalln("Failed to gwmux:", err)
	}

	mux := http.NewServeMux()

	datafs := &assetfs.AssetFS{
		Asset:    swaggerdata.Asset,
		AssetDir: swaggerdata.AssetDir,
	}

	serveFileFn := func(w http.ResponseWriter, r *http.Request, root *assetfs.AssetFS, path string) error {
		file, err := root.Open(path)
		if err != nil {
			log.Println("error ", err)
			return err
		}
		defer file.Close()
		http.ServeContent(w, r, filepath.Base(path), time.Time{}, file)
		return nil
	}

	mux.HandleFunc("/swagger-ui/LICENSE", func(w http.ResponseWriter, r *http.Request) {
		log.Println("license request ===")
		err := serveFileFn(w, r, datafs, "LICENSE")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to open file: %s", err.Error()), http.StatusBadRequest)
		}
	})

	mux.HandleFunc("/swagger-ui/document", func(w http.ResponseWriter, r *http.Request) {
		err := serveFileFn(w, r, datafs, "document")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to open file: %s", err.Error()), http.StatusBadRequest)
		}
	})

	mux.HandleFunc("/swagger-ui/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		err := serveFileFn(w, r, datafs, "openapi.yaml")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to open file: %s", err.Error()), http.StatusBadRequest)
		}
	})

	uiServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swaggerui.Asset,
		AssetDir: swaggerui.AssetDir,
		Prefix:   "swagger-ui",
		Fallback: "index.html",
	})
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", uiServer))
	mux.Handle("/", gwmux)

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
