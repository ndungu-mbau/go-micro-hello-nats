package main;

import (
	"context"
	"log"

	pb "github.com/mbau_ndungu/go-micro-hello/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/transport/nats"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	transport := nats.newTransport()
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Transport(transport)
	)

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}