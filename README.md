# File Sharing API with JWT Auth, Upload, and Public Sharing 

---

## Project Overview

This project is a **secure file sharing system** built with a Golang backend and a modern React frontend. It allows users to:

- Upload files (currently `.txt` only).
- Store and list their uploaded files.
- Generate public shareable links.
- Download files securely.

The system features JWT-based authentication, protected routes, and a modern, responsive UI using Tailwind CSS.

---

## Features

- JWT-based Authentication  
- File Upload (Only `.txt`)  
- File Listing  
- Public File Sharing via Tokenized URLs  
- Secure File Download  
- Auth-Protected Routes  
- GitHub Actions CI/CD
- Prometheus Exporter  
- Grafana Dashboard
---

## Component Details

### 1. Authentication

- Users register/login using `/register` and `/login`.
- JWT tokens are returned and stored in `localStorage`.
- Protected endpoints check for valid tokens using middleware.

### 2. File Upload

- Authenticated users upload `.txt` files via `multipart/form-data`.
- Files are stored on disk or cloud (configurable).

### 3. File Listing

- Authenticated users can view their uploaded files.
- Each file entry includes a share button and download option.

### 4. Public Sharing

- Clicking "Share" generates a public URL like `/share/<token>`.
- Anyone with this link can download the file.

### 5. Monitoring

- Prometheus endpoint `/metrics`
- Grafana dashboard shows:
  - Requests/sec
  - Response status codes
  - File upload/download count
---

## Technology Stack

### Frontend
- **React** + **TypeScript**
- **Tailwind CSS**
- **React Router DOM**
- **Axios** for HTTP requests

### Backend
- **Golang**
- **PostgreSQL** (with GORM ORM)
- **JWT** for Authentication

### Monitoring
- **Prometheus**
- **Grafana**

### DevOps
- **GitHub Actions (CI/CD)**
- **Docker + Docker Compose**

---

## Setup and Installation

### Prerequisites

- Go 1.18+
- Node.js 18+
- PostgreSQL
- Docker & Docker Compose
- Prometheus + Grafana

---

### Backend Setup

```bash
git clone https://github.com/shantanugupta2004/HTTP-Server-Go.git
cd backend


go mod tidy
go run main.go
```

Example `.env`:

```
PORT=5000
DATABASE_URL=postgres://user:pass@localhost:5432/files_db
JWT_SECRET=your-secret
```
---

### Frontend Setup

```bash
cd frontend
pnpm install         # or npm install / yarn install
pnpm run dev         # or npm run dev / yarn dev
```

Ensure the API base URL in `utils/api.ts` points to your backend (e.g., `http://localhost:5000`).

---

### Docker Compose Setup

Add a `docker-compose.yml` to run:

- Backend  
- PostgreSQL  
- Prometheus  
- Grafana  

```bash
docker-compose up --build
```

> Example services in `docker-compose.yml`:  
> `api`, `db`, `prometheus`, `grafana`

---

## API Endpoints

### Authentication

#### `POST /register`
- **Description**: Register a new user.
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Response**: `{ "token": "JWT_TOKEN" }`

#### `POST /login`
- **Description**: Log in and receive a JWT.
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Response**: `{ "token": "JWT_TOKEN" }`

---

### File Upload

#### `POST /upload`
- **Description**: Upload a `.txt` file (authenticated).
- **Headers**:  
  `Authorization: Bearer <JWT>`
- **Body**: `multipart/form-data` with `file=<File>`
- **Response**:
  ```json
  {
    "message": "File uploaded successfully",
    "fileId": "uuid"
  }
  ```

---

### File Listing

#### `GET /files`
- **Description**: Get list of user's uploaded files.
- **Headers**:  
  `Authorization: Bearer <JWT>`
- **Response**:
  ```json
  [
    {
      "id": "uuid",
      "filename": "example.txt",
      "uploadedAt": "2025-07-10T12:00:00Z"
    }
  ]
  ```

---

### Public Sharing

#### `GET /share/:token`
- **Description**: Public download route.
- **Access**: No auth required.
- **Response**: Returns the file directly for download.

---

### File Download

#### `GET /download/:fileId`
- **Description**: Download your own file.
- **Headers**:  
  `Authorization: Bearer <JWT>`
- **Response**: File stream (download)
---
### Monitoring
- `GET /metrics` — Prometheus-compatible metrics
---
## Usage

1. **Register** or **Login** to get a JWT token.
2. Upload a `.txt` file.
3. View the list of uploaded files.
4. Click "Share" to get a public download link.
5. Others can use the share link to download the file without login.

---
## CI/CD with GitHub Actions

This project uses **GitHub Actions** for:

- Linting and formatting
- Running unit tests
- Building backend Docker image
- Pushing to Docker Hub or any registry

### Example Workflow: `.github/workflows/ci.yml`

```yaml
name: CI Pipeline

on:
  push:
    branches: [ main ]

jobs:
  build-and-test:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'

      - name: Install Dependencies
        run: go mod tidy

      - name: Run Tests
        run: go test ./...

      - name: Docker Build
        run: docker build -t your-dockerhub/file-sharing-backend .

      - name: Docker Push
        run: |
          echo "${{ secrets.DOCKER_PASSWORD }}" | docker login -u "${{ secrets.DOCKER_USERNAME }}" --password-stdin
          docker push your-dockerhub/file-sharing-backend
```

---

## Monitoring with Prometheus and Grafana

### Step 1: Expose Metrics in Go

Add in your backend:

```go
import "github.com/prometheus/client_golang/prometheus/promhttp"

r.GET("/metrics", gin.WrapH(promhttp.Handler()))
```

### Step 2: prometheus.yml config

```yaml
scrape_configs:
  - job_name: 'file-sharing-app'
    static_configs:
      - targets: ['api:5000']
```

### Step 3: Grafana Dashboard

- Connect Grafana to Prometheus (`http://prometheus:9090`)
- Import dashboard or create:
  - Upload count over time
  - Request duration
  - Status codes
  - Auth success/failures

### Access:
- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000` (default creds: `admin/admin`)

---
## References

- [React Docs](https://react.dev)
- [Go Docs](https://go.dev/)
- [GORM](https://gorm.io)
- [JWT](https://jwt.io/)
- [Prometheus](https://prometheus.io/)
- [Grafana](https://grafana.com/)
- [GitHub Actions](https://docs.github.com/en/actions)
---
