// Code generated by protoc-gen-grpc-gateway
// source: proto/login.proto
// DO NOT EDIT!

/*
Package go_micro_example_login is a reverse proxy.

It translates micro into RESTful JSON APIs.
*/
package go_micro_example_login

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

func request_Login_Create_0(ctx context.Context, marshaler runtime.Marshaler, serviceName string, conn client.Client, r *http.Request, pathParams map[string]string, opts ...client.CallOption) (proto.Message, error) {
	var protoReq Request

	if r.ContentLength > 0 {
		if err := marshaler.NewDecoder(r.Body).Decode(&protoReq); err != nil {
			return nil, errors.BadRequest("go.micro.api", err.Error())
		}
	}

	req := conn.NewProtoRequest(serviceName, "Login.Create", &protoReq)

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

var (
	filter_Login_Get_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_Login_Get_0(ctx context.Context, marshaler runtime.Marshaler, serviceName string, conn client.Client, r *http.Request, pathParams map[string]string, opts ...client.CallOption) (proto.Message, error) {
	var protoReq Request

	if err := runtime.PopulateQueryParameters(&protoReq, r.URL.Query(), filter_Login_Get_0); err != nil {
		return nil, errors.BadRequest("go.micro.api", err.Error())
	}

	req := conn.NewProtoRequest(serviceName, "Login.Get", &protoReq)

	var response GetResponse
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

type LoginHTTPAuthHandler func(next client.CallFunc, scheme *gw.SecurityScheme) client.CallFunc

var (
	security_basicAuth = &gw.SecurityScheme{Alias: "basicAuth", Type: "basic", Description: "", Name: "", In: "", Terminate: false}

	security_basicWithoutError = &gw.SecurityScheme{Alias: "basicWithoutError", Type: "basic", Description: "", Name: "", In: "", Terminate: false}
)

// RegisterLoginHTTPHandler registers the http handlers for service Login to "mux".

func RegisterLoginHTTPHandler(ctx context.Context, mux *runtime.ServeMux, serviceName string, conn client.Client, auth LoginHTTPAuthHandler) error {

	if len(serviceName) == 0 {
		serviceName = "go.micro.example.login"
	}

	opts := []client.CallOption{
		client.WithSelectOption(selector.WithStrategy(selector.Random)),
	}

	opts_Create_0 := append(opts, client.WithCallWrapper(func(next client.CallFunc) client.CallFunc {
		return auth(next, security_basicAuth)
	}))

	mux.Handle("POST", pattern_Login_Create_0, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := helper.RequestToContext(r)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		resp, err := request_Login_Create_0(ctx, inboundMarshaler, serviceName, conn, r, pathParams, opts_Create_0...)
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

	opts_Get_0 := append(opts, client.WithCallWrapper(func(next client.CallFunc) client.CallFunc {
		return auth(next, security_basicWithoutError)
	}))

	mux.Handle("GET", pattern_Login_Get_0, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := helper.RequestToContext(r)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		resp, err := request_Login_Get_0(ctx, inboundMarshaler, serviceName, conn, r, pathParams, opts_Get_0...)
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
	pattern_Login_Create_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "login"}, ""))

	pattern_Login_Get_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "me"}, ""))
)
