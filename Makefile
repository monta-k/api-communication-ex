run-go-server:
	docker compose -f docker-compose.go-server.yml up

generate-oapi-codegen:
	docker compose -f docker-compose.go-codegen.yml up oapi-codegen

generate-gqlgen:
	docker compose -f docker-compose.go-codegen.yml up gqlgen