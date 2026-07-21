# 🔗 URL Shortener

A lightweight URL shortening service written in **Go** using **Gin** and **Redis**.

The application allows you to:

- Create short URLs
- Retrieve original URLs
- Redirect users using short links
- Track link clicks
- View all stored short links

---

## 🚀 Tech Stack

- Go 1.26+
- Gin Web Framework
- Redis
- Docker Compose
- slog (structured logging)

---

## 📁 Project Structure

```
.
├── cmd/                    # Application entry point
├── configs/                # Configuration files
├── internal/
│   ├── config/             # Configuration loader
│   ├── handlers/           # HTTP handlers
│   ├── models/             # Data models
│   ├── repositories/       # Redis repository
│   └── services/           # Business logic
├── docker-compose.yaml
├── Taskfile.yaml
├── go.mod
└── .env
```

---

## ⚙️ Requirements

- Go 1.26+
- Docker
- Docker Compose

---

## 🔧 Installation

Clone the repository:

```bash
git clone <repository-url>
cd url-shortener
```

Download dependencies:

```bash
go mod download
```

---

## 🐳 Run Redis

Start Redis using Docker Compose:

```bash
docker compose up -d
```

---

## ▶️ Run the application

Using Go:

```bash
go run ./cmd -config=./configs/config.yaml
```

Or using Taskfile:

```bash
task run
```

The server starts on:

```
http://localhost:8000
```

---

## 🔌 API

### Create short URL (POST)

```
POST /v1/set
```

Request body:

```json
{
    "url": "https://google.com"
}
```

Response:

```json
{
    "url": "abc123"
}
```

---

### Create short URL (GET)

```
GET /v2/set?url=https://google.com
```

---

### Get original URL

```
GET /v1/g/{shortUrl}
```

Example:

```
GET /v1/g/abc123
```

Response:

```json
{
    "url": "https://google.com"
}
```

---

### Redirect to original URL

```
GET /v2/g/{shortUrl}
```

Example:

```
GET /v2/g/abc123
```

Returns:

```
301 Moved Permanently
```

and redirects the client to the original URL.

---

### List all short URLs

```
GET /v2/all
```

Returns all stored shortened links.

---

## 📊 Features

- ✅ URL shortening
- ✅ Permanent redirects
- ✅ Redis storage
- ✅ Click tracking
- ✅ Structured logging
- ✅ REST API
- ✅ Docker support

---

## ⚙️ Configuration

The application loads its configuration from:

```
configs/config.yaml
```

Environment variables are stored in:

```
.env
```

Redis configuration is defined in:

```
docker-compose.yaml
```

---

## 📝 Logging

The project uses Go's `log/slog`.

Supported log levels:

- local
- dev
- prod

---

## Future Improvements

- User authentication
- Custom aliases
- URL expiration
- Analytics dashboard
- Swagger/OpenAPI documentation
- Rate limiting
- Unit and integration tests
- Docker image for the application

---