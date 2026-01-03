
.Phony:
docker-build-database:
	docker build -f ./dockerfiles/database.Dockerfile . -t gcr.io/small-biz-template/markenshop/lottery-db:latest

.Phony:
docker-run-database:
	docker run --rm -d -p 5432:5432 -v mtg_card-db:/var/lib/postgresql/data -t gcr.io/small-biz-template/markenshop/calendar-db:latest

.Phony:
migrate-up:
	goose -dir ./database/postgres/migrations postgres "postgresql://user:password@localhost:5432/postgres?sslmode=require" up

.Phony:
migrate-down:
	goose -dir ./database/postgres/migrations postgres "postgresql://user:password@localhost:5432?sslmode=require" down

.Phony:
migration-status:
	goose -dir ./database/postgres/migrations postgres "postgresql://user:password@localhost:5432?sslmode=require" status
