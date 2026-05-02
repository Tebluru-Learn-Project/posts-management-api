# Golang REST API Project

A robust and scalable REST API built with [Golang](https://golang.org/) and the [Gin Web Framework](https://gin-gonic.com/). This project implements a clean layered architecture (`Controller -> Service -> Repository -> Database`) and provides core backend features such as authentication, profile management, and file uploads.

## ✨ Features

- **Layered Architecture:** Clean separation of concerns (Controllers, Services, Repositories) for maintainability and scalability.
- **Authentication & Authorization:** 
  - User Registration with Email OTP verification.
  - Secure Login with Session/Token management.
  - Secure password hashing using Bcrypt.
- **Profile Management:** Get and update comprehensive user profile details including bio and location.
- **File Management:** Secure file and avatar uploads (with MIME type & file size validation).
- **Database Migrations:** Structured SQL-based migrations.
- **Standardized Responses:** Consistent JSON formatting for all successes and errors.

## 🛠 Tech Stack

- **Language:** Go 1.25+
- **Framework:** Gin Web Framework (`github.com/gin-gonic/gin`)
- **Database:** MySQL (`github.com/go-sql-driver/mysql`)
- **Security:** Bcrypt (`golang.org/x/crypto/bcrypt`)
- **Other Utilities:** Google UUID, standard `net/smtp` for email delivery.

## 🚀 Getting Started

### Prerequisites

- [Go 1.25+](https://golang.org/dl/) installed on your machine.
- A running MySQL database server.

### Installation & Setup

1. **Clone the repository:**
   ```bash
   git clone <repository_url>
   cd go-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Environment Setup:**
   Copy the environment template and configure your database and SMTP credentials.
   ```bash
   cp .env.example .env
   ```
   *Edit the `.env` file to match your local setup.*

4. **Database Migration:**
   Execute the `.up.sql` scripts located in the `migrations/` folder against your MySQL database sequentially to set up the necessary tables (`users`, `sessions`, `profiles`, `files`).

5. **Run the Application:**
   ```bash
   go run cmd/api/main.go
   ```
   The server will start on the port defined in your `.env` (default is `5001`).

## 📁 Project Structure Overview

```text
go-api/
├── cmd/             # Application entry points (e.g., main.go)
├── config/          # Environment, App, and Database configuration
├── internal/        # Private application and business logic
│   ├── controller/  # Handles HTTP requests/responses
│   ├── service/     # Core business logic
│   ├── repository/  # Database queries and interactions
│   ├── model/       # Database entities
│   ├── dto/         # Data Transfer Objects (Requests & Responses)
│   ├── route/       # API Endpoint definitions
│   ├── middleware/  # HTTP Middlewares (Auth, etc.)
│   └── helper/      # Reusable utilities (Password, Validation, Files)
├── migrations/      # SQL up/down migration files
└── pkg/             # Public/shared libraries
```
*(For detailed architectural explanations, please refer to the `golang-api-project-structure.md` document)*

## 📡 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register a new user and trigger OTP email.
- `POST /api/v1/auth/verify-otp` - Verify the account using the OTP code.
- `POST /api/v1/auth/login` - Authenticate user and retrieve session token.

### User & Profile (Requires Auth Header)
- `GET  /api/v1/user/me` - Get basic information of the currently authenticated user.
- `GET  /api/v1/user/profile` - Get detailed user profile data.
- `PUT  /api/v1/user/profile` - Update user profile information.
- `POST /api/v1/user/avatar` - Upload a new user avatar image.

## 📝 License

This project is open-source and available under the [MIT License](LICENSE).
