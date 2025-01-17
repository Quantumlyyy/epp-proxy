module github.com/as207960/epp-proxy/v2

go 1.15

require (
	github.com/golang/glog v1.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	google.golang.org/grpc v1.45.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.0
	github.com/as207960/epp-proxy/gen/go/epp v0.0.0
)

replace github.com/as207960/epp-proxy/gen/go/epp => ./gen/go