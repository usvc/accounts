# Security

## Updating a User's Password

### Example
Before running the example, run:

```bash
curl -X POST -d '{"email":"email@domain.com","password":"password1!"}' 'http://localhost:54321/user'
```

Get the user's UUID from the above and input it below as `$USER_UUID`

```bash
curl -X PATCH -d '{"password":"password!2"}' 'http://localhost:54321/security/$USER_UUID'
```

#### Success Case
```json
{
  "error_code": "E_SECURITY_PASSWORD_CHANGE_OK",
  "message": "ok",
  "data": null,
  "timestamp": "Fri, 14 Dec 2018 17:07:59 UTC"
}
```

#### Fail Case #1: Password too short
```json
{
  "error_code": "E_PASSWORD_TOO_SHORT",
  "message": "password should be at least of length 8",
  "data": {},
  "timestamp": "Fri, 14 Dec 2018 17:10:49 UTC"
}
```

#### Fail Case #2: Password no numerics
```json
{
  "error_code": "E_PASSWORD_NO_NUMBERS",
  "message": "password should contain at least one numerical character",
  "data": {},
  "timestamp": "Fri, 14 Dec 2018 17:11:30 UTC"
}
```

#### Fail Case #3: Password no special characters
```json
{
  "error_code": "E_PASSWORD_NO_SPECIAL_CHARACTERS",
  "message": "password should contain at least one special character",
  "data": {},
  "timestamp": "Fri, 14 Dec 2018 17:11:10 UTC"
}
```
