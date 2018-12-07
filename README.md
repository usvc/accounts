# usvc/accounts
An accounts microservice.

# endpoints

| Method | Endpoint | Description | Version |
| --- | --- | --- | --- |
| POST | /user | Creates a new user | **TODO** |
| GET | /user/:id | Returns the ID of a user | **TODO** |
| PATCH | /user/:id | Updates a field of the user identifie by :id | **TODO** |
| PUT | /user/:id | Replaces the data of the user identified by :id | **TODO** |
| DELETE | /user/:id | Deletes the user identified by :id | **TODO** |

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
Run `make start` to get started in development.

# license
This project is licensed under [the MIT license](./LICENSE).
