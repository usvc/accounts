---
title: Configuration Documentation
authors:
- zephinzer
tags:
- environment
- configuration
- database
- server
- application
created_at: 2018-12-22
updated_at: 2018-12-22
tldr: Documentation for environment variables that can be used to configure this
path: docs/Config
---
# Configuration
Configurations are initialised in the `config.go` file

## `ENVIRONMENT`
- Defines the environment we will be running in
- Defaults to `"development"`
- Other example values: `"production"`

## `INTERFACE`
- Defines the interface which the application will bind to for listening
- Defaults to `"0.0.0.0"`
- Other example values: `"14.55.123.44"`

## `PORT`
- Defines the port which the application server will listen on
- Defaults to `3000`
- Other example values: `1000` - `65535`

## `LOG_FORMAT`
- Defines the format of the logs
- Defaults to `"text"`
- Other example values: `"json"`

## `LOG_LEVEL`
- Defines the minimum level at which logs are streamed to the output
- Defaults to `"trace"`
- Other example values: `"debug"`, `"info"`, `"warn"`, `"error"`

## `LOG_SOURCE_MAP`
- Defines whether the caller function should be recorded in the log
- Defaults to `true` (production recommendation: `false`)
- Other example values: `false`

## `LOG_PRETTY_PRINT`
- Defaults to `true` (production recommendation: `false`)
- Other example values: `false`

## `DB_HOST`
- Defines the host of the database instance for persistent data
- Defaults to `"database"`
- Other example values: any valid hostname

## `DB_PORT`
- Defines the port of the database instance for persistent data
- Defaults to `3306`
- Other example values: `1000` - `65535` depending on your database configuration

## `DB_DATABASE`
- Defines the schema to use
- Defaults to `"database"`
- Other example values: any valid schema name

## `DB_USER`
- Defines the user to use to login to the database instance
- Defaults to `"user"`
- Other example values: any valid username

## `DB_PASSWORD`
- Defines the password that matches the user defined in `DB_USER`
- Defaults to `"password"`
- Other example values: any valid password
