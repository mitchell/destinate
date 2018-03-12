all:
	make migrate
	make destinations

migrate:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate functions/migrate/main.go

destinations:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/destinations-index functions/destinations/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/destinations-post functions/destinations/post/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/destinations-get functions/destinations/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/destinations-put functions/destinations/put/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/destinations-delete functions/destinations/delete/main.go

clean:
	rm -r ./bin
