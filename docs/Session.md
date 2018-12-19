# Sessions

## Creating New Sessions
> See [the User module documentation](./User.md) for information on how to create a new User. You'll need the User's UUID for the calls below.

| Request | Value |
| --- | --- |
| Method | `POST` |
| URL | `/session` |

| Body Parameter | Description |
| --- | --- |
| account_uuid | UUID of the user (**required**) |
| ipv4 | IP address (version 4) where the session was started |
| ipv6 | IP address (version 6) where the session was started |
| source | A source for where the user session is being created for |
| device | A device code for where the device the user is using |
| token_refresh | The refresh token for the user |
| token_access | The access token for the user |
| date_expires | A timestamp for the expiry of this session |

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