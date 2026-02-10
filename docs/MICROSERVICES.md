# 🏗️ Microservices Architecture

## System Overview

```mermaid
flowchart TB
    subgraph Clients["📱 Clients"]
        Mobile["Mobile/Tablet<br/>(PWA)"]
        Web["Web Browser"]
        WhatsApp["WhatsApp<br/>(Future)"]
    end

    subgraph Gateway["🚪 API Gateway (Go)"]
        Kong["Kong/Traefik<br/>- Rate Limiting<br/>- SSL Termination<br/>- Routing"]
    end

    subgraph Auth["🔐 Auth Service (Go)"]
        AuthAPI["Auth API<br/>- Login/Register<br/>- JWT Tokens<br/>- Permissions"]
        AuthDB[(Auth DB<br/>PostgreSQL)]
    end

    subgraph Inventory["📦 Inventory Service (Go)"]
        InvAPI["Inventory API<br/>- Products CRUD<br/>- Stock Tracking<br/>- Movements"]
        InvDB[(Inventory DB<br/>PostgreSQL)]
        InvCache[(Redis Cache)]
    end

    subgraph AI["🤖 AI Service (Go/Python)"]
        AIAPI["AI API<br/>- Chat Processing<br/>- Intent Detection<br/>- Clarification"]
        RAG["RAG Engine<br/>- Context Retrieval<br/>- Product Matching"]
        LLM["LLM Provider<br/>Ollama/OpenAI"]
        VectorDB[(Vector DB<br/>pgvector)]
    end

    subgraph Notification["🔔 Notification Service (Go)"]
        NotifAPI["Notification API<br/>- Push Notifications<br/>- WhatsApp<br/>- Email"]
        NotifQueue[(Message Queue<br/>Redis/RabbitMQ)]
    end

    subgraph Analytics["📊 Analytics Service (Go)"]
        AnalyticsAPI["Analytics API<br/>- Usage Patterns<br/>- Waste Reports<br/>- Predictions"]
        AnalyticsDB[(Analytics DB<br/>TimescaleDB)]
    end

    Mobile --> Kong
    Web --> Kong
    WhatsApp --> Kong
    
    Kong --> AuthAPI
    Kong --> InvAPI
    Kong --> AIAPI
    Kong --> NotifAPI
    Kong --> AnalyticsAPI
    
    AuthAPI --> AuthDB
    InvAPI --> InvDB
    InvAPI --> InvCache
    AIAPI --> RAG
    RAG --> LLM
    RAG --> VectorDB
    RAG --> InvAPI
    NotifAPI --> NotifQueue
    AnalyticsAPI --> AnalyticsDB
    
    InvAPI -.->|"Events"| NotifAPI
    InvAPI -.->|"Events"| AnalyticsAPI
```

---

## Kubernetes Architecture

```mermaid
flowchart TB
    subgraph K8s["☸️ Kubernetes Cluster"]
        subgraph Ingress["Ingress Layer"]
            Nginx["Nginx Ingress<br/>Controller"]
            Cert["Cert Manager<br/>(SSL)"]
        end
        
        subgraph Services["Services Layer"]
            subgraph AuthNS["Namespace: auth"]
                AuthDeploy["Auth Service<br/>Deployment<br/>replicas: 2"]
                AuthSvc["Auth Service<br/>(ClusterIP)"]
                AuthPVC["PVC: auth-db"]
            end
            
            subgraph InvNS["Namespace: inventory"]
                InvDeploy["Inventory Service<br/>Deployment<br/>replicas: 2-5 (HPA)"]
                InvSvc["Inventory Service<br/>(ClusterIP)"]
                InvPVC["PVC: inventory-db"]
            end
            
            subgraph AINS["Namespace: ai"]
                AIDeploy["AI Service<br/>Deployment<br/>replicas: 2"]
                AISvc["AI Service<br/>(ClusterIP)"]
            end
            
            subgraph NotifNS["Namespace: notification"]
                NotifDeploy["Notification Service<br/>Deployment<br/>replicas: 1"]
                NotifSvc["Notification Service<br/>(ClusterIP)"]
            end
        end
        
        subgraph Data["Data Layer"]
            subgraph DBs["Databases"]
                PG["PostgreSQL<br/>StatefulSet"]
                Redis["Redis<br/>StatefulSet"]
                Ollama["Ollama<br/>Deployment<br/>(GPU node)"]
            end
        end
        
        subgraph Config["Configuration"]
            CM["ConfigMaps<br/>- App configs<br/>- Feature flags"]
            Secrets["Secrets<br/>- DB passwords<br/>- API keys<br/>- JWT secret"]
        end
    end
    
    Internet((Internet)) --> Nginx
    Nginx --> AuthSvc
    Nginx --> InvSvc
    Nginx --> AISvc
    Nginx --> NotifSvc
    
    AuthDeploy --> PG
    InvDeploy --> PG
    InvDeploy --> Redis
    AIDeploy --> Ollama
    AIDeploy --> PG
```

