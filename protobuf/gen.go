//go:generate sh -c "protoc -I../../../ -I. --go_out=:. ./go/micro/gw/{annotations.proto,http.proto}"
package protobuf
