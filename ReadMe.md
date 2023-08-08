# Golang - GIN Boilerplate

Boilerplate ready to extend and use

## Env File
Create a env file, in env_back is the template where you have to replace the values correctly, it is needed to run correctly the service

## Migrations

Create new migration file: 
> ` migrate create -ext sql -dir migrations/ -seq <<migration_sequence_name>>`

Run Migrations
> ` migrate -path migrations/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up `

Rollback Migrations

> ` migrate -path migrations/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down ` 

Force Migration

> ` migrate -path migrations/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force <VERSION> `

## To set an API Secret using ssl

> `openssl rand -hex 32`

## Create mocks

To create mocks of the interfaces go to the folder and run
> ` mockery --all `

## Run tests
> `go test ./...`

To get test coverage, run
> `go test ./... -coverprofile coverage.out`

To see the coverage output in a web browser
> `go tool cover -html coverage.out`

## Documentation

After running the service the documentation can be found in 

> `/docs/index.html`

After creating and endpoint run `swag init` to create the documentation files, and `swag fmt` to run formatting in the file