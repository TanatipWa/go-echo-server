protoc ^
-I. ^
-I./third_party/googleapis/ ^
--go-grpc_out=:pkg/proto/ ^
--go_out=:pkg/proto/ ^
api/server/*.proto

go build -o ../../bin/server