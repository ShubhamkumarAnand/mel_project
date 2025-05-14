
# Workout Planning

This is a lightweight Go backend application for managing workouts and exercises. It allows users to create and retrieve structured workout routines, including detailed entries for each exercise, such as sets, reps, duration, and weight. The backend is built with Go's standard library and leverages PostgreSQL for data persistence.

### Key features include:

- RESTful API structure for workout creation and retrieval
- PostgreSQL integration using database/sql and pgx
- Embedded SQL migrations via goose
- Modular codebase with clear separation between handlers, routes, and storage
- Transaction-safe creation of workouts and related entries
- Easily extensible architecture for future enhancements like user authentication or progress tracking

Ideal for fitness apps, personal workout logs, or as a base for a full-stack workout tracking solution.

## Setup

The API project is built from scratch. Before watching the course, you should install:
- [Go](https://go.dev/doc/install) (version 1.24.2 or higher)
- [Postgres](https://www.postgresql.org/download/) and any DB tool like psql or Sequel Ace to run SQL queries.
- [Docker and Docker Compose](https://www.docker.com/)

## Setup Tips
- In the [Postgres Database Container lesson][database], the Docker container exposes Postgres on the default port of `5432`. If you already have Postgres or something else running on that port and you get a connection error, you can use an alternate port but updating the `docker-compose.yml` to be something like `"5433:5432"`.
- In the [SQL Migrations with Goose lesson][goose], if you get a "command not found" error when running `goose -version`, it's because the `$HOME/go/bin` directory is not added to your `PATH`. You can fix this temporarily by running `export PATH=$HOME/go/bin:$PATH`, but this will not persist if you close your terminal. A permanent fix would require adding `export PATH=$HOME/go/bin:$PATH` to your `.zshrc` or `.bashrc`.
