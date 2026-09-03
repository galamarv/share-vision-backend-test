
# Sharing Vision - Article Backend Microservice

A backend microservice built with **Golang**, **Echo**, **GORM**, and **MySQL**, following **Clean Architecture** principles.

---

## Architecture & Features

- **Clean Architecture:** Organized cleanly into Domain, Repository, Usecase, and Delivery (HTTP) layers.
- **Database & Migration:** Automated MySQL database creation and table migration for `posts`.
- **Strict Validation Rules:**
  - **Title:** Required, minimum 20 characters.
  - **Content:** Required, minimum 200 characters.
  - **Category:** Required, minimum 3 characters.
  - **Status:** Required, must be `Publish`, `Draft`, or `Thrash`.
- **Paging Support:** Handled via custom `limit` and `offset` parameters.
- **Logging Middleware:** Built-in request tracing and logging via Echo.

---

## Prerequisites

- **Go:** v1.18 or higher
- **MySQL:** Local database server instance running

---

## Environment Setup

Copy the example environment file to create your active `.env` file:
   ```bash
   cp .env.example .env

```

---

## Running the Application

1. Install dependencies:
```bash
go mod tidy

```


2. Run the server:
```bash
go run main.go

```

The server will automatically check and create the `article` database if it doesn't exist, execute GORM auto-migrations to generate the `posts` table, and launch on port `8080`.

---

## API Endpoints Summary

| Method | Endpoint | Description |
| --- | --- | --- |
| **POST** | `/article/` | Create a new article with payload validation |
| **GET** | `/article/:limit/:offset` | Fetch paginated articles list |
| **GET** | `/article/:id` | Fetch a specific article by ID |
| **PUT/PATCH** | `/article/:id` | Update an existing article by ID |
| **DELETE** | `/article/:id` | Delete an article by ID |

