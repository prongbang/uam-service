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

gen_role:
	make gen name=role in=internal/shared/role out=internal/shared

gen_auth:
	make gen name=auth in=internal/shared/auth out=internal/shared

run:
	go run cmd/user/main.go

# brew install grpcurl
# Postman: https://learning.postman.com/docs/sending-requests/grpc/grpc-request-interface/#creating-a-new-request
username_login:
	grpcurl -plaintext -import-path ./internal/service/uam -proto uam.proto \
	-d '{"username": "superadmin", "password": "super.admin"}' \
	'[::1]:50052' uam.Uam/Login

email_login:
	grpcurl -plaintext -import-path ./internal/service/uam -proto uam.proto \
	-d '{"email": "super.admin@gmail.com", "password": "super.admin"}' \
	'[::1]:50052' uam.Uam/Login