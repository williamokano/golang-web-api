# Basic Web API

This is a project for simple experimentation with golang.

The objectives of this project is to make a web api that communicates using json.

The API capabilities will be described below

## How to create a new migration

Just run `make new-migration NAME=migration_name`. It will run the docker `migrate/migrate`. Please check `Makefile`.

## Capabilities

This API is intended for studying purposes only. Here's a list of features that this api intends to have:

- [ ] Simple http endpoint with json support
- [ ] The endpoint should have validation in place
- [ ] Should have openapi documentation
- [ ] Add tests examples
- [ ] Should be easily extensible
- [ ] Must have database connection (postgresql maybe) (try avoid using ORM initially)
- [ ] Authentication (building from zero even though isn't recommended)
- [ ] Authentication must support MFA (Multi-factor authentication)
- [ ] Authentication should support WebAuthn with multiple keys
- [ ] Authorization using either jwt/oauth2 and/or saml (RBAC initially)
- [ ] Implement alternativa authn, like ReBAC (Zanzibar) instead of RBAC
- [ ] Support different communication protocols
  - [ ] GraphQL
  - [ ] gRPC
- [ ] Implement ORM (gorm maybe?)
- [ ] Example of a pub/sub (redis, redic, valkey, sns)
- [ ] Streaming (kafka/rabbit/kinesis, whatever)
- [ ] Some form of queue (sqs probably)
- [ ] TBD
