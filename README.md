# 🍕 Restaurant Inventory AI - Smart Stock Management

> An AI-powered inventory and logistics management system for restaurants

[![Learning Project](https://img.shields.io/badge/Purpose-Learning%20%2B%20Production-blue)]()
[![Go](https://img.shields.io/badge/Backend-Go-00ADD8)]()
[![Kubernetes](https://img.shields.io/badge/Deploy-Kubernetes-326CE5)]()
[![AI](https://img.shields.io/badge/AI-RAG%20%2B%20LLM-green)]()

---

## 📖 Project Story - How We Got Here

### The Journey

This project started as an interview preparation exercise. While planning hands-on practice for technical interviews (Go, Kubernetes, Linux, Docker), we realized:

> "Why practice with fake projects when we can build something real?"

The idea evolved through several iterations:

```
Initial Ideas:
├── ❌ Generic Todo App (too simple, not exciting)
├── ❌ AI Security Monitor (good, but no real user)
├── ❌ Stock Market Analyzer (interesting, but no real user)
│
└── ✅ Restaurant Inventory System
    • Real user: Father owns a restaurant!
    • Real problem: Stock management is painful
    • Real value: Reduce waste = save money
    • AI makes sense: Demand prediction is genuinely useful
```

### Why This Project is Perfect for Learning

| Reason | Explanation |
|--------|-------------|
| **Real User** | My father will actually use this |
| **Real Feedback** | We'll know immediately if something doesn't work |
| **Full Stack** | Frontend, Backend, Database, AI, Infrastructure |
| **AI Integration** | RAG, LLMs, embeddings - all industry-relevant |
| **Production Ready** | Must be reliable (restaurant can't have downtime!) |
| **Portfolio Gold** | "I built a production system for a real business" |

---

## 🎯 What We're Building

### The Problem
Restaurant owners struggle with:
- 📦 Tracking what's in stock across multiple storage areas
- ⏰ Items expiring before being used (waste = money lost)
- 📉 Running out of ingredients during busy hours
- 💰 Not knowing true ingredient costs
- 📊 Ordering too much or too little

### The Solution

```
┌─────────────────────────────────────────────────────────────────┐
│              Restaurant Inventory AI                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  📦 Track Stock → 🤖 AI Analysis → 📱 Smart Alerts & Orders     │
│                                                                 │
│  "Know what to order, when, and how much - automatically"       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Core Features

| Feature | Description | AI-Powered? |
|---------|-------------|-------------|
| **Inventory Tracking** | Track ingredients, quantities, locations | ❌ |
| **Expiration Alerts** | Warn before items expire | ❌ |
| **Low Stock Alerts** | Notify when reordering needed | ❌ |
| **Demand Prediction** | "You'll need 20kg tomatoes Friday" | ✅ |
| **Smart Order Suggestions** | Recommend what/when to order | ✅ |
| **Recipe Cost Calculator** | Know cost per dish | ❌ |
| **Supplier Comparison** | Find better prices | ✅ |
| **Waste Analysis** | Identify waste patterns | ✅ |
| **Natural Language Queries** | "What should I order?" | ✅ RAG |

---

## 🧠 Learning Goals

This project is designed to teach:

### 1. Go Programming
- HTTP servers and REST APIs
- Database interactions
- Concurrent programming
- Error handling patterns
- Project structure best practices

### 2. AI Engineering
- LLM API integration (OpenAI, Ollama)
- RAG (Retrieval Augmented Generation)
- Embeddings and vector search
- Prompt engineering
- AI application patterns

### 3. Databases
- PostgreSQL fundamentals
- Data modeling
- Migrations
- Vector databases (Qdrant)

### 4. Containers & Kubernetes
- Docker and Dockerfiles
- Multi-stage builds
- Kubernetes deployments
- Services, ConfigMaps, Secrets
- Persistent storage

### 5. Linux & DevOps
- Command line proficiency
- Shell scripting
- Logging and monitoring
- Debugging techniques
- CI/CD basics

---

## 📁 Project Structure

```
restaurant-inventory-ai/
│
├── README.md                    ← You are here
├── PROJECT_STORY.md             ← Why and how we built this
├── LEARNING_JOURNEY.md          ← Track learning progress
│
├── docs/                        ← Documentation
│   ├── 01_ARCHITECTURE.md       ← System design
│   ├── 02_API_DESIGN.md         ← API specifications
│   ├── 03_DATABASE_SCHEMA.md    ← Data models
│   ├── 04_AI_FEATURES.md        ← AI/RAG design
│   └── 05_DEPLOYMENT.md         ← Production guide
│
├── learning/                    ← Learning materials
│   ├── 01_go_basics/            ← Go tutorials
│   ├── 02_api_development/      ← REST API lessons
│   ├── 03_databases/            ← SQL & PostgreSQL
│   ├── 04_docker/               ← Container lessons
│   ├── 05_ai_engineering/       ← AI/ML/RAG lessons
│   └── 06_kubernetes/           ← K8s lessons
│
├── backend/                     ← Go API (will build)
├── frontend/                    ← Web UI (will build)
├── ai-service/                  ← AI/RAG service (will build)
├── kubernetes/                  ← K8s manifests (will build)
└── docker-compose.yaml          ← Local development
```

---

## 🚀 Quick Links

- [📋 Learning Journey](LEARNING_JOURNEY.md) - Track your progress
- [🏗️ Architecture](docs/01_ARCHITECTURE.md) - System design
- [🔌 API Design](docs/02_API_DESIGN.md) - Endpoints
- [💾 Database Schema](docs/03_DATABASE_SCHEMA.md) - Data models
- [🤖 AI Features](docs/04_AI_FEATURES.md) - RAG & predictions

---

## 🛠️ Tech Stack

| Layer | Technology | Why |
|-------|------------|-----|
| **Backend** | Go | Fast, simple, great for APIs |
| **Database** | PostgreSQL | Reliable, feature-rich |
| **Vector DB** | Qdrant | Fast vector search for RAG |
| **AI/LLM** | Ollama + OpenAI | Local dev + production |
| **Frontend** | HTML/JS (HTMX) | Simple, fast to learn |
| **Containers** | Docker | Industry standard |
| **Orchestration** | Kubernetes | Production-grade deployment |
| **Cache** | Redis | Fast data access |

---

## 💰 Cost Estimates

### Development (FREE)
Everything runs locally on your machine.

### Production Options

| Option | Monthly Cost | Notes |
|--------|--------------|-------|
| **Budget** | ~$30-50 | VPS + managed DB + OpenAI |
| **Self-Hosted AI** | ~$20 + hardware | Mini PC at restaurant |
| **Full Cloud** | ~$80-120 | Managed K8s + all services |

---

## 👨‍💻 Author

Built as a learning project with a real-world purpose.

---

*Let's build something amazing! 🚀*
