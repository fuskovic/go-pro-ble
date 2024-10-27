#!/bin/bash

echo "pulling grpc-tooling image"
docker pull fuskovic/grpc-tooling:latest
if [ $? -ne 0 ]; then
    echo "failed to pull grpc-tooling image"
    exit 1
fi
echo "successfully pulled grpc-tooling image"

echo "generating protos..."
docker run \
    --name compile-protos \
    -v $(pwd)/protos/:/protos/ \
    --rm \
    fuskovic/grpc-tooling:latest \
    protoc \
    --proto_path=./protos/definitions \
    --go_out=./protos \
    --go-grpc_out=./protos \
    --go-grpc_opt=paths=source_relative \
        cohn.proto \
        live_streaming.proto \
        media.proto \
        network_management.proto \
        preset_status.proto \
        request_get_preset_status.proto \
        response_generic.proto \
        set_camera_control_status.proto \
        turbo_transfer.proto 
if [ $? -ne 0 ]; then
    echo "failed to generate protos"
    exit 1
fi
rm -rf identity
echo "successfully generated protos"