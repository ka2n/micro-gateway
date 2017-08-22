//go:generate sh -c "cd ./service/login && go generate"
//go:generate sh -c "cd ./service/service1 && go generate"
package main

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"

	login "github.com/ka2n/micro-gateway/example/service/login/proto"
	service1 "github.com/ka2n/micro-gateway/example/service/service1/proto"
	"github.com/ka2n/micro-gateway/protobuf/go/micro/gw"
)

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

func newGateway(ctx context.Context, conn client.Client) (http.Handler, error) {
	mux := runtime.NewServeMux()

	var err error
	err = login.RegisterLoginHTTPHandler(ctx, mux, "", conn, authHandler)
	if err != nil {
		return nil, err
	}

	err = service1.RegisterSayHTTPHandler(ctx, mux, "", conn, authHandler)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

func authHandler(next client.CallFunc, scheme *gw.SecurityScheme) client.CallFunc {
	switch scheme.GetType() {
	case "basic":
		allowGuest := scheme.GetAlias() == "basicWithoutError"
		return func(ctx context.Context, address string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			var user string
			md, ok := metadata.FromContext(ctx)
			if !ok {
				md = make(metadata.Metadata)
			}

			if uname, pass, ok := parseBasicAuth(md["Authorization"]); ok {
				user = getUser(uname, pass)
			}

			if user == "" {
				if allowGuest {
					md["Userid"] = "guest"
				} else {
					return errors.Unauthorized("go.micro.auth", "Unauthorized")
				}
			} else {
				md["Userid"] = user
			}

			return next(metadata.NewContext(ctx, md), address, req, rsp, opts)
		}
	case "apiKey":
		return func(ctx context.Context, address string, req client.Request, rsp interface{}, opts client.CallOptions) error {
			return nil
		}
	default:
		panic("Unsupported security scheme")
	}
}

func getUser(username, password string) string {
	if username == "admin" && password == "password" {
		return username
	}
	return ""
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
