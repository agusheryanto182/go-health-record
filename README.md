## What Is Health Record?
This is a backend that nurses can use to record patient medical records

## Key Technologies

1. **Go**: Go (or Golang) is the programming language that can develop an API with high performance and scalable.

2. **Fiber**: Fiber is a high-performance and minimalist web framework for the Go programming language.

3. **PostgreSQL**: PostgreSQL is a powerful, open-source relational database management system. It is used for storing and managing data related to products, customers, orders, and more in the online store application.

4. **JWT (JSON Web Tokens)**: JWT is a standard for securely transmitting information between parties as a JSON object. In the online store application, JWT is used for implementing authentication and authorization mechanisms.

5. **Pgx**: Pgx is a pure Go driver and toolkit for PostgreSQL. The pgx driver is a low-level, high performance interface that exposes PostgreSQL-specific features such as LISTEN / NOTIFY and COPY. It also includes an adapter for the standard database/sql interface.

6. **Sqlx**: Sqlx is a library for Go that provides extensions to the standard database/sql package. It aims to simplify working with SQL databases in Go by offering additional features and utilities. Sqlx includes functions for working with query results, struct scanning, named parameters, and more. It enhances the productivity of Go developers when interacting with SQL databases like PostgreSQL by reducing boilerplate code and providing a more intuitive API.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/agusheryanto182/go-health-record.git
   ```

2. Go to folder go-raide-hailings

   ```bash
   cd go-health-record
   ```

3. Set up an env and paste in terminal

4. Create a new database

5. Copy this and paste it in terminal for migrate database

   ```bash
   make migrate-dev
   ```

6. ```bash
   go run .
   ```

## Documentation

1. Postman

   ```bash
   https://documenter.getpostman.com](https://documenter.getpostman.com/view/32137512/2sA3Qqfsg7
   ```

2. Database

   ```bash
   https://dbdiagram.io/d/health-record-6641bfc79e85a46d55ac25a0
   ```