---

## Service Details

### 1. Auth Service
```yaml
Responsibility: User authentication and authorization
Port: 8081
Endpoints:
  POST /auth/login
  POST /auth/register
  POST /auth/refresh
  GET  /auth/me
  POST /auth/logout
  
Database Tables:
  - users (id, email, password_hash, role, employee_id)
  - refresh_tokens (id, user_id, token, expires_at)
  
K8s Resources:
  - Deployment (2 replicas)
  - Service (ClusterIP)
  - ConfigMap (jwt settings)
  - Secret (jwt secret, db password)
```

### 2. Inventory Service
```yaml
Responsibility: Products, stock, movements
Port: 8082
Endpoints:
  # Categories
  GET    /categories
  POST   /categories
  PUT    /categories/:id
  DELETE /categories/:id
  
  # Products
  GET    /products
  GET    /products/:id
  POST   /products
  PUT    /products/:id
  DELETE /products/:id
  GET    /products/search?q=cola
  GET    /products/low-stock
  GET    /products/expiring
  
  # Stock
  GET    /stock
  GET    /stock/:product_id
  POST   /stock/movement  (IN/OUT/WASTE/ADJUST)
  GET    /stock/movements?product_id=&from=&to=
  POST   /stock/sync      (bulk update)
  
Database Tables:
  - categories
  - products
  - stock
  - stock_movements
  - suppliers (future)
  
K8s Resources:
  - Deployment (2-5 replicas, HPA)
  - Service (ClusterIP)
  - HorizontalPodAutoscaler
  - PersistentVolumeClaim (for DB)
```

### 3. AI Service
```yaml
Responsibility: Natural language processing, chat
Port: 8083
Endpoints:
  POST /chat/message
  GET  /chat/history
  POST /chat/voice (audio input)
  
Internal Calls:
  → Inventory Service (get stock, products)
  → Notification Service (send alerts)
  
Components:
  - Intent Classifier
  - Entity Extractor (product names)
  - Clarification Engine
  - RAG for context
  
K8s Resources:
  - Deployment (2 replicas)
  - Service (ClusterIP)
  - ConfigMap (LLM settings)
  - Secret (OpenAI API key)
```

### 4. Notification Service
```yaml
Responsibility: Alerts and notifications
Port: 8084
Endpoints:
  POST /notify/push
  POST /notify/whatsapp
  POST /notify/email
  GET  /notify/preferences/:user_id
  PUT  /notify/preferences/:user_id
  
Event Listeners:
  - Low stock event → Send alert
  - Expiring soon event → Send alert
  - Large movement event → Send alert
  
K8s Resources:
  - Deployment (1-2 replicas)
  - Service (ClusterIP)
  - Secret (WhatsApp API, email creds)
```

### 5. Analytics Service (Future)
```yaml
Responsibility: Reports and predictions
Port: 8085
Endpoints:
  GET /analytics/usage?product_id=&period=
  GET /analytics/waste?period=
  GET /analytics/predictions
  GET /analytics/report/daily
  GET /analytics/report/weekly
```

---

## Inter-Service Communication

```mermaid
sequenceDiagram
    participant C as Client
    participant G as API Gateway
    participant Auth as Auth Service
    participant Inv as Inventory Service
    participant AI as AI Service
    participant Notif as Notification Service

    C->>G: POST /chat/message<br/>"לקחתי 3 קולה קטן"
    G->>Auth: Validate JWT
    Auth-->>G: Valid (user: manager)
    G->>AI: Process message
    
    AI->>Inv: GET /products/search?q=קולה קטן
    Inv-->>AI: [Cola Small Can, Cola Small Plastic, Cola Small Glass]
    
    AI-->>G: Need clarification
    G-->>C: "איזה סוג? פחית / פלסטיק / זכוכית"
    
    C->>G: "פחית"
    G->>AI: Continue conversation
    
    AI->>Inv: GET /stock/product_id_123
    Inv-->>AI: {boxes: 5, units: 12}
    
    AI-->>G: "כמה ארגזים או בקבוקים?"
    G-->>C: "כמה ארגזים או בקבוקים?"
    
    C->>G: "3 בקבוקים"
    G->>AI: "3 בקבוקים"
    
    AI->>Inv: POST /stock/movement<br/>{product_id, type: OUT, units: -3}
    Inv-->>AI: Success
    
    Inv->>Notif: Event: Stock changed
    Note over Notif: Check if low stock → Alert if needed
    
    AI-->>G: "עדכנתי: -3 פחיות קולה קטן"
    G-->>C: Response
```

---

## Product Data Model (Revised)

```mermaid
erDiagram
    CATEGORY {
        uuid id PK
        string name "משקאות"
        string name_en "Drinks"
        int sort_order
    }
    
    PRODUCT {
        uuid id PK
        uuid category_id FK
        string name "קוקה קולה 330 פחית"
        string name_en "Coca Cola 330ml Can"
        string brand "Coca Cola"
        string size "330ml"
        string container_type "can/plastic/glass/bag/box"
        string unit_type "bottle/kg/piece/unit"
        int box_size "24"
        int default_expiry_days "365"
        int min_stock_level "48"
        decimal price_estimate "nullable"
        boolean is_active "true"
        timestamp created_at
        timestamp updated_at
    }
    
    STOCK {
        uuid id PK
        uuid product_id FK
        int quantity_boxes
        int quantity_units
        date expiry_date "nullable"
        timestamp updated_at
    }
    
    STOCK_MOVEMENT {
        uuid id PK
        uuid product_id FK
        uuid user_id FK
        string type "IN/OUT/WASTE/ADJUSTMENT"
        int boxes_change
        int units_change
        string reason "nullable"
        string notes "nullable"
        timestamp created_at
    }
    
    CATEGORY ||--o{ PRODUCT : contains
    PRODUCT ||--o{ STOCK : has
    PRODUCT ||--o{ STOCK_MOVEMENT : tracks
```

---

## AI Clarification Flow

```mermaid
flowchart TD
    Start["User: לקחתי 3 קולה קטן"] --> Parse["Parse Intent<br/>Action: OUT<br/>Quantity: 3<br/>Product: קולה קטן"]
    
    Parse --> Search["Search Products<br/>WHERE name LIKE '%קולה%'<br/>AND name LIKE '%קטן%'"]
    
    Search --> Found{"How many<br/>matches?"}
    
    Found -->|"0"| NotFound["לא מצאתי מוצר כזה.<br/>התכוונת ל...?<br/>(suggest similar)"]
    
    Found -->|"1"| OneMatch["Check stock of match"]
    
    Found -->|"2+"| MultiMatch["Multiple matches found"]
    
    MultiMatch --> AskType["איזה סוג?<br/>🥫 פחית<br/>🧴 פלסטיק<br/>🍾 זכוכית"]
    
    AskType --> UserSelect["User selects"]
    
    UserSelect --> OneMatch
    
    OneMatch --> AskUnit{"Box size > 0?"}
    
    AskUnit -->|"Yes"| AskBoxOrUnit["3 ארגזים או 3 בודדים?"]
    AskUnit -->|"No (kg/piece)"| Confirm
    
    AskBoxOrUnit --> UserAnswers["User: בודדים"]
    
    UserAnswers --> Validate{"Stock >= 3?"}
    
    Validate -->|"Yes"| Confirm["מעדכן: -3 פחיות קולה קטן"]
    Validate -->|"No"| Warn["יש רק 2 במלאי.<br/>לעדכן ל-0?"]
    
    Confirm --> Update["POST /stock/movement"]
    Warn --> UserConfirm["User confirms"]
    UserConfirm --> Update
    
    Update --> CheckLow{"Stock < min_level?"}
    CheckLow -->|"Yes"| Alert["🔔 Trigger low stock alert"]
    CheckLow -->|"No"| Done["✅ Done"]
    Alert --> Done
```

---

## Directory Structure (Microservices)

```
restaurant-inventory-ai/
├── services/
│   ├── auth/
│   │   ├── cmd/
│   │   │   └── server/
│   │   │       └── main.go
│   │   ├── internal/
│   │   │   ├── api/
│   │   │   ├── models/
│   │   │   ├── repository/
│   │   │   └── service/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   └── go.sum
│   │
│   ├── inventory/
│   │   ├── cmd/
│   │   │   └── server/
│   │   │       └── main.go
│   │   ├── internal/
│   │   │   ├── api/
│   │   │   ├── models/
│   │   │   ├── repository/
│   │   │   └── service/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   └── go.sum
│   │
│   ├── ai/
│   │   ├── cmd/
│   │   │   └── server/
│   │   │       └── main.go
│   │   ├── internal/
│   │   │   ├── api/
│   │   │   ├── chat/
│   │   │   ├── intent/
│   │   │   ├── rag/
│   │   │   └── llm/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   └── go.sum
│   │
│   └── notification/
│       ├── cmd/
│       │   └── server/
│       │       └── main.go
│       ├── internal/
│       │   ├── api/
│       │   ├── whatsapp/
│       │   ├── push/
│       │   └── email/
│       ├── Dockerfile
│       ├── go.mod
│       └── go.sum
│
├── shared/
│   └── pkg/
│       ├── auth/        (JWT validation)
│       ├── errors/      (common errors)
│       ├── logger/      (structured logging)
│       └── middleware/  (common middleware)
│
├── deployments/
│   └── k8s/
│       ├── base/
│       │   ├── namespace.yaml
│       │   ├── auth/
│       │   ├── inventory/
│       │   ├── ai/
│       │   └── notification/
│       ├── overlays/
│       │   ├── dev/
│       │   └── prod/
│       └── kustomization.yaml
│
├── docker-compose.yaml  (local development)
├── Makefile
└── docs/
    ├── ARCHITECTURE.md
    ├── REQUIREMENTS.md
    └── DECISIONS.md
```

---

## Development Plan (Updated)

### Phase 1: Foundation (Weeks 1-3)
Learn Go + Build Inventory Service core
- Go fundamentals
- Product/Category/Stock models
- Basic CRUD APIs
- PostgreSQL setup
- Docker container

### Phase 2: Auth Service (Week 4)
- User model
- JWT authentication
- Role-based access
- Service-to-service auth

### Phase 3: Docker & Local Dev (Week 5)
- Dockerize both services
- docker-compose for local
- Service communication
- Shared database vs separate DBs

### Phase 4: Kubernetes Basics (Weeks 6-7)
- Deploy to local K8s (minikube/kind)
- Deployments, Services, ConfigMaps
- Ingress setup
- Basic monitoring

### Phase 5: AI Service (Weeks 8-10)
- Chat endpoint
- Ollama integration
- Intent detection
- Product search & matching
- Clarification flow

### Phase 6: Notification Service (Week 11)
- Alert triggers
- WhatsApp integration
- Event-driven architecture

### Phase 7: Production Ready (Week 12+)
- HPA (auto-scaling)
- Health checks
- Logging & monitoring
- CI/CD pipeline
