## Run

run:
	go run app/to-do/main.go

tidy:
	go mod tidy

to-do:
	docker build \
	-f setup/docker/dockerfile.to-do \
	-t to-do-service:1.0 \
	.
	