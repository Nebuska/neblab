# Task Tracker

Task management backend api project written in go.

---
## Table of Contents

1. [Project Overview](#project-overview)
2. [Features](#features)
3. [Tech Stack](#tech-stack)
4. [Architecture](#architecture)
5. [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Running the Application](#running-the-application)
6. [API Documentation](#api-documentation)
    - [Authentication](#authentication)
    - [Endpoints](#endpoints)
7. [Testing](#testing)
8. [Environment Variables](#environment-variables)
9. [Contributing](#contributing)
10. [License](#license)

---

## Project Overview

Not yet written

---

## Features

- Board based task management
- Shared task boards
- JWT based user authentication
- Dependency Injection
- Structured logging


---

## Tech Stack

- **Language:** Go 1.25.3
- **Web Framework:** Gin v1.11.0
- **Dependency Injection** fx v1.24.0
- **Database:** MySQL
- **ORM:** GORM v1.31.1
- **Logger:** Zerolog v1.34.0
- **Configuration:** godotenv v1.5.1

---

## Architecture

///Insert diagram or explanation of layers
HTTP Handler --> Service Layer --> Repository / Database
- **Handler:** Handles HTTP requests and responses
- **Service:** Business logic
- **Repository:** Database operations

---

## Getting Started

### Prerequisites

- **git:** for cloning the repository
- **Go:** 1.25 or higher
- **Database:** MySQL (Can use other sql databases with minimal change)

### Installation

1. **Clone the repository and download dependencies**

```bash
  git clone https://github.com/Nebuska/task-tracker.git
  cd task-tracker
  go mod download
```

### 2. Set up environments

- Create a `.env` file in the root directory using `.env.example`
- Adjust values according to your setup. [See how](#environment-variables)

### 3. Set up database

- Create a database for the application (e.g., tasktracker)
- 



### Running the Application

#### Without Docker
```bash
    make all
```
By default, the server should start on http://localhost:8080 (or the port you configured).

#### With Docker

- Not implemented yet

---

## Api Documentation

- Not finished documentation only short example
- Might use Swagger/OpenAPI later



### Authentication

### Endpoints

| Method | Path                   | Description                         | Auth Required |
|--------|------------------------|-------------------------------------|---------------|
| POST   | /api/v1/register       | Create an account                   | No            |
| POST   | /api/v1/login          | Get jwt token for account           | No            |
| GET    | /api/v1/boards         | List every board user has access    | Yes           |
| GET    | /api/v1/boards/:id     | Get the details of a board          | Yes           |
| POST   | /api/v1/boards         | Create new board                    | Yes           |
| DELETE | /api/v1/boards/:id     | Delete board                        | Yes           |
| GET    | /api/v1/tasks?:filters | List all tasks according to filters | Yes           |
| POST   | /api/v1/tasks          | Create new task                     | Yes           |
| PUT    | /api/v1/tasks/:id      | Update task                         | Yes           |
| DELETE | /api/v1/tasks/:id      | Delete task                         | Yes           |

---

## Testing

---

## Environment Variables

---

## Contributing

---

## License