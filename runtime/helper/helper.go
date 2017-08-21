package helper

import (
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

// RequestToContext converts a *http.Request to context.Context
func RequestToContext(r *http.Request) context.Context {
	ctx := r.Context()
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata)
	}

	for k, v := range r.Header {
		md[k] = strings.Join(v, ",")
	}
	return metadata.NewContext(ctx, md)
}

// Strategy is a hack for selection
func Strategy(services []*registry.Service) selector.Strategy {
	return func(_ []*registry.Service) selector.Next {
		// ignore input to this function, use services above
		return selector.Random(services)
	}
}
