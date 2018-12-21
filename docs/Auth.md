# Authentication Mechanisms

## Authenticating via Credentials
This method covers authenticating a user based on their email/username, and their password.

### Example
This example requires that [a User be created](./User.md).

```sh
curl -d '{"password":"password1!","email":"email@domain.com"}' -X POST 'http://localhost:54321/auth/credentials'
```

> The above will provide a success case with the user creation process outlined in the [User documentation](./User.md), to generate the error cases, change or omit the password/email

#### Success Case
```json
{
  "error_code": "E_AUTH_API_OK",
  "message": "ok",
  "data": null,
  "timestamp": "Fri, 21 Dec 2018 16:30:46 UTC"
}
```

#### Error Case 1: Missing Parameters
```json
{
  "error_code": "E_AUTH_CREDENTIALS_MISSING_PARAMS",
  "message": "either 'username' or 'email', and 'password', should be specified",
  "data": null,
  "timestamp": "Fri, 21 Dec 2018 16:34:30 UTC"
}
```

#### Error Case 2: Everything Else
For all other errors such as no user found or wrong password, we return a generic error to avoid exposing any meta-information which could be used for information gathering by a malicious actor.

```json
{
  "error_code": "E_AUTH_CREDENTIALS_INVALID_PARAMS",
  "message": "the email/username/password combination does not exist",
  "data": null,
  "timestamp": "Fri, 21 Dec 2018 16:36:52 UTC"
}
```
