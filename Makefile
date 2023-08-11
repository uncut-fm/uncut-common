user-proto:
	cd pkg/proto/auth/user && protoc --go_out=. --go_opt=paths=source_relative \
             --go-grpc_out=. --go-grpc_opt=paths=source_relative \
             user.proto