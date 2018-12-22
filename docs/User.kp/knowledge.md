---
title: User Documentation
authors:
- zephinzer
tags:
- account
- user
- registration
- sign up
created_at: 2018-12-22 00:00:00
updated_at: 2018-12-22 12:41:37.992101
tldr: Documentation for the `/user` endpoint and related entities
path: docs/User
---

# User Accounts

## Creating a User

| Request | Value |
| --- | --- |
| Method | `POST` |
| URL | `/user` |

| Body Parameter | Description |
| --- | --- |
| email | Email of the user |
| password | Password of the user |

### Example
```bash
curl -X POST -d '{"email":"email@domain.com","password":"password1!"}' 'http://localhost:54321/user'
```

#### Success Case
```json
{
  "error_code": "E_USER_API_CREATE_OK",
  "message": "ok",
  "data": {
    "uuid": "807dce47-fe11-11e8-a683-0242ac120002",
    "email": "email@domain.com",
    "username": "",
    "date_created": "2018-12-12 13:26:15",
    "last_modified": "2018-12-12 13:26:15"
  },
  "timestamp": "Wed, 12 Dec 2018 13:26:15 UTC"
}
```

## Retrieving a list of Users

| Request | Value |
| --- | --- |
| Method | `GET` |
| URL | `/user` |

| URL Query Parameter | Description |
| --- | --- |
| start | 0 based start index to start retrieving users from |
| limit | Number of users to return |

### Example #1: Mindless retrieve
```bash
curl -X GET 'http://localhost:54321/user'
```

#### Success Case
```json
{
  "error_code": "E_USER_API_QUERY_OK",
  "message": "ok",
  "data": [
    {
      "uuid": "f5889750-fdaf-11e8-a683-0242ac120002",
      "email": "emailOne@domain.com",
      "username": "",
      "date_created": "2018-12-12 01:48:01",
      "last_modified": "2018-12-12 01:48:01"
    },
    {
      "uuid": "807dce47-fe11-11e8-a683-0242ac120002",
      "email": "email@domain.com",
      "username": "",
      "date_created": "2018-12-12 13:26:15",
      "last_modified": "2018-12-12 13:26:15"
    }
  ],
  "timestamp": "Wed, 12 Dec 2018 13:28:03 UTC"
}
```

### Example #2: Sensible retrieve
```bash
curl -X GET 'http://localhost:54321/user?start=0&limit=1'
```

#### Success Case
```json
{
  "error_code": "E_USER_API_QUERY_OK",
  "message": "ok",
  "data": [
    {
      "uuid": "f5889750-fdaf-11e8-a683-0242ac120002",
      "email": "emailOne@domain.com",
      "username": "",
      "date_created": "2018-12-12 01:48:01",
      "last_modified": "2018-12-12 01:48:01"
    }
  ],
  "timestamp": "Wed, 12 Dec 2018 13:28:28 UTC"
}
```

## Retrieving a User

| Request | Value |
| --- | --- |
| Method | `GET` |
| URL | `/user/:uuid` |

| Path Parameter | Description |
| --- | --- |
| uuid | UUID of the user |

### Example
```bash
curl -X GET 'http://localhost:54321/user/$USER_UUID'
```

#### Success Case
```json
{
  "error_code": "E_USER_API_GET_OK",
  "message": "ok",
  "data": {
    "uuid": "f5889750-fdaf-11e8-a683-0242ac120002",
    "email": "emailOne@domain.com",
    "username": "",
    "date_created": "2018-12-12 01:48:01",
    "last_modified": "2018-12-12 01:48:01"
  },
  "timestamp": "Wed, 12 Dec 2018 13:30:19 UTC"
}
```

## Updating a User

| Request | Value |
| --- | --- |
| Method | `PATCH` |
| URL | `/user/:uuid` |

| Path Parameter | Description |
| --- | --- |
| uuid | UUID of the user |

| Body Parameter | Description |
| --- | --- |
| email | Email address of the user |
| username | Username of the user |

### Example #1: Single field update
```bash
curl -X PATCH -d '{"email":"modified@cooldomain.com"}' 'http://localhost:54321/user/$USER_UUID'
```

#### Success Case
```json
{
  "error_code": "E_USER_API_UPDATE_OK",
  "message": "",
  "data": {
    "uuid": "f5889750-fdaf-11e8-a683-0242ac120002",
    "email": "modified@cooldomain.com",
    "username": "",
    "date_created": "",
    "last_modified": ""
  },
  "timestamp": "Wed, 12 Dec 2018 14:02:34 UTC"
}
```

### Example #2: Batch field updates
```bash
curl -X PATCH -d '{"email":"modified@otherdomain.com","password":"password2!"}' 'http://localhost:54321/user/$USER_UUID'
```

#### Success Case
```json
{
  "error_code": "E_USER_API_UPDATE_OK",
  "message": "",
  "data": {
    "uuid": "f5889750-fdaf-11e8-a683-0242ac120002",
    "email": "modified@otherdomain.com",
    "username": "username",
    "date_created": "",
    "last_modified": ""
  },
  "timestamp": "Wed, 12 Dec 2018 14:01:41 UTC"
}
```

## Deleting a User

### Example
```bash
curl -X DELETE 'http://localhost:54321/user/$USER_UUID'
```

#### Success Case
```json
{
  "error_code": "E_USER_API_DELETE_OK",
  "message": "ok",
  "data": {
    "uuid": "f5889750-fdaf-11e8-a683-0242ac120002"
  },
  "timestamp": "Wed, 12 Dec 2018 14:11:14 UTC"
}
```

#### Failed Case
```json
{
  "error_code": "E_USER_NOT_FOUND",
  "message": "user with uuid 'f5889750-fdaf-11e8-a683-0242ac120002' could not be found",
  "data": null,
  "timestamp": "Wed, 12 Dec 2018 14:11:26 UTC"
}
```

- - -

[Back to Main README.md](../README.md)