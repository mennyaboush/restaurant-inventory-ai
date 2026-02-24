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
### 2. Go Programming **[IN PROGRESS]**
- HTTP servers and REST APIs ✅
- Database interactions (PostgreSQL) ✅
- Error handling patterns ✅
- Project structure best practices ✅
- Testing (unit + integration) ⬜

### 3. Databases **[IN PROGRESS]**
- PostgreSQL fundamentals ✅
- Data modeling ✅
- Migrations ✅
- SQL queries ✅
- Transactions ⬜

### 4. Frontend Development **[NEXT]**
- HTML/JavaScript/CSS
- REST API consumption
- Mobile-responsive design
- Hebrew RTL layout

### 5. AI Engineering **[FUTURE]**
- LLM API integration (OpenAI, Ollama)
- RAG (Retrieval Augmented Generation)
- Natural language understanding
- Prompt engineering

### 6. Containers & Kubernetes **[FUTURE]**
- Docker and Dockerfiles
- Kubernetes deployments
- Production deployment

---

## 📁 Project Structure

```
restaurant-inventory-ai/
├── cmd/
│   ├── server/main.go          ← Main application entry
│   └── dev/main.go             ← Dev utilities
├── internal/
│   ├── api/api.go              ← HTTP handlers
│   ├── models/product.go       ← Data structures
│   ├── repository/             ← Database layer
│   │   ├── repository.go       ← Interface
│   │   ├── postgres_store.go   ← PostgreSQL implementation
│   │   └── memory_store.go     ← In-memory (testing)
│   └── service/service.go      ← Business logic (future)
├── config/
│   └── config.go               ← Configuration loading
├── migrations/
│   ├── 001_create_products_table.sql
│   └── 002_create_stock_tables.sql
├── docs/
│   ├── REQUIREMENTS.md         ← Product requirements
│   ├── ARCHITECTURE.md         ← System architecture
│   ├── DATA_MODELS_FINAL.md    ← Database models
│   └── DECISIONS.md            ← Design decisions
├── learn/
│   └── 01_basics/              ← Learning exercises
├── .env                        ← Environment config
├── docker-compose.yml          ← PostgreSQL setup
├── go.mod                      ← Go dependencies
└── README.md                   ← You are here
```

---

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Docker Desktop
- Git

### Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/restaurant-inventory-ai.git
cd restaurant-inventory-ai

# Start PostgreSQL
docker-compose up -d postgres

# Set environment variables
export $(cat .env | xargs)

# Run the server
go run cmd/server/main.go

# Test the API
curl http://localhost:8080/products
```

### Development Commands

```bash
# Run tests
go test ./...

# Run integration tests
go test ./internal/repository/ -tags=integration

# Check database
docker exec -it postgres-inventory psql -U postgres -d inventory

# View logs
docker-compose logs -f postgres
```

---

## 📚 Documentation

- [📋 Requirements](docs/REQUIREMENTS.md) - What we're building
- [🏗️ Architecture](docs/ARCHITECTURE.md) - How it's structured
- [💾 Data Models](docs/DATA_MODELS_FINAL.md) - Database schema
- [🎯 Decisions](docs/DECISIONS.md) - Design choices
- [📖 Learning Journey](LEARNING_JOURNEY.md) - Progress tracking
- [🎯 Work Plan](WORK_PLAN.md) - Weekly schedule

---

## 🛠️ Tech Stack (MVP)

| Layer | Technology | Status |
|-------|------------|--------|
| **Backend** | Go 1.21 | ✅ Working |
| **Database** | PostgreSQL 16 | ✅ Working |
| **Router** | Chi | ✅ Working |
| **Frontend** | HTML/JS | ⬜ Next |
| **Auth** | JWT | ⬜ Next |
| **Deploy** | VPS/K8s | ⬜ Future |

### Future Additions
- **AI/LLM:** Ollama/OpenAI for chat
- **Cache:** Redis for sessions
- **Vector DB:** pgvector for RAG
- **Mobile:** PWA or native app

---

## 📊 Current Status

### ✅ Completed (Week 1-3)
- [x] Go project setup
- [x] Data models (Product, Stock, StockMovement)
- [x] PostgreSQL database setup
- [x] Migrations
- [x] REST API (products CRUD)
- [x] Stock operations API
- [x] Repository pattern with interface

### 🔄 In Progress (Week 3-4)
- [ ] Movement history API
- [ ] Basic authentication
- [ ] Input validation
- [ ] Error handling improvements

### ⬜ Next Steps (Week 4-6)
- [ ] Web frontend UI
- [ ] Hebrew language support
- [ ] Low stock alerts
- [ ] User management
- [ ] Deploy to production

---

## 💡 Learning Approach

**Philosophy:** Learn by building something real

1. **Start Simple:** Basic CRUD API with PostgreSQL
2. **Iterate:** Add features one at a time
3. **Understand Why:** Every decision is documented
4. **Production Ready:** Not just a toy project
5. **Best Practices:** Industry-standard patterns

**Weekly Structure:**
- Monday-Friday: Learning + coding (2-3 hours/day)
- Weekend: Review, refactor, plan next week

---

## � MVP Deployment Cost

### Development (FREE)
Everything runs locally on your machine.

### Production MVP Options

| Option | Monthly Cost | What You Get |
|--------|--------------|--------------|
| **VPS (Recommended)** | $6-12 | DigitalOcean/Hetzner droplet, PostgreSQL, Nginx |
| **Platform-as-a-Service** | $0-10 | Railway.app or Render.com (free tier available) |
| **Self-Hosted** | ~$5 | Raspberry Pi at restaurant + domain |

### Future Costs (v2.0+)
- **AI (OpenAI API):** ~$10-30/month for natural language features
- **WhatsApp Business API:** ~$20/month + per-message fees
- **Kubernetes Cluster:** ~$20-50/month (DigitalOcean Kubernetes)

**MVP Strategy:** Start with free/cheap VPS, add features and costs gradually.

---

## 🎯 Roadmap

### Phase 1: MVP (Weeks 1-6) **[IN PROGRESS]**
- [x] Database setup
- [x] Product management API
- [x] Stock tracking API
- [ ] Web interface
- [ ] Basic authentication
- [ ] Deploy to production

### Phase 2: Enhanced Features (Weeks 7-10)
- [ ] Movement history
- [ ] Low stock alerts
- [ ] Multi-user support (Manager/Employee roles)
- [ ] Hebrew language UI
- [ ] PWA (works offline)

### Phase 3: AI Features (Weeks 11-14)
- [ ] Natural language input
- [ ] Smart suggestions
- [ ] Demand prediction
- [ ] Auto-ordering suggestions

### Phase 4: Advanced (Future)
- [ ] Supplier management
- [ ] Recipe costing
- [ ] POS integration
- [ ] Multiple locations

---

## 🤝 Contributing

This is a learning project with a real user (restaurant owner). 

**Suggestions Welcome:**
- Open an issue to discuss features
- Share your own learning journey
- Suggest improvements to documentation

---

## 📝 License

MIT License - Feel free to learn from and adapt this project.

---
