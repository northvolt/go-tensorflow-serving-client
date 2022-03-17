# go-tensorflow-serving-client

Go language client wrapper for Tensorflow Server generated from protobuf files.

# Installing to use in another project

    go get -d github.com/northvolt/go-tensorflow-serving-client

# Using

See example program in ./examples/basic/main.go

# Rebuilding

You will need to install protoc and the needed plugins:

        brew install protoc
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

To rebuild the wrappers, for example to update for a new version:

        make update-tf-server-wrappers
