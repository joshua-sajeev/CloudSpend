# CloudSpend

CloudSpend is a personal project for learning Golang and Docker. The goal of this project is to explore backend development with Go, API handling, database management with PostgreSQL, and containerization using Docker.

## Overview

- **Built with Golang** using Gin and Gorilla Mux for handling API requests.
- **Uses PostgreSQL** as the database with GORM for ORM functionality.
- **Dockerized setup** with Docker Compose to manage services.
- **Basic transaction management** to experiment with CRUD operations.
- **Includes pgAdmin** for database administration.

## Features

- API for handling transactions.
- Database migrations using Makefile.
- Graceful shutdown handling in the server.
- Retry logic for database connections.
- Containerized environment for easy setup.

## Technologies Used

- **Golang** – Backend development.
- **Gin/Gorilla Mux** – API handling.
- **PostgreSQL** – Database.
- **GORM** – ORM for database interactions.
- **Docker & Docker Compose** – Containerization.
- **pgAdmin** – Database management.

## Setup

### Prerequisites

- Docker & Docker Compose installed.
- Go installed (if running locally without Docker).

### Steps

1. Clone the repository:

   ```sh
   git clone https://github.com/joshua-sajeev/cloudspend.git
   cd cloudspend
   ```

2. Create a `.env` file with database credentials.

3. Start the services using Docker Compose:

   ```sh
   docker-compose up --build
   ```

4. Access the API at `http://localhost:8080`.

## Notes

This project is primarily for learning purposes and not intended for production use.  
Development on **CloudSpend** has been paused for now. Currently working on [taskflow](https://github.com/joshua-sajeev/taskflow).
