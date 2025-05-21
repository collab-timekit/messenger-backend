# 📬 Messenger Microservice

This project is a microservice for handling conversations, messages, and conversation members. It follows the principles of **Clean Architecture** and is built in **Go** using the **Gin** web framework.

---

## 📁 Project Structure

```
.
├── cmd/                    # Application entry point
│   └── main.go
├── config/                 # Static configuration (YAML)
├── internal/
│   ├── app/                # Application logic (use cases)
│   │   ├── port/
│   │   │   ├── in/         # Use case interfaces
│   │   │   └── out/        # Repository interfaces
│   │   └── service/        # Use case implementations
│   ├── bootstrap/          # Dependency injection and server setup
│   ├── domain/             # Domain entities
│   └── infra/              # Infrastructure
│       ├── adapter/
│       │   ├── in/
│       │   │   ├── rest/   # HTTP handlers (Gin)
│       │   │   └── ws/     # WebSocket support
│       │   └── out/
│       │       ├── keycloak/       # Auth provider integration
│       │       └── persistence/    # Repositories and models
│       ├── config/        # Configuration loaders
│       └── security/      # Authentication & middleware
├── .github/workflows/      # GitHub Actions CI/CD
├── .vscode/                # VS Code settings
├── Dockerfile              # Container specification
├── Makefile                # Build and test automation
├── go.mod / go.sum         # Go dependencies
```

---

## 🚀 Features

* JWT-based authentication (via Keycloak)
* REST API for message, conversation, and member handling
* WebSocket support for real-time updates
* Clean Architecture separation of concerns
* GitHub Actions CI/CD workflows

---

## 🛠️ Development

### Prerequisites

* Go 1.21+
* Docker (optional for DB and containerized runs)
* Make (or use Makefile targets manually)

### Running Locally

```bash
git clone https://github.com/your-org/messenger.git
cd messenger
make run
```

Or using Go:

```bash
go run ./cmd/main.go
```

### Configuration

Edit `config/config.yaml` for environment-specific settings like database or Keycloak.

### Running Tests

```bash
make test
```

---

## 📦 CI/CD

CI/CD is configured using GitHub Actions:

* `ci.yml`: Build, test and SonarQube analysis
* `release.yml`: Tag-based Docker builds
* `release-please.yml`: Automated changelogs & version bumps

---

## 📚 License

MIT License. See `LICENSE` file.

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

```bash
git checkout -b feature/my-feature
make test
```

Then open a Pull Request.

---

## 📞 Contact

For questions or feedback, contact the maintainers via your internal SmartWorkplace channels.