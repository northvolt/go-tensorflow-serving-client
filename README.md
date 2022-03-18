# go-tensorflow-serving-client

Go language client wrappers for Tensorflow Server generated from the protobuf files.

# Why This Is Needed

In order to call Tensorflow Server from Go using GPRC, you need to have Go wrappers to match the .proto files for the version of the Tensorflow Server and also the Tensorflow framework.

The Tensorflow framework has all of the needed `option` directives to generate Go wrappers that can be used from a client application using the standard `protoc` tools.

However Tensorflow Server does not yet have the needed `option` directives in its own .proto files, so Go wrapper generation is made somewhat more difficult.

See this PR which adds the needed directives to Tensorflow Server upstream repo: https://github.com/tensorflow/serving/pull/1988

This repo includes all the needed client wrappers already generated, so that they can be used more easily from Go applications.

# Installing to use in another project

    go get -d github.com/northvolt/go-tensorflow-serving-client

If you want to install for a specific version of Tensorflow Server (currently only v2.6.3 is supported):

    go get -d github.com/northvolt/go-tensorflow-serving-client@v2.6.3

# Using

See example program in ./examples/basic/main.go

# Rebuilding

You will need to install protoc and the needed plugins:

        brew install protoc
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

To rebuild the wrappers, for example to update for a new version:

        make update-tf-server-wrappers

# Thanks

The information and scripts in this repo are heavily based on the knowledge learned by @AlexanderJLiu as documented in the https://github.com/AlexanderJLiu/tensorflow-serving-api repo.
