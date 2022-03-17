module github.com/northvolt/go-tensorflow-serving-client

go 1.17

replace github.com/tensorflow/serving/tensorflow_serving => ./third_party/github.com/tensorflow/serving/tensorflow_serving

replace github.com/tensorflow/tensorflow/tensorflow => ./third_party/github.com/tensorflow/tensorflow/tensorflow

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/tensorflow/serving/tensorflow_serving v0.0.0-00010101000000-000000000000 // indirect
	github.com/tensorflow/tensorflow/tensorflow v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)
