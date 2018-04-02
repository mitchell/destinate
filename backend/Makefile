all:
	make migrate
	make destinations
	make reviews
	make users

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

reviews:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/reviews-index functions/reviews/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/reviews-post functions/reviews/post/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/reviews-get functions/reviews/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/reviews-put functions/reviews/put/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/reviews-delete functions/reviews/delete/main.go

users:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-index functions/users/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-post functions/users/post/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-get functions/users/get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-put functions/users/put/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users-delete functions/users/delete/main.go

clean:
	rm -r ./bin
