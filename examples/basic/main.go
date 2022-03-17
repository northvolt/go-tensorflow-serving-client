package main

import (
	"context"
	"log"
	"time"

	"github.com/tensorflow/serving/tensorflow_serving/apis/model_go_proto"
	"github.com/tensorflow/serving/tensorflow_serving/apis/predict_go_proto"
	"github.com/tensorflow/serving/tensorflow_serving/apis/prediction_service_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	"google.golang.org/grpc"
)

var (
	// TensorFlow Serving grpc address.
	address = "127.0.0.1:8500"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}

	client := prediction_service_go_proto.NewPredictionServiceClient(conn)

	// Create a grpc request.
	request := &predict_go_proto.PredictRequest{
		ModelSpec: &model_go_proto.ModelSpec{},
		Inputs:    make(map[string]*tensor_go_proto.TensorProto),
	}
	request.ModelSpec.Name = "test-model"
	request.ModelSpec.SignatureName = "serving_default"

	request.Inputs["inputs"] = &tensor_go_proto.TensorProto{
		Dtype: types_go_proto.DataType_DT_INT64,
		TensorShape: &tensor_shape_go_proto.TensorShapeProto{
			Dim: []*tensor_shape_go_proto.TensorShapeProto_Dim{
				&tensor_shape_go_proto.TensorShapeProto_Dim{
					Size: int64(2),
				},
				&tensor_shape_go_proto.TensorShapeProto_Dim{
					Size: int64(31),
				},
			},
		},
		Int64Val: []int64{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		},
	}

	// Send the grpc request.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Predict(ctx, request)
	if err != nil {
		log.Fatalf("couldn't get response: %v", err)
	}
	log.Printf("%+v", response)
}
