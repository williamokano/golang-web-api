# Basic Web API

This is a project for simple experimentation with golang.

The objectives of this project is to make a web api that communicates using json.

The API capabilities will be described below

## How to create a new migration

Just run `make new-migration NAME=migration_name`. It will run the docker `migrate/migrate`. Please check `Makefile`.

## How to run

### Requirements

A kind of comprehensive list of dependencies needed to run this application

#### Database

You need to have a postgresql running. You can have one for local development using docker and docker compose.

Just type `docker-compose up -d` and the database will be running.

#### Exported variables

At the moment, the only way to configure the database "DSN" is through environment variable. You must provide
the bellow variables.

```bash
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_USER="dev"
export DB_PASSWORD="dev"
export DB_NAME="webapi"
```

You can you your own database, the service itself will run the migrations and create the required schemas as long
it has the database user has the proper permissions.

### Running the service

As easy as running `go run ./cmd/webserver` as it only have a webserver at the moment. More binaries to come... for study 😉

## Capabilities

This API is intended for studying purposes only. Here's a list of features that this api intends to have:

- [x] Simple http endpoint with json support
- [ ] The endpoint should have validation in place
- [x] Should have openapi documentation
- [ ] Add tests examples
- [ ] Should be easily extensible
- [x] Must have database connection (postgresql maybe) (try avoid using ORM initially)
- [x] Database migrations added
- [ ] Authentication (building from zero even though isn't recommended)
- [ ] Authentication must support MFA (Multi-factor authentication)
- [ ] Authentication should support WebAuthn with multiple keys
- [ ] Authorization using either jwt/oauth2 and/or saml (RBAC initially)
- [ ] Implement alternative authn, like ReBAC (Zanzibar) instead of RBAC
- [ ] Support different communication protocols
  - [ ] GraphQL
  - [ ] gRPC
- [ ] Implement ORM (gorm maybe?)
- [ ] Example of a pub/sub (redis, redic, valkey, sns)
- [ ] Streaming (kafka/rabbit/kinesis, whatever)
- [ ] Some form of queue (sqs probably)
- [ ] TBD
