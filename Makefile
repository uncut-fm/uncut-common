user-proto:
	cd pkg/proto/auth/user && protoc --go_out=. --go_opt=paths=source_relative \
             --go-drpc_out=. --go-drpc_opt=paths=source_relative \
             user.proto