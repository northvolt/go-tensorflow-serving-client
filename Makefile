TENSORFLOW_VERSION := v2.6.3
TENSORFLOW_SERVER_VERSION := go-package-proto-2.6.3

build/tensorflow:
	mkdir -p ./build/
	git clone https://github.com/tensorflow/tensorflow.git ./build/tensorflow

update-tensorflow: build/tensorflow
	cd ./build/tensorflow && git fetch && git checkout $(TENSORFLOW_VERSION)

# TODO: change to upstream fork once PR #1988 has been merged:
# https://github.com/tensorflow/serving/pull/1988
build/serving:
	mkdir -p ./build/
	git clone https://github.com/northvolt/serving.git ./build/serving

update-serving: build/serving
	cd ./build/serving && git fetch && git checkout $(TENSORFLOW_SERVER_VERSION)

gen-tf-wrappers:
	mkdir -p ./third_party
	protoc -I=./build/serving -I=./build/tensorflow --go_out=./third_party --go-grpc_out=./third_party ./build/serving/tensorflow_serving/*/*.proto
	protoc -I=./build/serving -I=./build/tensorflow --go_out=./third_party ./build/serving/tensorflow_serving/sources/storage_path/*.proto
	protoc -I=./build/serving -I=./build/tensorflow --go_out=./third_party ./build/tensorflow/tensorflow/core/framework/*.proto
	protoc -I=./build/serving -I=./build/tensorflow --go_out=./third_party ./build/tensorflow/tensorflow/core/example/*.proto
	protoc -I=./build/serving -I=./build/tensorflow --go_out=./third_party ./build/tensorflow/tensorflow/core/protobuf/*.proto
	protoc -I=./build/serving -I=./build/tensorflow --go_out=./third_party ./build/tensorflow/tensorflow/stream_executor/*.proto

# adds files needed to correctly compile things
third_party/github.com/tensorflow/tensorflow/tensorflow/go/genop/internal/doc.go:
	mkdir -p ./third_party/github.com/tensorflow/tensorflow/tensorflow/go/genop ./third_party/github.com/tensorflow/tensorflow/tensorflow/go/genop/internal
	echo "package internal" >> ./third_party/github.com/tensorflow/tensorflow/tensorflow/go/genop/internal/doc.go

third_party/github.com/tensorflow/tensorflow/tensorflow/go/op/wrappers.go:
	mkdir -p ./third_party/github.com/tensorflow/tensorflow/tensorflow/go/op
	cp ./build/tensorflow/tensorflow/go/op/wrappers.go ./third_party/github.com/tensorflow/tensorflow/tensorflow/go/op/wrappers.go

third_party/github.com/tensorflow/tensorflow/tensorflow/go/doc.go:
	mkdir -p ./third_party/github.com/tensorflow/tensorflow/tensorflow/go
	echo "package tensorflow" >> ./third_party/github.com/tensorflow/tensorflow/tensorflow/go/doc.go

third_party/github.com/tensorflow/tensorflow/tensorflow/go.mod:
	cd third_party/github.com/tensorflow/tensorflow/tensorflow && go mod init github.com/tensorflow/tensorflow/tensorflow

tf-fixes-needed: third_party/github.com/tensorflow/tensorflow/tensorflow/go/genop/internal/doc.go third_party/github.com/tensorflow/tensorflow/tensorflow/go/op/wrappers.go third_party/github.com/tensorflow/tensorflow/tensorflow/go/doc.go third_party/github.com/tensorflow/tensorflow/tensorflow/go.mod

third_party/github.com/tensorflow/serving/tensorflow_serving/go.mod:
	cd third_party/github.com/tensorflow/serving/tensorflow_serving && go mod init github.com/tensorflow/serving/tensorflow_serving

serving-fixes-needed: third_party/github.com/tensorflow/serving/tensorflow_serving/go.mod

update-tf-server-wrappers: update-tensorflow update-serving gen-tf-wrappers tf-fixes-needed serving-fixes-needed
