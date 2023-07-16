# Halyk Todo List Api

Halyk Todo List Api Provides Simple CRUD Service for Tasks

## Configuration

Swagger Documentation at : http://localhost:8000/internal/docs/index.html

Configuration is fully loaded from environment. Table below shows the list of environment variables that could be used.

| Name | Description | Example |
|------|-------------|---------|
|API_NAME|Name of the API|halyk-todo-list|
|API_HOST|API Host|localhost|
|API_PORT|API Port|8000|
|DATABASE_TYPE|Type of database to use|postgres|
|POSTGRES_DATABASE_HOST|Postgres host|localhost|
|POSTGRES_DATABASE_PORT|Postgres port|5432|
|POSTGRES_DATABASE_USERNAME|Postgres username|postgres|
|POSTGRES_DATABASE_PASSWORD|Postgres password|postgres|
|POSTGRES_DATABASE_DBNAME|Postgres database|halyk-todolist|

## TODO

Cache for service (At this moment not fully understand how cache will be work with high dynamic updating service)

Logger (In this version loggin in data layer. With adding all service logging, it will be more informational)