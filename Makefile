pb:
	@protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

run-server:
	@go run server/grpc-server.go

run-client:
	@go run client/grpc-client.go

install-protoc:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest