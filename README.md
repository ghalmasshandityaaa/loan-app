# Loan APP Api

A secure backend application for loan management built with **Golang (Fiber Framework)** and **MySQL**. Features include loan applications, admin approvals, and user management with OWASP Top 10-compliant security layers.

## 📋 Table of Contents

- [🏗️ Architecture Overview](#architecture-overview)
- [📁 Entity RelationShip Diagram](#erd-overview)
- [🔒 Security Features (OWASP Top 10)](#security-features-owasp-top-10)
- [🔌 API Endpoints](#api-endpoints)
- [🚀 Setup Instructions](#setup-instructions)


## 🏗️ Architecture Overview 

### System Architecture
```mermaid
graph TB
    subgraph "External Interfaces"
        A[HTTP API<br/>Fiber]
        B[Database<br/>MySQL]
    end
    
    subgraph "Application Layer"
        D[Handlers]
        E[UseCases]
        F[Middleware]
    end
    
    subgraph "Domain Layer"
        G[Entities]
        H[Repositories]
    end
    
    A --> D
    B --> H
    D --> E
    E --> H
    F --> D
    
    style A fill:#e3f2fd
    style D fill:#f3e5f5
    style G fill:#e8f5e8
```

### Project Structure

```
loan-app/
├── cmd/                    # Application entry point
│   └── main.go            # Main application bootstrap
├── config/                 # Configuration management
│   ├── config.go          # Configuration loader
│   ├── config.json        # Application configuration
│   └── types.go           # Configuration types
├── internal/              # Internal application code
│   ├── app/              # Application bootstrap
│   ├── entity/           # Domain entities/models
│   ├── handler/          # HTTP request handlers
│   ├── middleware/       # Custom middleware
│   ├── model/            # Request/response models
│   ├── repository/       # Data access layer
│   ├── route/            # Route definitions
│   ├── usecase/          # Business logic layer
│   ├── utils/            # Utility functions
│   └── vm/               # View models
├── pkg/                   # Reusable packages
│   ├── database/         # Database utilities
│   ├── fiber/            # Fiber framework setup
│   ├── logger/           # Logging utilities
│   ├── middleware/       # Common middleware
│   └── validator/        # Validation utilities
├── docs/                  # Documentation
│   └── postman/          # Postman collections
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── Makefile              # Build and deployment scripts
└── README.md             # Project documentation
```

## 📁 Erd Overview

```mermaid
erDiagram
    users {
        CHAR(26) id PK
        VARCHAR(16) nik UK "National ID number"
        VARCHAR(100) full_name
        VARCHAR(100) legal_name "Name as shown on ID"
        VARCHAR(100) place_of_birth
        DATE date_of_birth
        BIGINT salary
        TEXT id_card_photo_url
        TEXT selfie_photo_url
        TEXT password
        BOOLEAN is_admin "DEFAULT FALSE"
        TIMESTAMP created_at "DEFAULT CURRENT_TIMESTAMP"
        TIMESTAMP updated_at "NULL DEFAULT NULL"
    }

    customer_limits {
        CHAR(26) id PK
        CHAR(26) user_id FK
        SMALLINT tenor "CHECK (1,2,3,4)"
        BIGINT limit_amount "DEFAULT 0"
        BIGINT used_amount "DEFAULT 0"
        BIGINT available_amount "DEFAULT 0"
        DATETIME created_at "DEFAULT CURRENT_TIMESTAMP"
        DATETIME updated_at "NULL"
    }

    partners {
        VARCHAR(26) id PK
        VARCHAR(100) name UK
        ENUM partner_type "ecommerce, dealer"
        DATETIME created_at "DEFAULT CURRENT_TIMESTAMP"
        VARCHAR(26) created_by FK
        DATETIME updated_at "NULL"
        VARCHAR(26) updated_by FK
    }

    assets {
        VARCHAR(26) id PK
        VARCHAR(26) partner_id FK
        VARCHAR(100) name
        BIGINT price "CHECK price > 0"
        DATETIME created_at "DEFAULT CURRENT_TIMESTAMP"
        VARCHAR(26) created_by FK
        DATETIME updated_at "NULL"
        VARCHAR(26) updated_by FK
    }

    transactions {
        VARCHAR(26) id PK
        VARCHAR(26) user_id FK
        VARCHAR(26) asset_id FK
        VARCHAR(50) contract_number
        BIGINT otr_price
        BIGINT admin_fee "DEFAULT 0"
        BIGINT installment_amount
        BIGINT interest_amount
        DATETIME created_at "DEFAULT CURRENT_TIMESTAMP"
        VARCHAR(26) created_by FK
        DATETIME updated_at "NULL"
        VARCHAR(26) updated_by FK
    }

    %% Relationships
    users ||--o{ customer_limits : "has"
    users ||--o{ partners : "creates/updates"
    users ||--o{ assets : "creates/updates"
    users ||--o{ transactions : "creates/updates"
    partners ||--o{ assets : "owns"
    users ||--o{ transactions : "makes"
    assets ||--o{ transactions : "involved_in"
```

## 🔒 Security Features (OWASP Top 10)

This application has implemented several security features to protect against common attacks, including:

### Prevention of SQL Injection (OWASP A03)

* Use of prepared statements or parameterized queries to prevent SQL injection
* Use of a secure template engine to prevent template injection

### Prevention of Broken Authentication (OWASP A07)

* Implementation of authentication using JSON Web Tokens (JWT) to ensure secure data transmission
* Use of secure password hashing to store user passwords

### Security Misconfiguration Prevention (OWASP A05)

* Use of secure configuration for the application, including CORS and CSRF settings
* Use of the latest library and framework versions to ensure security

### Prevention of Insecure Direct Object References (IDOR) (OWASP A01)

* Implementation of role-based access control to ensure that only authorized users can access certain data
* Use of middleware to check user roles before accessing data

### Prevention of Cross-Site Request Forgery (CSRF) (OWASP A01)

* Implementation of CSRF tokens to ensure that requests sent by users are valid
* Use of middleware to check CSRF tokens before processing requests

### Prevention of API Security Risks (OWASP API Top 10) (Bonus for API)

* Implementation of JWT (JSON Web Tokens) to ensure secure data transmission
* Use of an API gateway to ensure security and access control to the API

## 🔌 API Endpoints

### Auth Endpoints

* `POST /v1/auth/sign-in`: Sign in user
* `POST /v1/auth/sign-up`: Sign up user

### User Endpoints

* `GET /v1/user/me`: Get self user information
* `GET /v1/user/limit`: Get user limits

### Partner Endpoints

* `POST /v1/partner`: Create partner
* `GET /v1/partner`: Get partner lists

### Asset Endpoints

* `POST /v1/asset`: Create asset
* `GET /v1/asset`: Get asset lists

### Transaction Endpoints

* `POST /v1/transaction`: Create User Transaction
* `GET /v1/transaction`: Get List User Transaction

### Swagger Endpoints

* `GET /swagger/*`: Swagger documentation

### 🔄 API Versioning

The API uses URL versioning (`/v1/`) to ensure backward compatibility. When making breaking changes, create a new version while maintaining the old one for a transition period.

## 🚀 Setup Instructions

### Prerequisites
- Go 1.24.2 or higher
- MySQL
- Make (for build automation)

### Installation

1. **Clone the repository**
```bash
git clone <repository-url>
cd loan-app
```

2. **Install dependencies**
```bash
go mod download
```

3. **Configure the application**

modify the configuration on `config/config.json`


4. **Set up the database**
```bash
# Run database migrations
make migrate-up
```

5. **Build and run the application**
```bash
# Development mode with hot reload
make run-dev

# Production mode
make run
```

### Available Make Commands

```bash
# Database migrations
make migrate-create name=<migration_name>  # Create new migration
make migrate-up                            # Apply migrations
make migrate-down                          # Rollback last migration
make migrate-clean                         # Rollback all migrations
make migrate-status                        # Show migration status

# Build and run
make build                                 # Build the application
make run-dev                               # Run with hot reload
make run                                   # Run in production mode
make clean                                 # Clean build artifacts
make rebuild                               # Force rebuild

# Swagger documentation
make swagger-gen                           # Generate Swagger documentation
make swagger-serve                         # Serve Swagger UI (requires server running)

# Help
make help                                  # Show all available commands
```

### Development Workflow

1. **Create a new feature branch**
```bash
git checkout -b feature/your-feature-name
```

2. **Make your changes and test**
```bash
make run-dev
```

3. **Create database migration if needed**
```bash
make migrate-create name=add_new_table
```

4. **Run tests and build**
```bash
make build
```

5. **Commit and push your changes**
```bash
git add .
git commit -m "feat(<module>): add new feature"
git push origin feature/your-feature-name
```