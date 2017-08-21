//go:generate protoc -I../../../protobuf -I. --go_out=plugins=micro:. ./proto/service1.proto
//go:generate protoc -I../../../protobuf -I. --microgw_out=:. ./proto/service1.proto
package main

import (
	"golang.org/x/net/context"

	proto "github.com/ka2n/micro-gateway/example/service/service1/proto"
	micro "github.com/micro/go-micro"
)

type handler struct{}

func (h *handler) Hello(ctx context.Context, req *proto.Request, resp *proto.Response) error {
	name := req.GetName()
	resp.Name = "Hello " + name
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.example.s1"),
	)

	srv.Init()
	proto.RegisterSayHandler(srv.Server(), new(handler))

	srv.Run()
}
