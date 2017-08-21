//go:generate sh -c "cd ./service/login && go generate"
//go:generate sh -c "cd ./service/service1 && go generate"
package main

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"

	login "github.com/ka2n/micro-gateway/example/service/login/proto"
	service1 "github.com/ka2n/micro-gateway/example/service/service1/proto"
)

func newGateway(ctx context.Context, conn client.Client) (http.Handler, error) {
	mux := runtime.NewServeMux()

	var err error
	err = login.RegisterLoginHTTPHandler(ctx, mux, "", conn)
	if err != nil {
		return nil, err
	}

	err = service1.RegisterSayHTTPHandler(ctx, mux, "", conn)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	svc := micro.NewService(
		micro.Name("go.micro.gw"),
	)
	svc.Init()

	gw, err := newGateway(ctx, svc.Client())
	if err != nil {
		glog.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		glog.Fatal(err)
	}
}
