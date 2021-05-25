NAME=spamhouse
IMAGE=$(NAME)
TEST_IMAGE=$(NAME)-test
DOCKERFILE=ci/Dockerfile
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

.PHONY: regenerate
regenerate:
	@echo "--> regenerating graphql resolvers"
	@go generate ./...
	
.PHONY: test
test:
	@echo "--> running go tests for $(NAME)"
	@go test -v ./...

.PHONY: test-integration
test-integration: stop
	@echo "--> starting integration tests for $(TEST_IMAGE)"
	@PORT=$(PORT) docker-compose up -d spamhouse 
	@docker build --target build1 -t $(TEST_IMAGE) .
	@PORT=$(PORT) docker-compose up --remove-orphans test


.PHONY: db-shell
db-shell:
	docker exec -it spamhouse sqlite3 spamhouse.db


