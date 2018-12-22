---
title: "ADR: Decoupling of Session and Authentication logical entities"
authors:
- zephinzer
tags:
- session
- authentication
- api
- controller
created_at: 2018-12-20
updated_at: 2018-12-20
tldr: We shall separate authentication mechanisms from session creation/maintenance
path: docs/adr/20181220-decoupling-session-authentication
---

# Title
ADR: Decoupling of Session and Authentication logical entities

# Status
Approved

# Context
When it comes to sessions and authentications, they seem coupled, but if implemented as such in code, an endpoint at `/session` should be able to receive a `POST` verb and be expected to autheticate the user.

A simple RESTful API might implement authentication as `POST /session` for logging in, and `DELETE /session` for logging out, but we keep in mind that we need to remain open to multiple types of authentications. In the near future, email-only/Faceobok/Google-based authentication is due, and tying the authentication entity to the session CRUD operations will result in having to extract it later.

# Decision
Since we are certain of the direction where multiple modes of authentication are provisioned for, we go with using a separate logical entity, Authentication, to handle authentication related flows.

# Consequences
This logic might seem conventional to new developers, hence this ADR.

No other long-term consequences are foreseen.
