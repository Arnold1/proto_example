# proto_example

Error:
```
$ go test
Wanted: {
		"Foo": "1",
    	"Bar": 2,
    	"FooBar": "3.1"
	}
, Got: {"Foo":"1","Bar":"2","FooBar":3.1}
--- FAIL: TestWinNotifyJsonToProtobuf (0.00s)
    jsonToProto_test.go:14: %!s(bool=false) != %!s(bool=true)
FAIL
exit status 1
FAIL	github.com/Arnold1/proto_example	0.329s
```
