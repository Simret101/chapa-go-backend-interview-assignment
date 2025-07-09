
---

# 📘 Chapa Backend API Documentation

This documentation describes the **Chapa Backend API**, a Go-based service that simulates key functionalities of the [Chapa API](https://chapa.co/) for handling transactions and transfers. It is built with clean architecture principles, GORM for persistence, Redis caching for speed, and Docker for containerization.

---

## 🔗 Base URL

```
http://localhost:8080/api/v0
```

---

## 🏗 Architecture Overview

| Layer          | Responsibility                                                                   |
| -------------- | -------------------------------------------------------------------------------- |
| `domain`       | Models and interfaces (e.g., `Transaction`, `Transfer`, `UseCase`, `Repository`) |
| `usecases`     | Business logic (e.g., `TransactionUseCase`, `TransferUseCase`)                   |
| `controllers`  | HTTP request handlers using Gin                                                  |
| `repositories` | DB operations using GORM                                                       |
| `routes`       | Registers routes with dependency injection                                       |
| `pkg/config`   | Error handlers, Redis cache manager, environment loading                         |

---

## 🚀 Getting Started

### 1. 🔨 Build Docker Image

```bash
docker build -t chapa-backend .
```

### 2. 🐳 Run Chapa Backend

```bash
docker run -p 8080:8080 --name chapa-backend-container chapa-backend
```

---

## 🧠 Redis Setup (via Docker)

### 🧱 Start Redis with Docker

```bash
docker run -d --name redis -p 6379:6379 redis
```

### 🔧 Access Redis CLI (for debugging)

```bash
docker exec -it redis redis-cli
```

---

## 📂 API Endpoints

### 1. ✅ Initiate Transaction

**POST** `/transactions/initiate`

#### Request Body

```json
{
  "txn_ref": "txn12345",
  "amount": 100.00,
  "currency": "ETB",
  "customer_email": "customer@example.com",
  "description": "Payment for order #123"
}
```

#### Response

```json
{
  "message": "transaction initiated",
  "txn_ref": "txn12345"
}
```

---

### 2. 🔍 Verify Transaction

**GET** `/transactions/verify/{txn_ref}`

```json
{
  "id": 1,
  "txn_ref": "txn12345",
  "amount": 100.00,
  "currency": "ETB",
  "customer_email": "customer@example.com",
  "description": "Payment for order #123",
  "status": "success",
  "created_at": "...",
  "updated_at": "..."
}
```

---

### 3. 📩 Transaction Webhook

**POST** `/transactions/webhook`

#### Headers

```
Chapa-Signature: <HMAC-SHA256 signature>
```

#### Body

```json
{
  "event": "charge.success",
  "status": "success",
  "txn_ref": "txn12345",
  "amount": 100.00,
  "currency": "ETB",
  "customer_email": "customer@example.com",
  "description": "Payment for order #123"
}
```

---

### 4. 🚚 Initiate Transfer

**POST** `/transfers/initiate`

```json
{
  "transfer_ref": "transfer12345",
  "account": "1234567890",
  "amount": 500.00,
  "bank_code": "CBE"
}
```

---

### 5. 🔎 Verify Transfer

**GET** `/transfers/verify/{transfer_ref}`

```json
{
  "id": 1,
  "transfer_ref": "transfer12345",
  "account": "1234567890",
  "amount": 500.00,
  "bank_code": "CBE",
  "status": "success",
  "created_at": "...",
  "updated_at": "..."
}
```

---

### 6. 🔔 Transfer Webhook

**POST** `/transfers/webhook`

```json
{
  "event": "transfer.success",
  "status": "success",
  "transfer_ref": "transfer12345",
  "amount": 500.00,
  "currency": "ETB",
  "recipient_email": "recipient@example.com",
  "description": "Transfer for order #123"
}
```

---

## 🧪 Postman Collection

👉 [View Postman Docs](https://documenter.getpostman.com/view/42403797/2sB34eJN9V)

---

## ⚠️ Error Handling (Consistent Format)

All responses follow:

```json
{
  "error": "transaction not found"
}
```

### Common Errors (From `pkg/config/errors.go`)

| Error                      | HTTP Code |
| -------------------------- | --------- |
| ErrBadRequest              | 400       |
| ErrTransactionNotFound     | 404       |
| ErrInvalidWebhookSignature | 401       |
| ErrPaymentAlreadyHandled   | 422       |
| ErrInternalServer          | 500       |

---

## ✅ Final Notes

* 📦 Built using Go, Gin, GORM, Redis
* 🧼 Clean architecture 
* 🔐 Secure webhook validation with HMAC-SHA256
* 🚀 Ready for production deployment via Docker

---

