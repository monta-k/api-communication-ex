run-go-server:
	docker compose -f docker-compose.go-server.yml up

generate-oapi-codegen:
	docker compose -f docker-compose.oapi-codegen.yml up