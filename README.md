# Blogging Platform

## Project Setup and Execution

This README provides instructions to set up, build, and run the project with PostgreSQL and Go. Follow the steps below to initialize the database and run the Go application.

## Features

- CRUD operations for blog posts.
- Filtering posts by author or creation date.
- Pagination for large datasets.
- Basic JWT-based authentication and authorization.
- Role-based access control (admin and regular users).
- Unit tests for key functionalities.

## Technologies Used

Ensure you have the following installed:

- [Go](https://golang.org/) for the backend.
- [PostgreSQL](https://www.postgresql.org/) as the database.
- [Gorilla Mux](https://www.gorillatoolkit.org/pkg/mux) for routing.
- [JWT](https://jwt.io/) for authentication.
- [Make](https://www.gnu.org/software/make/) for build automation
- [Docker](https://www.docker.com/) (if you have Docker-related tasks)
- [Psql](https://www.postgresql.org/) (if you use Postgresql directly )


## Prerequisites

Ensure you have the following installed:
- Docker and Docker Compose
- Go (version 1.18 or higher)
- psql ( version 15.1 and above, if you use database without docker)


## Project Structure

Here's an overview of the project structure:

```plaintext
blogging-platform/
├── cmd/              
    ├── main.go          # Entry point for the API server
├── handler/  
    ├── apihandler/          # API endpoint handlers
    ├── middleware/          # Middleware for authentication/authorization
    ├── authservice/         # Provides AuthServices
    ├── config/              # Configuration files
    ├── logger/              # Error handling
    ├── models/              # Data structures
    ├── psql/                # Database structure
    ├── models/              # Data structures
    ├── routs/               # Roting handler
├── tests/                   # Unit and integration tests
├── thunder_clint/           # API testing
├── docs/                    # API documentation
├── blogger_schema.sql       # Unit and integration tests
├── config,yml               # for config the environments
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
└── Makefile                 # Makefile for common tasks like testing and building
```


## Installation
To install and run this project, follow these steps:
1. **Clone the repository:**
   ```bash
   git clone https://github.com/KIRANKUMAR_HS/blogging-platform.git
   cd blogging-platform
   ```


## Setting Up PostgreSQL with Docker

To set up PostgreSQL with Docker, follow these steps:

1. **Build the Docker Image**:
   - Open a terminal in the project's root directory.
   - Run the following command to build the PostgreSQL image:
     ```bash
     docker build -t my-postgres-image .
     ```

2. **Run the Docker Container**:
   - Run the container and map port 5432:
     ```bash
     docker run -d --name my-postgres-container -p 5432:5432 my-postgres-image
     ```

3. **Access PostgreSQL**:
   - Connect to PostgreSQL with a client like `psql`:
     ```bash
     psql -h localhost -p 5432 -U myuser -d mydatabase
     ```


## Creating the Database and Tables

Once connected to PostgreSQL, create the necessary tables:

1. **Create the `users` Table**:
   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(50) UNIQUE NOT NULL,
       password_hash VARCHAR(100) NOT NULL,
       role VARCHAR(20) NOT NULL,
       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
   );
2. **Create the `users` Table**:
   ```sql
   CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    author_id INT REFERENCES users(id),  -- Foreign key to `users`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );


## Alternatively Or you can also use the Psql Databse running in your locall system
Reffer to the official documentation of Postgresql to install an and run psql Database:

1. **Create Database with user and posts tables**:
   - create PostgreSQL with using .sql file:
     ```bash
     psql -U your_user_name -h localhost -p 2909 -d postgres -f blogger_schema.sql
     ```

## Running the Application

With PostgreSQL set up and the tables created,and along wityh the proper configfile you can build and run the Go application:
To build the application, use the `build` target in the `Makefile`:


1. **Build the application**:
   - Building application using make file:
     ```bash
     make build
     ```
    
2. **start the application**:
   - Starting the application using make file:
     ```bash
     make start
     ```

3. **clean the app buildlication**:
   - To clean build artifacts, use the clean target in make file:
     ```bash
     make clean
     ```

4. **Run the application without building**:
   - Running application usig make file:
     ```bash
     make run
     ```

5. ** To perform unit test use the **:
    - Testing the application using make file:
     ```bash
     make test
     ```

## API usage of application
For detailed API usage instructions, see the [API documentation](docs/api_doc.md):
