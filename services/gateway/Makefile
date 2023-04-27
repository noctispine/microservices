proto:
	protoc --proto_path ./pkg/baseProto --go_out=. --go_opt=paths=source_relative \
	base.proto \
	&& \
	protoc --go_out=. --go_opt=module=github.com/capstone-project-bunker/backend/services/gateway \
    --go-grpc_out=. --go-grpc_opt=module=github.com/capstone-project-bunker/backend/services/gateway \
    ./pkg/**/pb/*.proto
run:
	APP_ENV=DEVELOPMENT go run cmd/main.go
