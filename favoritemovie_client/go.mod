module test.com/favoritemovie_client

go 1.16

replace test.com/go-favoritemovie-grpc/favoritemovie v0.0.0 => ../../favoritemovie/impl

require (
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.40.0
	test.com/go-favoritemovie-grpc/favoritemovie v0.0.0
)
