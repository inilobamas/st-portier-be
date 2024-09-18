# st-portier-be
# Key Management Backend

This project is a backend implementation for managing buildings, floors, rooms, locks, and key copies in a multi-tenant environment. The system is built using Golang, Gin, PostgreSQL, and JWT for authentication.

## Features

- **Multi-Tenancy Support**: Data isolation across companies using a `company_id` field.
- **CRUD Operations**: Supports Create, Read, Update, and Delete operations for buildings, floors, rooms, and locks.
- **Key Copy Assignment**: Functionality for assigning and revoking key copies to employees.
- **Role-Based Access Control (RBAC)**: Different levels of access for Super Admins, Admins, and Normal Users.
- **Global Search**: Search across buildings, floors, rooms, locks, and employees within a company.

## Prerequisites

- Go
- PostgreSQL
- Git

## Installation

#1. Clone the repository:

   ```bash
   git clone https://github.com/inilobamas/st-portier-be.git
   cd your-project
   ```

#2. Install dependencies:
   ```bash
   go mod tidy
   ```

#3. Set up PostgreSQL:
   ```sql
   CREATE DATABASE your_db_name;
   CREATE USER your_user WITH ENCRYPTED PASSWORD 'your_password';
   GRANT ALL PRIVILEGES ON DATABASE your_db_name TO your_user;
   ```

#4. Create a .env file with the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=your_db_name
   JWT_SECRET=your_secret
  ```

#5. Running the Application
   ```bash
   go run main.go
   ```

   The application will start on localhost:8080.

## Reference

[Assets on Spreadsheet](https://docs.google.com/spreadsheets/d/1_23qT2cwyXJeu0vR586Iv_6GSm21OEDjbDPyKq7xWv8/edit?usp=sharing)