build:
	@go build -o bin/api
run:build
	@./bin/api
test:
	@go test -v --cover ./...
gen:
	@sqlc generate