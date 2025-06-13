# Payroll Microservice System

A comprehensive microservice architecture for enterprise payroll management built with Go, featuring distributed tracing, role-based access control, and audit logging.

## System Architecture

This system implements a modern microservices architecture with the following components:

- **API Gateway**: Single entry point that routes requests to appropriate services
- **Service Mesh**: Inter-service communication with context propagation
- **Distributed Tracing**: End-to-end request tracking with OpenTelemetry and Jaeger
- **Role-Based Access Control**: Admin and Employee permission levels
- **Audit Logging**: Comprehensive activity tracking across all services
- **Database Per Service**: Each microservice manages its own data store

## Services

### Core Services

- **Gateway Service** (Gin): Routes requests to appropriate microservices, handles request tracing, and forwards authentication headers
- **Auth Service**: Manages authentication, JWT token generation, and validation
- **User Service**: Handles user management and profile information

### Business Logic Services

- **Payroll Service** (Chi): Manages pay periods, and coordinates with other services during payroll runs
- **Overtime Service** (Chi): Calculates overtime hours based on attendance data
- **Attendance Service** (Chi): Tracks employee clock-in/clock-out times and validates attendance records
- **Reimbursement Service** (Chi): Processes expense reimbursement requests

### Distributed Tracing & Monitoring

The system implements OpenTelemetry for distributed tracing with:
- Jaeger for trace visualization
- OTel Collector for telemetry data processing
- Context propagation across service boundaries
- Span correlation for end-to-end request tracking

### Role-Based Access Control

Two primary roles are supported:
- **Admin**: Full system access with management capabilities
- **Employee**: Limited access to personal records and submissions

### Inter-Service Communication

- HTTP-based service communication with context propagation
- Shared client package for standardized service calls
- Header propagation for maintaining trace context and user information

### Audit Logging

- Comprehensive audit trail for all data modifications
- Records user ID, IP address, request ID, and trace ID
- Stores before/after state for all changes

## Setup Instructions

### Prerequisites

- Go 1.18 or higher
- PostgreSQL 13 or higher
- Jaeger and OTel Collector for tracing (optional but recommended)

### Environment Setup

1. Create a `.env` file by copying the example configuration:

```bash
cp .env.example .env
```

2. Edit the `.env` file and update the following variables with your configuration:
   - `DB_*`: Database connection details
   - `*_PORT`: Ports for each service
   - `JWT_SECRET`: A secure secret key for JWT token generation
   - `HOST`: The host address for your services