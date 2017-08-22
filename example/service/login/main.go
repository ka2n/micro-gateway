//go:generate protoc -I../../../protobuf -I. --go_out=plugins=micro:. ./proto/login.proto
//go:generate protoc -I../../../protobuf -I. --microgw_out=:. ./proto/login.proto
package main

import (
	"encoding/base64"
	"strings"

	"github.com/micro/go-micro/metadata"

	"github.com/golang/glog"

	"golang.org/x/net/context"

	proto "github.com/ka2n/micro-gateway/example/service/login/proto"
	micro "github.com/micro/go-micro"
)

type handler struct{}

func (h *handler) Create(ctx context.Context, req *proto.Request, resp *proto.Response) error {
	user := userFromContext(ctx)
	resp.JwtToken = user // TODO: generate token
	return nil
}

func (h *handler) Get(ctx context.Context, req *proto.Request, resp *proto.GetResponse) error {
	resp.Hello = "Hello " + userFromContext(ctx) + "!"
	return nil
}

func userFromContext(ctx context.Context) string {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return ""
	}
	return md["Userid"]
}

func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.example.login"),
	)

	srv.Init()
	proto.RegisterLoginHandler(srv.Server(), new(handler))

	if err := srv.Run(); err != nil {
		glog.Fatal(err)
	}
}
