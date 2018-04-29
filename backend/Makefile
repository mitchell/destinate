all:
	make migrate
	make destinations
	make users

migrate:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate functions/migrate/main.go

destinations:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/destinations-index functions/destinations/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/destinations-get functions/destinations/get/main.go

users:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-index functions/users/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-post functions/users/post/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-get functions/users/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-destinations-put-like functions/users/destinations/like/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-destinations-put-dislike functions/users/destinations/dislike/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-put functions/users/put/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-delete functions/users/delete/main.go

clean:
	rm -r ./bin ./vendor
