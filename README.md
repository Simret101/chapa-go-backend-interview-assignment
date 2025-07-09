
# Chapa Backend Clone API

This is a Go-based backend service that replicates core functionalities of the Chapa payment API, including transactions, transfers, and secure webhook handling using clean architecture principles.

---

## 📌 Features

- ✅ Initiate & verify transactions  
- ✅ Initiate & verify transfers  
- ✅ Webhook handling with HMAC-SHA256 signature verification  
- ✅ Redis caching support  
- ✅ Dockerized for container-based deployment  
- ✅ Clean architecture (Domain, Usecases, Controllers, Repositories)

---
```
## 📁 Folder Structure

├── api
│   ├── controllers
│   │   ├── transaction.controller.go
│   │   └── transfer.controller.go
│   ├── middleware
│   └── routes
│       ├── routes.go
│       ├── transaction.routes.go
│       └── transfer.routes.go
│   ├── cmd
│   │   ├── main.go
│   │   └── .env
├── docs
│   ├── Chapa Backend.postman_collection.json
│   └── Documentation.MD
├── internal
│   ├── domain
│   │   ├── transaction.go
│   │   ├── transfer.go
│   │   ├── transaction.interface.go
│   │   ├── transfer.interface.go
│   │   ├── webhookhandler.interface.go
│   │   ├── webhooklog.go
│   │   └── cache.interface.go
│   ├── repositories
│   │   ├── transaction.repository.go
│   │   └── transfer.repository.go
│   └── usecases
│       ├── transaction.usecase.go
│       └── transfer.usecase.go
├── migration
├── mocks
├── pkg
│   └── config
│       ├── error.go
│       └── cache.go
├── scripts
├── storage
├── tmp
├── utils
├── go.mod
├── go.sum
├── Dockerfile
└── README.md

```

---

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL
- **Cache**: Redis (via Docker)
- **Containerization**: Docker

---

## 🐳 Docker Setup

### 🔧 Dockerfile (multi-stage build)
Already included in the project root.

### ✅ Build Docker image:

```bash
docker build -t chapa-backend .
````

### ✅ Run Docker container:

```bash
docker run -p 8080:8080 chapa-backend
```

### ✅ Run Redis (via Docker):

```bash
docker run --name redis -p 6379:6379 -d redis
```

### ✅ Access Redis CLI:

```bash
docker exec -it redis redis-cli
```

---

## 🔗 API Endpoints

| Endpoint                           | Method | Description                   |
| ---------------------------------- | ------ | ----------------------------- |
| `/api/v0/transactions/initiate`    | POST   | Initiate a transaction        |
| `/api/v0/transactions/verify/:ref` | GET    | Verify transaction by ref     |
| `/api/v0/transactions/webhook`     | POST   | Webhook for transaction event |
| `/api/v0/transfers/initiate`       | POST   | Initiate a transfer           |
| `/api/v0/transfers/verify/:ref`    | GET    | Verify transfer by ref        |
| `/api/v0/transfers/webhook`        | POST   | Webhook for transfer event    |

> For full documentation, refer to [`docs/Documentation.MD`](./docs/Documentation.MD) and Postman collection.

---

## 🧪 Postman Collection

* Import `docs/Chapa Backend.postman_collection.json` into Postman to test all endpoints easily.

---

## 🔐 Webhook Security

Webhooks are verified using HMAC-SHA256. Ensure to set `CHAPA_WEBHOOK_SECRET` correctly in `.env`.

Sample Header:

```http
Chapa-Signature: <HMAC-SHA256 signature>
```

---

## 📦 Redis Caching

Redis is used for caching verification status or preventing replay attacks.
Code for caching is found in `pkg/config/redis.go`.

---

## 🚀 Running Locally

```bash
go mod tidy
go run ./cmd/main.go
```

Access API: `http://localhost:8080/api/v0`

---

```
