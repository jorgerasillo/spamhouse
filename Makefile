NAME=spamhouse
IMAGE=$(NAME)
DOCKERFILE=ci/Dockerfile
TEST_DOCKERFILE=test/Dockerfile
# This value matches the docker volume created in via docker-compose
VOL_NAME=data
PORT=8080

.PHONY: build
build:
	@echo "--> building image: $(NAME)"
	@docker build -t $(IMAGE) .

.PHONY: dev
dev: stop build run

.PHONY: run
run: stop
	@echo "--> starting $(NAME) on port $(PORT)"
	@PORT=$(PORT) docker-compose up spamhouse

.PHONY: stop
stop:
	@echo "--> stopping $(NAME)"
	@PORT=$(PORT) docker-compose down

.PHONY: test
test:
	@echo "--> running go tests for $(NAME)"
	@go test -v ./...

.PHONY: test-integration
test-integration:
	@echo "--> starting integration tests for $(NAME)"
	@PORT=$(PORT) docker-compose up test

.PHONY: delete-db
delete-integration-db:
	@echo "--> Deleting docker volume with db data $(VOL_NAME)"
	@docker kill $(NAME) || true
	@docker rm $(NAME) || true
	@docker volume rm $(VOL_NAME)

.PHONY: delete-local-db
delete-local-db:
	@echo "--> Deleting local db data $(NAME).db"
	@rm ./$(NAME).db || true

.PHONY: db-shell
db-shell:
	docker exec -it spamhouse sqlite3 spamhouse.db


