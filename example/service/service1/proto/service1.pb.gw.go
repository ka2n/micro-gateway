// Code generated by protoc-gen-grpc-gateway
// source: proto/service1.proto
// DO NOT EDIT!

/*
Package go_micro_example_s1 is a reverse proxy.

It translates micro into RESTful JSON APIs.
*/
package go_micro_example_s1

import (
	"io"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"github.com/ka2n/micro-gateway/protobuf/go/micro/gw"
	"github.com/ka2n/micro-gateway/runtime/helper"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/selector"
	"golang.org/x/net/context"
)

var _ io.Reader
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_Say_Hello_0(ctx context.Context, marshaler runtime.Marshaler, serviceName string, conn client.Client, r *http.Request, pathParams map[string]string, opts ...client.CallOption) (proto.Message, error) {
	var protoReq Request

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, errors.BadRequest("go.micro.api", err.Error())
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, err
	}

	req := conn.NewProtoRequest(serviceName, "Say.Hello", &protoReq)

	var response Response
	if err := conn.Call(ctx, req, &response, opts...); err != nil {
		ce := errors.Parse(err.Error())
		switch ce.Code {
		case 0:
			ce.Code = 500
			ce.Id = "go.micro.api"
			ce.Status = http.StatusText(500)
			ce.Detail = "error during request: " + ce.Detail
		default:
			return nil, ce
		}
	}
	return &response, nil
}

type SayHTTPAuthHandler func(next client.CallFunc, scheme *gw.SecurityScheme) client.CallFunc

var (
	security_jwt = &gw.SecurityScheme{Type: "apiKey", Description: "", Name: "Authorization", In: "header", Terminate: false}
)

// RegisterSayHTTPHandler registers the http handlers for service Say to "mux".

func RegisterSayHTTPHandler(ctx context.Context, mux *runtime.ServeMux, serviceName string, conn client.Client, auth SayHTTPAuthHandler) error {

	if len(serviceName) == 0 {
		serviceName = "go.micro.example.s1"
	}

	opts := []client.CallOption{
		client.WithSelectOption(selector.WithStrategy(selector.Random)),
	}

	opts_Hello_0 := append(opts, client.WithCallWrapper(func(next client.CallFunc) client.CallFunc {
		return auth(next, security_jwt)
	}))

	mux.Handle("GET", pattern_Say_Hello_0, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := helper.RequestToContext(r)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		resp, err := request_Say_Hello_0(ctx, inboundMarshaler, serviceName, conn, r, pathParams, opts_Hello_0...)
		if err != nil {
			if err, ok := err.(*errors.Error); ok {
				http.Error(w, err.Error(), int(err.Code))
				return
			}
			http.Error(w, err.Error(), 500)
			return
		}

		b, _ := outboundMarshaler.Marshal(resp)
		w.Header().Set("Content-Length", strconv.Itoa(len(b)))
		w.Write(b)
	})

	return nil
}

var (
	pattern_Say_Hello_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "service1", "hello", "name"}, ""))
)
