![GraphQL Logo](assets/image2.png)![GraphQL Logo](assets/image3.png)![GraphQL Logo](assets/image1.png)

# Author Management System
This repository is a practice and example for working with GraphQL API, contains the source code for the Author Management System, a GraphQL-based API for managing authors, books, genres, and languages.

## Table of Contents
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)


## Getting Started

These instructions will help you set up and run the project on your local machine for check code and testing purposes.

### Prerequisites

- Go 1.16 or higher
- Docker (for running the database)
- Make (optional, for using the Makefile)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/AssassinRobot/author.git
    cd author
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Set up the database:
    ```sh
    docker-compose up -d
    ```

### Usage

1. Run the application:
    ```sh
    go run main.go
    ```

2. Access the GraphQL playground at `http://localhost:8000`.
(Also you can deal with GraphQL direct by sending POST request to `http://localhost:8000/query` )

## Project Structure

- `main.go`: Entry point of the application.
- `config/`: Configuration files.
- `database/`: Database connection and repository implementations.
- `graph/`: GraphQL resolvers and schema.
- `internal/model/`: internal(database) models.
- `internal/repository/`: Repository interfaces.
