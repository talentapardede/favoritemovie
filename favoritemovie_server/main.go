package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	fm "test.com/go-favoritemovie-grpc/favoritemovie"
	impl "test.com/go-favoritemovie-grpc/favoritemovie/favoritemovie_server/impl"
)

const (
	port = ":8087"
)

func main() {
	listenAddr, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("Failed to listen : %v", err)
	}

	service := grpc.NewServer()
	fm.RegisterFavoriteMovieSearcherServer(service, &impl.FavoriteMovieServer{})
	logrus.Printf("server listening at %v", listenAddr.Addr())
	err = service.Serve(listenAddr)
	if err != nil {
		logrus.Fatalf("fatal to serve: %v", err)
	}

}
