# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

# Migration parameters
DATABASE_URL="postgres://postgres:@localhost:5432/odisha_dev?sslmode=disable"
MIGRATION_LOCATION=./db/migrations/
GO_MIGRATE=migrate -verbose
MIGRATE=${GO_MIGRATE} -path ${MIGRATION_LOCATION} -database ${DATABASE_URL}

# SqlBoiler
SQLBOILER=sqlboiler
DB_CONFIG_FILE=./db/sqlboiler.toml
DB_CONFIG_TO_USE=psql


all: debug

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run: build
	./$(BINARY_NAME)

debug:
	$(GOCMD) run ./main.go

deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

# Migration
create_migration:
	${GO_MIGRATE} create -ext sql -dir ${MIGRATION_LOCATION} -seq ${name}

migrate_up:
	${MIGRATE} up

migrate_up_to:
	${MIGRATE} up ${version}

migrate_down:
	${MIGRATE} down

migrate_down_to:
	${MIGRATE} down ${version}

migrate_version:
	${MIGRATE} version

migrate_drop:
	${MIGRATE} drop

migrate_force:
	${MIGRATE} force ${version}

models_update:
	${SQLBOILER} ${DB_CONFIG_TO_USE} -c ${DB_CONFIG_FILE}


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v