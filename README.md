# usvc/accounts
An accounts microservice for easy addition of a login requirement to your system.

# Current Scope of Component (For Version 1.0)

- [x] User account creation with email/password
- [x] Updating of user's account information
- [x] Deletion of an existing user
- [x] Retrieval of a user's account information
- [x] Retrieval of a list of users
- [x] Session creation
- [x] Authentication with email/username and password
- [ ] Session status retrieval (still logged in?)
- [ ] Session deletion

See [the roadmap section](#roadmap) for future developements.

# Endpoints

| Method | Endpoint | Description | Version | Reference |
| --- | --- | --- | --- | --- |
| GET | /user | Retrieves users | 0.0.2 | [#6](//github.com/usvc/accounts/issues/6) |
| POST | /user | Register a new user | 0.0.1 | [#1](//github.com/usvc/accounts/issues/1) |
| GET | /user/:uuid | Returns the user identified by the :uuid | 0.0.1 | [#2](//github.com/usvc/accounts/issues/2) |
| PATCH | /user/:uuid | Updates a field of the user identified by :uuid | 0.0.3 | [#4](//github.com/usvc/accounts/issues/4) |
| DELETE | /user/:uuid | Deletes the user identified by the :uuid | 0.0.1 | [#3](//github.com/usvc/accounts/issues/3) |
| PATCH | /security/:user_uuid | Updates the password for the user identified by the :user_uuid | 0.0.4 | [#7](//github.com/usvc/accounts/issues/7) |
| POST | /auth/credentials | Verifies a user's email/username and password matches | 0.0.7 | [#24](//github.com/usvc/accounts/issues/24) |
| POST | /session | Creates the user session | 0.0.6 | [#12](//github.com/usvc/accounts/issues/12) |
| DELETE | /session | Destroys the user session | **TODO** | [#13](//github.com/usvc/accounts/issues/13) |
| GET | /metrics | Returns Prometheus metrics | **TODO** | N/A |
| GET | /healthz | Returns 200 OK if healthy | **TODO** | [#8](//github.com/usvc/accounts/issues/8) |
| GET | /readyz | Returns 200 OK if ready to accept connections | **TODO** | [#9](//github.com/usvc/accounts/issues/9) |

> Refer to [the User documentation](./docs/User.md) for details on the user endpoints.  
>
> Refer to [the Security documentation](./docs/Security.md) for detials on the security endpoints.  
>
> Refer to [the Session documentation](./docs/Session.md) for details on the session endpoints.
>
> Refer to [the Auth documentation](./docs/Auth.md) for details on the authentication endpoints.

# Configuration

| Key | Default | Description |
| --- | --- | --- |
| ENVIRONMENT | `"development"` | Defines the environment we will be running in |
| INTERFACE | `"0.0.0.0"` | Defines the interface the server should bind to |
| PORT | `3000` | Defines the port the server should listen on |
| LOG_FORMAT | `"text"` | Either 'text' or 'json' |
| LOG_LEVEL | `"trace"` | Defines minimum level of the logs to print - choose from 'trace', 'debug', 'info', 'warn', 'error', 'panic', 'fatal' |
| LOG_SOURCE_MAP | `true` | Defines whether to print the source of the log call |
| LOG_PRETTY_PRINT | `true` | Defines whether to output the logs in a human-optimised manner |
| DB_HOST | `"database"` | Defines the hostname of the database instance |
| DB_PORT | `3306` | Defines the port which the database instance is listening on |
| DB_DATABASE | `"database"` | Defines the schema name to use upon a successful connection |
| DB_USER | `"user"` | Defines the user of the database instance |
| DB_PASSWORD | `"password"` | Defines the password of the user defined in `DB_USER` |

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
1. [Create a new issue](https://github.com/usvc/accounts/issues/new) (if one isn't already available)
1. Create a new branch with a descriptive title like `feature-<ISSUE_NUMBER>/describe-it-briefly` and create a pull request (PR). Prefix the PR with `WIP` while it's still in progress
1. Make changes to the codebase on your branch
1. Finalise your PR, remove the `WIP` and prefix your PR title with a `[#?]` where `?` is the issue number
1. Squash and merge when ready and close the issue

## Others
1. Fork this repository
2. [Create a new issue on this repository](https://github.com/usvc/accounts/issues/new)
3. Make changes to the `master` branch of your fork
4. Push your changes to your fork
5. Raise a pull request to this repository and prefix the title with a `[#?]` where `?` is the issue number
6. Ping a contributor if you're getting impatient

# Roadmap

## Version 1.0
- [x] Registration of new users using an email address and password
- [x] Removal of an existing user from the system
- [x] Retrieval of a list of users
- [x] Updating of existing users' account information
- [x] Updating of existing users' password
- [x] Logging in of existing users via password
- [ ] Logging out of existing users
- [ ] Session maintenance of existing users

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

# Changelog

### 0.0.7
- [x] Added endpoint for verification of credentials at `POST /auth/credentials`. See [the Auth documentation](./docs/Auth.md) for more information.

### 0.0.6
- [x] Added session creation at `POST /session`. See [the Session documentation](./docs/Session.md) for more information.

### 0.0.5
- [x] Enable configuring of the logging format via `LOG_FORMAT`, `LOG_LEVEL`, `LOG_SOURCE_MAP` and `LOG_PRETTY_PRINT` environment variables.

### 0.0.4
- [x] Added endpoint for updating of existing users' password. See [the Security documentation](./docs/Security.md) for more information.

### 0.0.3
- [x] Updating of existing users' account information. See [the User documentation](./docs/User.md) for more information.

### 0.0.2
- [x] Retrieval of a list of users. See [the User documentation](./docs/User.md) for more information.

### 0.0.1
- [x] Registration of new users using an email address and password.
- [x] Removal of an existing user from the system.

# License
This project is licensed under [the MIT license](./LICENSE). Feel free to fork it and create your own!
