package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	fm "test.com/go-favoritemovie-grpc/favoritemovie"
)

const (
	address = "localhost:8087"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := fm.NewFavoriteMovieSearcherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SearchMovie(ctx, &fm.MovieSpec{Searchword: "blackpink", Pagination: 1})
	if err != nil {
		logrus.Errorf("Ini error %v", err)
		return
	}
	done := make(chan bool)

	go func() {
		for {
			resp, err := response.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Title)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished stream")
}
