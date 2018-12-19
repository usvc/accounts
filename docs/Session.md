# Sessions

## Creating New Sessions
See [the User module documentation](./User.md) for information on how to create a new User. You'll need the User's UUID for the calls below.

### Example
```sh
curl -d '{"account_uuid":"e40b05ba-03a7-11e9-aaac-0242ac150002","ipv4":"127.0.0.1","ipv6":"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff","device":"mac","source":"website","token_refresh":"absdbasbdasbdbasd","token_access":"aoskdoaksdoksad_access","date_expires":"2018-12-24T10:10:00"}' -X POST 'http://localhost:54321/session'
```

#### Success Case
```json
{
  "error_code": "E_SESSIONS_API_CREATE_OK",
  "message": "ok",
  "data": null,
  "timestamp": "Wed, 19 Dec 2018 16:18:38 UTC"
}
```

#### Error Case: User UUID not found
```json
{
  "error_code": "E_USER_NOT_FOUND",
  "message": "the user identified by e40b05ba-03a7-11e9-aaac-0242ac150002 does not exist",
  "data": null,
  "timestamp": "Wed, 19 Dec 2018 16:19:15 UTC"
}
```