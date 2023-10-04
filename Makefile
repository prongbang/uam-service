install:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	brew install grpcurl

# make gen name=uam in=internal/shared/uam out=internal
gen:
	protoc --go_out=$(out) --go_opt=paths=import \
        --go-grpc_out=$(out) --go-grpc_opt=paths=import \
        $(in)/$(name).proto

gen_user:
	make gen name=user in=internal/shared/user out=internal/shared

run:
	go run cmd/user/main.go

# brew install grpcurl
test_login:
	grpcurl -plaintext -import-path ./internal/service/user -proto user.proto -d '{"username": "admin"}' '[::1]:50052' user.User/GetUser