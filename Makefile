run-go-server:
	docker compose -f docker-compose.go-server.yml up $(BUILD)

generate-go-oapi-codegen:
	docker compose -f docker-compose.go-codegen.yml up $(BUILD) oapi-codegen

generate-go-gqlgen:
	docker compose -f docker-compose.go-codegen.yml up $(BUILD) gqlgen

generate-go-grpc:
	docker compose -f docker-compose.go-codegen.yml up $(BUILD) grpc

generate-go-connect:
	docker compose -f docker-compose.go-codegen.yml up $(BUILD) connect-go