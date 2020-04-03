package main

import (
	"bytes"

	proto "github.com/Arnold1/proto_example/proto"
	"github.com/gogo/protobuf/jsonpb"
)

func JsonToProto(r *bytes.Reader) (*proto.Message, error) {
	var req = &proto.Message{}
	unmarshaler := &jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err := unmarshaler.Unmarshal(r, req); err != nil {
		return nil, err
	}
	return req, nil
}

func ProtoToJSON(res *proto.Message) ([]byte, error) {
	var b bytes.Buffer

	ma := jsonpb.Marshaler{}
	if err := ma.Marshal(&b, res); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
