# Workout Planning

This is a lightweight Go backend application for managing workouts and exercises. It allows users to create and retrieve structured workout routines, including detailed entries for each exercise, such as sets, reps, duration, and weight. The backend is built with Go's standard library and leverages PostgreSQL for data persistence.

## Key features include

- RESTful API structure for workout creation and retrieval
- PostgreSQL integration using database/sql and pgx
- Embedded SQL migrations via goose
- Modular codebase with clear separation between handlers, routes, and storage
- Transaction-safe creation of workouts and related entries
- Easily extensible architecture for future enhancements like user authentication or progress tracking

Ideal for fitness apps, personal workout logs, or as a base for a full-stack workout tracking solution.

## Getting Started

Follow these steps to run the project locally:

1. **Clone the repository:**

   ```sh
   git clone <your-repo-url>
   cd <repo-directory>
   ```

2. **Install dependencies:**
   - Ensure you have Go, Docker, and Docker Compose installed (see [Setup](#setup)).
3. **Start the Postgres database using Docker Compose:**

   ```sh
   docker-compose up -d
   ```

4. **Run database migrations:**

   ```sh
   goose -dir migrations postgres "<your-postgres-connection-string>" up
   ```

   Replace `<your-postgres-connection-string>` with your actual connection string (see your `docker-compose.yml`).
5. **Run the Go server:**

   ```sh
   go run main.go

   # if you have Air installed
   Air
   ```

You can now access the API locally as described in the project documentation.

## Setup

The API project is built from scratch. Before watching the course, you should install:

- [Go](https://go.dev/doc/install) (version 1.24.2 or higher)
- [Postgres](https://www.postgresql.org/download/) and any DB tool like psql or Sequel Ace to run SQL queries.
- [Docker and Docker Compose](https://www.docker.com/)

## Setup Tips

- In the [Postgres Database Container lesson][database], the Docker container exposes Postgres on the default port of `5432`. If you already have Postgres or something else running on that port and you get a connection error, you can use an alternate port but updating the `docker-compose.yml` to be something like `"5433:5432"`.
- In the [SQL Migrations with Goose lesson][goose], if you get a "command not found" error when running `goose -version`, it's because the `$HOME/go/bin` directory is not added to your `PATH`. You can fix this temporarily by running `export PATH=$HOME/go/bin:$PATH`, but this will not persist if you close your terminal. A permanent fix would require adding `export PATH=$HOME/go/bin:$PATH` to your `.zshrc` or `.bashrc`.

[database]: https://www.postgresql.org/
[goose]: https://github.com/pressly/goose

---

## Thanks for Visiting

Thank you for checking out this project! If you find it useful or have suggestions for improvement, feel free to open an issue or submit a pull request. Contributions are always welcome.

If you have any questions or want to connect, you can reach out via the repository's issues page or add your contact information here.

Happy coding and good luck with your fitness app development!
