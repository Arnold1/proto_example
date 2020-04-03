package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	proto "github.com/Arnold1/proto_example/proto"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func areEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func TestJsonToProtobuf(t *testing.T) {

	jsonDataIn := []byte(`{
		"Foo": "1",
    	"Bar": 2,
    	"FooBar": 3.1
	}`)
	expectedJsonDataOut := []byte(`{
		"Foo": "1",
    	"Bar": "2",
    	"FooBar": 3.1
	}`)

	out, err := JsonToProto(bytes.NewReader(jsonDataIn))
	assertEqual(t, err, nil)

	data, err := out.Marshal()
	assertEqual(t, err, nil)

	var reqOut = &proto.Message{}
	err = reqOut.Unmarshal(data)
	assertEqual(t, err, nil)

	jsonDataOut, err := ProtoToJSON(reqOut)
	assertEqual(t, err, nil)

	res, err := areEqualJSON(string(expectedJsonDataOut), string(jsonDataOut))
	if res != true {
		fmt.Printf("Wanted: %s\n, Got: %s\n", string(jsonDataIn), string(jsonDataOut))
	}
	assertEqual(t, err, nil)
	assertEqual(t, res, true)
}
