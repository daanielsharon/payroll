# Running the Payroll System

## Setup

Create a `.env` file by copying the example configuration:

```bash
cp .env.example .env
```

Edit the `.env` file and update the following variables with your configuration:
   - `DB_*`: Database connection details
   - `*_PORT`: Ports for each service
   - `JWT_SECRET`: A secure secret key for JWT token generation
   - `HOST`: The host address for your services

Example:
```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=payroll
DB_PORT=5432
GATEWAY_PORT=8080
PAYROLL_PORT=8081
OVERTIME_PORT=8082
ATTENDANCE_PORT=8083
REIMBURSEMENT_PORT=8084
AUTH_PORT=8085
USER_PORT=8086
HOST=localhost
JWT_SECRET=your-secret-key
```

## Start Infrastructure

Start the PostgreSQL database and Jaeger:

```bash
docker-compose up -d
```

## Run Migrations

Initialize the database schema:

```bash
cd migration
go run main.go
cd ..
```

## Start Services

Open a separate terminal for each service and run:

```bash
# Gateway Service
cd gateway
go run main.go
```

```bash
# Auth Service
cd auth
go run main.go
```

```bash
# User Service
cd user
go run main.go
```

```bash
# Payroll Service
cd payroll
go run main.go
```

```bash
# Overtime Service
cd overtime
go run main.go
```

```bash
# Attendance Service
cd attendance
go run main.go
```

```bash
# Reimbursement Service
cd reimbursement
go run main.go
```

## Access

- API Gateway: http://localhost:8080
- Jaeger UI: http://localhost:16686

