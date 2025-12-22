# Go Kit 2026 — Clean Architecture REST API Starter Kit

Starter kit **Go RESTful API** dengan **Clean Architecture**, terinspirasi dari:
- Go-kit style architecture
- Laravel (Controller / Service / Repository)
- Laravel Jetstream / Fortify
- Spatie Role & Permission

Dirancang untuk **scalable, testable, dan production-ready**.

---

## ✨ Features

- ✅ Clean Architecture (Controller / Service / Repository / Domain)
- ✅ Gin HTTP Framework
- ✅ GORM ORM (MariaDB default, DB-switch ready)
- ✅ JWT Authentication (Access & Refresh Token)
- ✅ Role & Permission (Spatie-like RBAC)
- ✅ Centralized Error Handling
- ✅ Transaction-safe Service Layer
- ✅ SQL Migration & Seeder
- ✅ API Versioning (`/api/v1`)
- ✅ Dependency Injection (explicit, no magic)
- ✅ Tested via Postman

---

## 🧱 Architecture Overview

```text
Controller (HTTP, v1)
   ↓
Handler (HTTP response & error mapping)
   ↓
Service (business logic, transaction orchestration)
   ↓
Domain (entity + repository interface)
   ↑
Repository (GORM implementation)

## 📂 Struktur Folder

```text
go-kit-2026/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   └── app/
│       ├── bootstrap/              # dependency wiring & app init
│       │   ├── app.go
│       │   ├── database.go
│       │   └── router.go
│       │
│       ├── config/                 # env-based configuration
│       │   ├── app.go
│       │   ├── database.go
│       │   ├── jwt.go
│       │   ├── redis.go
│       │   └── config.go
│       │
│       ├── controller/             # HTTP controllers
│       │   └── v1/
│       │       ├── auth_controller.go
│       │       ├── admin_controller.go
│       │       └── profile_controller.go
│       │
│       ├── handler/                # HTTP response & error presenter
│       │   └── error.go
│       │
│       ├── service/                # business logic
│       │   ├── auth_service.go
│       │   ├── auth_service_impl.go
│       │   ├── authorization_service.go
│       │   └── errors.go
│       │
│       ├── domain/                 # core business domain
│       │   ├── entity/
│       │   │   ├── user.go
│       │   │   ├── role.go
│       │   │   ├── permission.go
│       │   │   └── refresh_token.go
│       │   │
│       │   └── repository/         # repository interfaces
│       │       ├── user_repository.go
│       │       ├── role_repository.go
│       │       ├── permission_repository.go
│       │       └── refresh_token_repository.go
│       │
│       ├── repository/             # infra layer
│       │   ├── gorm/
│       │   │   ├── user_repository.go
│       │   │   ├── role_repository.go
│       │   │   ├── permission_repository.go
│       │   │   └── refresh_token_repository.go
│       │   └── transaction.go
│       │
│       ├── middleware/             # gin middleware
│       │   ├── auth_jwt.go
│       │   ├── role.go
│       │   ├── permission.go
│       │   └── logging.go
│       │
│       ├── dto/                    # request & response DTO
│       │   ├── request/
│       │   │   └── auth_request.go
│       │   └── response/
│       │       ├── base_response.go
│       │       └── auth_response.go
│       │
│       ├── routes/                 # route registration
│       │   └── v1.go
│       │
│       └── utils/                  # helpers
│           ├── password.go
│           └── jwt.go
│
├── migrations/                     # SQL migration & seeder
│   ├── 0001_create_users.sql
│   ├── 0002_create_roles.sql
│   ├── 0003_create_permissions.sql
│   ├── 0004_create_user_roles.sql
│   ├── 0005_create_role_permissions.sql
│   ├── 0006_create_refresh_tokens.sql
│   └── 1000_seed_roles_permissions.sql
│
├── .env.example
├── go.mod
├── go.sum
└── README.md
