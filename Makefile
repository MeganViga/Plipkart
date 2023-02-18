createdb:
	docker exec -it postgres createdb --username=root --owner=root plipkart
dropdb:
	docker exec -it postgres dropdb plipkart
migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateup:
		migrate -path db/migration -database "postgres://root:secret@localhost:5432/plipkart?sslmode=disable" -verbose up
migratedown:
		migrate -path db/migration -database "postgres://root:secret@localhost:5432/plipkart?sslmode=disable" -verbose down
.PHONY: migratecreate createdb dropdb migrateup migratedown