test:
	go test -coverprofile=coverage.out -race -covermode=atomic ./...
	go tool cover -html=coverage.out
	cd web && yarn test && cd ../

test-web:
	cd web && yarn test && cd ../

test-go:
	go test -v ./...

run-db:
	docker-compose -f docker-compose-db.yml up

run-data-collector:
	go run cmd/data-collector/main.go

run-api:
	go run cmd/api/main.go

run-web:
	cd web && yarn dev
