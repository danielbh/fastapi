#GET_ENV = $(export $(cat .env | xargs))

build:
	@go build

run:
	@export $(cat .env | xargs) && go run main.go

   