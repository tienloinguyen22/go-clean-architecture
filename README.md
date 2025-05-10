# Go Clean Architecture Project

A robust, modular Go application implementing clean architecture principles to ensure separation of concerns, maintainability, and scalability.

## Project Overview

This project demonstrates a clean architecture approach in Go, providing a structured and extensible application design. By separating concerns and defining clear boundaries between layers, the application becomes easier to develop, test, and maintain.

## Project Structure

```
go-clean-architecture/
│
├── internal/               # Main source code directory
│   ├── domain/             # Core business logic and entities
│   ├── infrastructure/     # External interfaces and implementations
│   ├── interfaces/         # Application interfaces and controllers
│   └── main.go             # Application entry point
│
├── pkg/                    # Shared packages and libraries
│   ├── logger/             # Logging utilities
│   ├── config/             # Configuration management
│   └── utils/              # Common utility functions
│
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

### Directory Explanation

#### `/internal/`
Contains the core application logic that is specific to this project and not intended to be imported by other projects.

#### `/pkg/`
Contains shared packages and libraries that can be imported by other projects. This directory follows Go best practices for creating reusable code:
- `logger/`: Centralized logging utilities with consistent logging format and configuration
- `config/`: Configuration management tools for handling environment-specific settings
- `utils/`: Common utility functions that can be used across different parts of the application

The separation between `internal/` and `pkg/` helps maintain a clear boundary between project-specific and potentially reusable code.

## Prerequisites

- Go (version 1.16 or higher)
- Git

## Getting Started

### Prerequisites
- Docker
- Docker Compose
- Make

### Local Development with Dependencies

1. Clone the repository
   ```bash
   git clone https://github.com/tienloinguyen22/go-clean-architecture.git
   cd go-clean-architecture
   ```

2. Start Infrastructure Dependencies
   ```bash
   # Starts PostgreSQL and Redis in the background
   docker-compose up -d postgres redis
   ```

3. Download Go dependencies
   ```bash
   go mod tidy
   ```

4. Run the application
   ```bash
   make start
   ```

## Development Principles

### Clean Architecture

This project follows clean architecture principles:

- **Domain Layer**: Contains core business logic and entities
- **Infrastructure Layer**: Implements external interfaces and dependencies
- **Interfaces Layer**: Defines application controllers and use cases
- **Dependency Rule**: Inner layers are independent of outer layers

### Key Benefits

- Highly testable
- Independent of frameworks
- Independent of UI
- Independent of database
- Independent of external interfaces

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Tien Loi Nguyen - [tienloinguyen22@gmail.com](mailto:tienloinguyen22@gmail.com)

Project Link: [https://github.com/tienloinguyen22/go-clean-architecture](https://github.com/tienloinguyen22/go-clean-architecture)
