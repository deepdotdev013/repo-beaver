# 🦫 Repo Beaver

Generate production-ready backend project structures in seconds.

Repo Beaver is a CLI tool that helps you bootstrap clean, scalable backend architectures for multiple languages like **Go** and **Node.js** — with zero setup friction.

---

## ✨ Features

* 🧱 Generate production-level folder structures
* ⚡ Supports multiple backend languages (Go, Node.js)
* 🎯 Clean CLI experience with interactive prompts
* 🎨 Beautiful terminal UI powered by Bubble Tea
* 📦 Auto-initialization (`go mod init`, `npm init`)
* 📄 Pre-configured templates (README, .gitignore, boilerplate files)
* 🔒 Safe directory handling (overwrite protection)

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
repo-beaver
```

---

### Interactive Flow

1. Select backend language (Go / Node.js)
2. Enter project name
3. (Go only) Enter module path
4. Sit back while Repo Beaver generates your project 🚀

---

## 📁 Example Output

### Go Project Structure

```
my-go-app/
├── cmd/
│   └── my-go-app/
│       └── main.go
├── internal/
├── pkg/
├── go.mod
├── README.md
└── .gitignore
```

---

### Node.js Project Structure

```
my-node-app/
├── src/
│   ├── controllers/
│   ├── models/
│   ├── services/
│   ├── routes/
│   └── app.js
├── configs/
├── package.json
├── README.md
└── .gitignore
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
* Bubble Tea (TUI)
* Go Embed (template system)
* OS Exec (project initialization)

---

## 📌 Roadmap

* [ ] Add support for more languages (Python, FastAPI, Django)
* [ ] Add configuration options (DB, auth, etc.)
* [ ] Convert into full CLI with commands (Cobra)
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
