# usvc/accounts
An accounts microservice for easy addition of a login requirement to your system

# Scope of Component

## Version 1.0
### 0.0.1
- [x] Registration of new users using an email address and password
- [x] Removal of an existing user from the system

### 0.0.2
- [x] Retrieval of a list of users
- [ ] Updating of existing users' account information

### 0.1.0
- [ ] Logging in of existing users via password
- [ ] Logging out of existing users

### 0.2.0
- [ ] Session maintenance of existing users

See [the roadmap section](#roadmap) for future developements.

# Endpoints

| Method | Endpoint | Description | Version | Reference |
| --- | --- | --- | --- | --- |
| GET | /user | Retrieves users | 0.0.2 | [#6](//github.com/usvc/accounts/issues/6) |
| POST | /user | Register a new user | 0.0.1 | [#1](//github.com/usvc/accounts/issues/1) |
| GET | /user/:uuid | Returns the user identified by the :uuid | 0.0.1 | [#2](//github.com/usvc/accounts/issues/2) |
| PATCH | /user/:uuid | Updates a field of the user identified by :uuid | **TODO** | [#4](//github.com/usvc/accounts/issues/4) |
| PUT | /user/:uuid | Replaces the data of the user identified by :uuid | **TODO** | [#5](//github.com/usvc/accounts/issues/5) |
| DELETE | /user/:uuid | Deletes the user identified by the :uuid | 0.0.1 | [#3](//github.com/usvc/accounts/issues/3) |
| POST | /session | Logs a user in | **TODO** | N/A |
| DELETE | /session | Logs a user out | **TODO** | N/A |
| GET | /metrics | Returns Prometheus metrics | **TODO** | N/A |
| GET | /healthz | Returns 200 OK if healthy | **TODO** | N/A |
| GET | /readyz | Returns 200 OK if ready to accept connections | **TODO** | N/A |

> Refer to [user.api.go](./user.api.go) for details on the user endpoints.  
> Refer to [session.api.go](./session.api.go) for details on the session endpoints.

# Configuration

| Key | Default | Description |
| --- | --- | --- |
| ENVIRONMENT | development | Defines the environment we will be running in |
| INTERFACE | 0.0.0.0 | Defines the interface the server should bind to |
| PORT | 3000 | Defines the port the server should listen on |
| DB_HOST | database | Defines the hostname of the database instance |
| DB_PORT | 3306 | Defines the port which the database instance is listening on |
| DB_DATABASE | database | Defines the schema name to use upon a successful connection |
| DB_USER | user | Defines the user of the database instance |
| DB_PASSWORD | password | Defines the password of the user defined in `DB_USER` |

> Refer to [config.go](./config.go) for details

# Development

You'll need Docker and Docker Compose installed on your machine for the development environment to be provisioned.

Run `make start` to get started in development.

Run `make test` to run automated tests.

> Refer to [the Makefile](./Makefile) for details.

# Deployment

## Support Services

- MySQL database

> Refer to [the Docker Compose file](./docker-compose.yml) for more information on support services required.

## Binary Entrypoints

Run `app` to start the application.

Run `app --migrate` to run the migrations.

> Refer to [main.go](./main.go) for the various usable flags.

# Contribution

## Contributors
1. [Create a new issue](https://github.com/usvc/accounts/issues/new)
2. Make changes to the codebase
3. Prefix your commits with a `[#?]` where `?` is the issue number
4. Push your changes

## Others
1. Fork this repository
2. [Create a new issue on this repository](https://github.com/usvc/accounts/issues/new)
3. Make changes to the `master` branch of your fork
4. Push your changes to your fork
5. Raise a pull request to this repository and prefix the title with a `[#?]` where `?` is the issue number
6. Ping a contributor if you're getting impatient

# Roadmap

## Version 2.0

- [ ] Creation of profiles associated with accounts
- [ ] Registration of new users using Facebook Login
- [ ] Logging in of existing users via Facebook Login
- [ ] Registration of new users using Google login
- [ ] Logging in of existing users via Google Login

## Version 3.0

- [ ] Logging in of existing users via an email message
- [ ] Audit trail of session creation/destruction
- [ ] Audit trail of user's actions

# License
This project is licensed under [the MIT license](./LICENSE).
