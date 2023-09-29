# make gen name=user in=internal/user out=internal
gen:
	protoc --go_out=$(out) --go_opt=paths=import \
        --go-grpc_out=$(out) --go-grpc_opt=paths=import \
        $(in)/$(name).proto

gen_user:
	make gen name=user in=internal/user out=internal


run:
	go run cmd/user/main.go

test_login:
	grpcurl -plaintext -import-path ./internal/user -proto user.proto -d '{"username": "admin"}' '[::1]:50052' user.User/GetUser