<<<<<<< HEAD
# go-kit-2026
Go Kit 2026 — Clean Architecture REST API Starter Kit
=======
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
>>>>>>> 7adc4aa (chore: add readme.md)
