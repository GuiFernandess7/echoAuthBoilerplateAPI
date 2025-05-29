# <div align="center">Echo API Boilerplate</div>


<div align="center">

<img src="https://github.com/user-attachments/assets/0f2e5144-5ac3-46dd-882b-d1496cfc973b" width="400" alt="Go Logo">

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

</div>

<div align="center">

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-v4-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

</div>

A Go boilerplate using Echo Framework with authentication and feature-based organization.

## 1. Features

- 🚀 Echo Framework for REST API
- 📦 Feature-based organization
- 🔐 Authentication system (login/signup)
- 📝 Custom logging
- ⚙️ Environment-based configuration
- 🗄️ PostgreSQL connection
- 🏥 Health check endpoint

## 2. Prerequisites

- Go 1.22.x or higher
- PostgreSQL
- Make (optional, for using make commands)

## 🔧 Setup

1. Clone the repository:
```bash
git clone https://github.com/GuiFernandess7/echoAuthBoilerplate.git
cd echoAuthBoilerplate
```

2. Install dependencies:
```bash
go mod download
```

3. Configure environment variables:
```bash
cp .env.example .env
```

Edit the `.env` file with your settings:
```env
SERVER_PORT=8080
LOG_LEVEL=INFO

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
DB_SSLMode=disable
```

## 🏃‍♂️ Running the Application

```bash
go run cmd/server.go
```

Or using make:
```bash
make run
```

## Project Structure

```
.
├── cmd/                    # Application entry points
│   └── server.go          # Main server initialization and configuration - entrypoint
│
├── internal/              # Private application code
│   ├── app/              # Core application setup and configuration
│   │   └── app.go        # Application struct and initialization
│   │
│   ├── config/           # Configuration management
│   │   └── config.go     # Environment and application settings
│   │
│   └── features/         # Business logic organized by features
│       ├── auth/         # Authentication feature
│       │   ├── handler.go    # HTTP request handlers
│       │   ├── repository.go # Database operations
│       │   ├── router.go     # Route definitions
│       │   └── model/        # Data models
│       │
│       └── health/       # Health check feature
│           └── router.go # Health check endpoints
│
├── pkg/                  # Public libraries that can be used by external applications
│   ├── database/        # Database connection and configuration
│   │   └── database.go  # Database setup and utilities
│   │
│   └── logger/          # Custom logging implementation
│       └── logger.go    # Logging configuration and utilities
│
├── logs/                # Application logs directory
│   └── app.log         # Main application log file
│
├── .env.example        # Example environment variables
├── .gitignore         # Git ignore rules
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
├── Makefile          # Build automation commands
└── README.md         # Project documentation
```

## 🔍 Endpoints

### Health Check
- `GET /health` - Check application status

### Authentication
- `POST /auth/login` - User login
- `POST /auth/signup` - New user registration

## 🛠️ Development

### Available Make Commands

```bash
make run        # Run the application
make build      # Build the application
make test       # Run tests
make clean      # Clean compiled files
```

### Logs

Logs are saved in `logs/app.log` and include different levels:
- DEBUG
- INFO
- WARN
- ERROR
- FATAL

## 📝 License

This project is under the MIT license. See the [LICENSE](LICENSE) file for more details.
