include .env

build:
	go build -o ${BINARY} ./cmd/api/main.go

start:
	@env MONGO_DB_USERNAME=${MONGO_DB_USERNAME} MONGO_DB_PASSWORD=${MONGO_DB_PASSWORD} ./${BINARY}

restart: build start