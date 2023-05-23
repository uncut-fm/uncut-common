user-proto:
	cd pkg/proto/auth/user && protoc --go_out=. --go_opt=paths=source_relative \
             --go-grpc_out=. --go-grpc_opt=paths=source_relative \
             user.proto

graph-proto:
	cd pkg/proto && protoc --go_out=. --go_opt=paths=source_relative \
             --go-grpc_out=. --go-grpc_opt=paths=source_relative \
             graph/network.proto