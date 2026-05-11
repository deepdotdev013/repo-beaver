# 🦫 Repo Beaver

Generate production-ready backend project structures in seconds.

Repo Beaver is a CLI tool that helps you bootstrap clean, scalable backend architectures for multiple languages like **Go** and **Node.js** — with zero setup friction.

---

## ✨ Features

* 🧱 Generate production-ready backend project structures
* ⚡ Supports multiple backend languages (Go, Node.js) with frameworks (Gin, Gorilla, Express, Fastify) or vanilla setups
* 🎯 Clean CLI experience with interactive prompts or direct flags
* 🎨 Beautiful terminal UI powered by Bubble Tea
* 📦 Auto-initialization (`go mod init`, `npm init`)
* 📄 Pre-configured templates (README, .gitignore, Dockerfile, .env.example, GitHub Actions workflow)
* 🔒 Safe directory handling (overwrite protection)
* 🚀 Quick project creation with framework-specific flags

---

## 🚀 Installation

### Option 1: Install via Go

```bash
go install github.com/deepdotdev013/repo-beaver/cmd@latest
```

> Make sure `$GOPATH/bin` is in your PATH.

---

### Option 2: Clone and run locally

```bash
git clone https://github.com/deepdotdev013/repo-beaver.git
cd repo-beaver
go run cmd/main.go
```

---

## 🛠️ Usage

Run the CLI:

```bash
repo-beaver create [project-name]
```

### Interactive Flow

1. Select backend language (Go / Node.js)
2. Enter project name (if not provided as argument)
3. (Go only) Enter module path
4. Select framework (Gin, Gorilla, Express, Fastify, or None)
5. Sit back while Repo Beaver generates your project 🚀

### Quick Creation with Flags

For faster project setup, use framework-specific flags:

```bash
# Express.js project
repo-beaver create my-api --express

# Fastify project
repo-beaver create my-api --fastify

# Gin project
repo-beaver create my-api --gin

# Gorilla Mux project
repo-beaver create my-api --gorilla
```

### Other Commands

```bash
# Show version
repo-beaver version

# Show help
repo-beaver --help
```

---

## 📁 Example Output

### Go Project Structure (with Gin/Gorilla)

```
my-go-app/
├── cmd/
│   └── my-go-app/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── stores/
│   ├── routes/
│   ├── clients/
│   ├── services/
│   ├── repositories/
│   ├── models/
│   ├── domains/
│   └── core/
├── pkg/
│   ├── logger/
│   ├── config/
│   ├── middleware/
│   ├── security/
│   └── utils/
├── configs/
├── infra/
│   └── db/
├── tests/
├── .github/
│   └── workflows/
│       └── ci.yml
├── Dockerfile
├── .env.example
├── .gitignore
├── go.mod
└── README.md
```

### Node.js Project Structure (with Express)

```
my-node-app/
├── src/
│   ├── controllers/
│   ├── models/
│   ├── services/
│   ├── repositories/
│   ├── routes/
│   ├── middlewares/
│   ├── utils/
│   ├── policies/
│   └── validators/
├── configs/
├── tests/
├── .github/
│   └── workflows/
│       └── ci.yml
├── app.js
├── Dockerfile
├── .env.example
├── .gitignore
├── package.json
└── README.md
```

---

## 🧠 Why Repo Beaver?

Setting up a backend project from scratch often involves:

* Creating folders manually
* Writing boilerplate files
* Configuring project structure
* Initializing dependencies

Repo Beaver automates all of this so you can:

> **Focus on building, not setting up.**

---

## 🧩 Tech Stack

* Go (CLI + generators)
* Cobra (Command-line interface)
* Bubble Tea (TUI)
* Go Embed (template system)
* OS Exec (project initialization)

---

## 📌 Roadmap

* [x] Support for "none" framework option
* [ ] Add support for more languages (Python, FastAPI, Django)
* [ ] Add configuration options (DB, auth, etc.)
* [ ] Convert into full CLI with commands (Cobra) ✅
* [ ] AI-assisted scaffolding (future vision 👀)

---

## 🤝 Contributing

Contributions are welcome!

1. Fork the repo
2. Create a feature branch
3. Submit a PR

---

## 📄 License

MIT License

---

## 💡 Final Note

> Let’s build something meaningful. Happy coding! 🦫
