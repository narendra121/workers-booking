get-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    # go get -u google.golang.org/grpc
#  make generate-grpc-models file_path=proto/models/*.proto
generate-grpc-models:
	protoc -I=. --go_out=. proto/models/*.proto

#  make generate-grpc-svc file_path=proto/svc/*.proto
generate-grpc-svc:
	protoc -I=. --go_out=. proto/svc/*.proto