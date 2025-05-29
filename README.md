# Echo Auth Boilerplate

A Go boilerplate using Echo Framework with authentication and feature-based organization.

## 🚀 Features

- Echo Framework for REST API
- Feature-based organization
- Authentication system (login/signup)
- Custom logging
- Environment-based configuration
- PostgreSQL connection
- Health check endpoint

## 📋 Prerequisites

- Go 1.22.x
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

## 📁 Project Structure

```
.
├── cmd/
│   └── server.go         # Application entry point
├── internal/
│   ├── app/             # Application configuration
│   ├── config/          # Configurations
│   └── features/        # Application features
│       ├── auth/        # Authentication feature
│       └── health/      # Health check feature
├── pkg/
│   ├── database/        # Database configuration
│   └── logger/          # Custom logger
└── logs/                # Logs directory
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
