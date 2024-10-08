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

## Scope

The backend for this project is built using Golang, Gin framework, and PostgreSQL as the database. It includes the following key features and functionalities:

1. User Management:
CRUD operations for users (Create, Read, Update, Delete).
Role-based access control (RBAC) with Super Admin, Admin, and Normal User roles.
JWT-based authentication for login and secure access to protected routes.
Password management (login, forgot password).
2. Company Management:
CRUD operations to manage companies.
Each user is linked to a specific company, and permissions are scoped to their company.
3. Employee Management:
CRUD operations for employees.
Employees can be assigned to companies and linked to specific key copies.
Employees do not need a user account, but key copies can be assigned to them.
4. Building and Lock Management:
Manage buildings with detailed CRUD operations for:
Buildings
Floors
Rooms
Locks
Each building can have multiple floors, rooms, and locks.
5. Key Copy Management:
CRUD operations for managing key copies.
Assign and revoke key copies to employees.
Ensure that key copies can only be assigned within the company’s scope.
6. Multi-Tenancy:
The system is built with multi-tenancy in mind, ensuring that:
Each company’s data is isolated.
Users, employees, and key copies are restricted to their respective companies.
Super Admins have broader access to manage multiple companies.
7. Global Search:
Implement a global search functionality to quickly locate users, buildings, rooms, and locks across the system, restricted by company scope.
8. Security:
JWT authentication for secure access.
Password hashing for user credentials.
Role-based access control to restrict features based on user roles.
9. API Endpoints:
RESTful API design, with clean endpoints for users, companies, employees, buildings, floors, rooms, locks, and key copies.
Pagination and filtering support for listing resources.

## Reference

[Assets on Spreadsheet](https://docs.google.com/spreadsheets/d/1_23qT2cwyXJeu0vR586Iv_6GSm21OEDjbDPyKq7xWv8/edit?usp=sharing)