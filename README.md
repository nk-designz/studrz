# RZ_STUD
A neat little school project with the goal of learning database-integration into a server-application.
The name has no meaning whatsoever.

![ui](assets/ui.png)
_Listing of users_

## Building from source
### Prerequisites
* docker
* docker-compose (Kubernetes planned)
* make
### Building
```bash
make docker-build
```
## Starting the application
```bash
make up
```

## Components
### Postgres
Just an basic instance of postgres.
All application state is inside this database.
#### Schema
The schema is auto-migrated by the gorm ORM.
![here](db/schema_erd.svg).
### Server
The REST-Backend of this app.  
#### Configuration:  
Configuration is done via environment variables.
| Variable | Example |
| --- | --- |
| STUDRZ_DBPASSWORD | mysecretpassword |
| STUDRZ_DATABASE | study |
| STUDRZ_DBUSER | rzstud |
| STUDRZ_DBHOST | db |
| STUDRZ_DBPORT | 5432 |
| STUDRZ_DBSSL | false |
| STUDRZ_TIMEZONE | Europe/Berlin |
| STUDRZ_ADMINPASSWORD | start-123 |
| * JAEGER_AGENT_HOST | deploy_jaeger-agent_1 |
| * JAEGER_AGENT_PORT | 6831 |
```*``` Optional for Jaeger tracing
#### API Docs
```yaml
routes:
  "/":
    description: returns a list of all routes
    methods: GET
  "/api":
    description: returns a list of all routes
    methods: GET
  "/api/consume":
    description: returns a list of consumes
    methods: GET
  "/api/consume/:id":
    description: returns a consume by it's id
    methods:
    - GET
    - POST
    - DELETE
  "/api/result":
    description: returns a list of all results
    methods: GET
  "/api/result/:id":
    description: returns a result by it's id
    methods:
    - GET
    - POST
    - DELETE
  "/api/result/:id/consume":
    description: returns a list of a consumes by it's result id
    methods: GET
  "/api/user":
    description: returns a list of all registered users
    methods: GET
  "/api/user/:id":
    description: returns a user by it's id
    methods:
    - GET
    - POST
    - DELETE
  "/api/user/:id/result":
    description: returns a list of a results by it's user id
    methods: GET
```
### UI
A user-friendly React-app.
after start of the docker containers you may go to the [ui](http://localhost:3000/)
Standard User is ```admin``` with the password ```start-123```
```
