run:
	go run cmd/api/main.go

build: swag
	go build -o bin/api cmd/api/main.go

swag:
	swag init -g cmd/api/main.go --output docs

clean:
	rm -rf bin/
	rm -rf docs/

dev:
	air