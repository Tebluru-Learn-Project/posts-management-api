# Golang API Project Structure

Struktur project Golang untuk API kompleks dengan pattern:

> Controller → Service → Repository → Database

Pattern ini cocok untuk project yang scalable, mudah di-maintain, dan enak dipakai untuk team.

---

## Struktur Folder

```bash
go-api/
├── cmd/
│   └── api/
│       └── main.go
│
├── config/
│   ├── app.go
│   ├── database.go
│   └── env.go
│
├── internal/
│   ├── controller/
│   │   ├── auth_controller.go
│   │   └── user_controller.go
│   │
│   ├── service/
│   │   ├── auth_service.go
│   │   └── user_service.go
│   │
│   ├── repository/
│   │   ├── auth_repository.go
│   │   └── user_repository.go
│   │
│   ├── model/
│   │   ├── user.go
│   │   └── session.go
│   │
│   ├── dto/
│   │   ├── auth_request.go
│   │   ├── auth_response.go
│   │   ├── user_request.go
│   │   └── user_response.go
│   │
│   ├── route/
│   │   └── api.go
│   │
│   ├── middleware/
│   │   ├── auth_middleware.go
│   │   └── logging_middleware.go
│   │
│   ├── helper/
│   │   ├── response.go
│   │   ├── password.go
│   │   └── token.go
│   │
│   └── exception/
│       ├── error_handler.go
│       └── custom_error.go
│
├── pkg/
│   ├── logger/
│   │   └── logger.go
│   │
│   └── utils/
│       └── utils.go
│
├── migrations/
│   ├── 001_create_users_table.up.sql
│   ├── 001_create_users_table.down.sql
│   ├── 002_create_sessions_table.up.sql
│   └── 002_create_sessions_table.down.sql
│
├── .env
├── go.mod
└── go.sum
```

---

## Penjelasan Folder

### `cmd/`

Tempat entry point aplikasi.

```bash
cmd/api/main.go
```

Biasanya dipakai untuk bootstrap aplikasi:

- load config
- connect database
- init router
- start server

Kenapa dipisah?

Karena nanti project bisa punya banyak entry point:

- `cmd/api` → REST API
- `cmd/worker` → background worker
- `cmd/ws` → websocket listener

---

### `config/`

Tempat semua konfigurasi aplikasi.

Isi utama folder ini:

- load environment
- database connection
- app configuration

File yang umum ada:

- `env.go` → load `.env`
- `database.go` → koneksi database
- `app.go` → config global app

Folder ini adalah fondasi project.

---

### `internal/`

Semua business logic utama ada di sini.

`internal` artinya package private, tidak untuk dipakai di luar app.

Semua logic utama aplikasi sebaiknya ditaruh di dalam folder ini.

---

### `internal/controller/`

Handle HTTP request & response.

Tugas controller:

- baca request
- parse input
- panggil service
- return response

Controller **jangan isi business logic**.

Controller hanya jadi jembatan antara HTTP dan service.

---

### `internal/service/`

Tempat business logic utama.

Semua rule dan flow aplikasi ditaruh di sini.

Contoh:

- register user
- login user
- validate session
- logout user

Service adalah inti logic aplikasi.

---

### `internal/repository/`

Tempat semua query database.

Tugas repository:

- insert data
- update data
- delete data
- find data

Repository hanya fokus ke database.

Repository **jangan isi business logic**.

---

### `internal/model/`

Representasi struktur data / entity database.

Contoh:

- `user.go`
- `session.go`

Model merepresentasikan bentuk data dari database.

Model **bukan** tempat request validation atau response formatting.

---

### `internal/dto/`

DTO = Data Transfer Object

Dipakai untuk request & response payload.

Kenapa dipisah dari model?

Karena:

- request API ≠ schema database
- response API ≠ model database

Contoh:

- `CreateUserRequest`
- `LoginRequest`
- `UserResponse`

DTO menjaga agar API tetap clean dan aman.

---

### `internal/route/`

Tempat definisi route / endpoint API.

Contoh:

- `POST /login`
- `POST /register`
- `GET /users`
- `DELETE /logout`

Route dipisah supaya `main.go` tetap bersih.

---

### `internal/middleware/`

Tempat middleware HTTP.

Contoh:

- auth middleware
- logging middleware
- recovery middleware
- CORS middleware

Middleware berjalan sebelum request masuk ke controller.

---

### `internal/helper/`

Tempat helper kecil yang reusable.

Helper dipakai lintas layer.

Contoh:

- `response.go` → format JSON response
- `password.go` → hash/check password
- `token.go` → generate session token

Helper berisi utility kecil yang sering dipakai.

---

### `internal/exception/`

Tempat handling error global.

Dipakai untuk:

- custom error
- error formatting
- centralized error response

Supaya response error konsisten.

Contoh:

```json
{
  "message": "validation error",
  "errors": {
    "email": "email is required"
  }
}
```

---

### `pkg/`

Reusable package yang generic.

Bedanya dengan `internal`:

- `internal` → private app logic
- `pkg` → reusable package

Biasanya isi:

- logger
- utils
- generic helpers

Kalau suatu package bisa dipakai ulang lintas project, taruh di `pkg`.

---

### `migrations/`

Tempat file migration database.

Dipakai untuk versioning schema database.

Contoh:

- create users table
- create sessions table
- alter table
- add index

Format:

- `.up.sql` → apply migration
- `.down.sql` → rollback migration

Migration wajib untuk project production.

---

### `.env`

Tempat environment variable.

Contoh:

- app name
- app port
- db host
- db user
- db password

Contoh isi:

```env
APP_NAME=MyApp
APP_PORT=8080
APP_ENV=development

DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=12345678
DB_NAME=myapp
```

---

### `go.mod`

Dependency manager Go.

Fungsinya:

- define module name
- manage dependencies

Contoh dependency:

- gin
- mysql driver
- bcrypt

---

### `go.sum`

Checksum dependency Go.

File ini auto-generated dan dipakai untuk validasi dependency integrity.

Jangan dihapus.

---

## Flow Request

Alur request di aplikasi:

```text
HTTP Request
   ↓
Route
   ↓
Middleware
   ↓
Controller
   ↓
Service
   ↓
Repository
   ↓
Database
```

Response balik:

```text
Database
   ↓
Repository
   ↓
Service
   ↓
Controller
   ↓
HTTP Response
```

---

## Urutan Development

Urutan development yang paling aman:

```text
1. config
2. helper
3. model
4. repository
5. service
6. controller
7. route
8. middleware
```

Kenapa urut begitu?

Karena setiap layer bergantung ke layer sebelumnya.

---

## Kesimpulan

Struktur ini cocok untuk:

- scalable API
- auth system
- ecommerce backend
- payment service
- websocket service
- trading engine
- arbitrage monitor

Kelebihan struktur ini:

- clean
- scalable
- testable
- mudah maintain
- cocok untuk team
- tidak cepat berantakan

Ini struktur yang aman untuk project Golang yang serius.
