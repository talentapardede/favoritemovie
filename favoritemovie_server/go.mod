module test.com/favoritemovie_server

go 1.16

replace test.com/go-favoritemovie-grpc/favoritemovie v0.0.0 => ../../favoritemovie/impl

replace test.com/go-favoritemovie-grpc/favoritemovie/favoritemovie_server v0.0.0 => ../favoritemovie_server

require (
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.40.0
	test.com/go-favoritemovie-grpc/favoritemovie v0.0.0
	test.com/go-favoritemovie-grpc/favoritemovie/favoritemovie_server v0.0.0
)
