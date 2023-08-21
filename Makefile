build:
	swag init --parseDependency --parseInternal --parseDepth 3 -g cmd/main.go --output assets/docs
	go build cmd/main.go

run:
	swag init --parseDependency --parseInternal --parseDepth 3 -g cmd/main.go --output assets/docs
	go run cmd/main.go