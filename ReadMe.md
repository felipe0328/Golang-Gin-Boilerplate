# Golang - GIN Boilerplate

Boilerplate ready to extend and use

## Migrations

Create new migration file: 
> ` migrate create -ext sql -dir migrations/ -seq <<migration_sequence_name>>`

Run Migrations
> ` migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up `

Rollback Migrations

> ` migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down ` 

Force Migration

> ` migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force <VERSION> `