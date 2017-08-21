package genswagger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/ka2n/micro-gateway/protoc-gen-microgw/descriptor"
	gen "github.com/ka2n/micro-gateway/protoc-gen-microgw/generator"
)

type generator struct {
	reg *descriptor.Registry
}

// New returns a new generator which generates grpc gateway files.
func New(reg *descriptor.Registry) gen.Generator {
	return &generator{reg: reg}
}

func (g *generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	var files []*plugin.CodeGeneratorResponse_File
	for _, file := range targets {
		glog.V(1).Infof("Processing %s", file.GetName())
		code, err := applyTemplate(param{File: file, reg: g.reg})
		if err != nil {
			return nil, err
		}

		var formatted bytes.Buffer
		json.Indent(&formatted, []byte(code), "", "  ")

		name := file.GetName()
		ext := filepath.Ext(name)
		base := strings.TrimSuffix(name, ext)
		output := fmt.Sprintf("%s.swagger.json", base)
		files = append(files, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(output),
			Content: proto.String(formatted.String()),
		})
		glog.V(1).Infof("Will emit %s", output)
	}
	return files, nil
}
