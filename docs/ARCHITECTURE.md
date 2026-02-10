# 🏗️ System Architecture

## High-Level Architecture

```mermaid
flowchart TB
    subgraph Clients["📱 Clients"]
        Mobile["Mobile App<br/>(iPhone/iPad)"]
        Web["Web Browser"]
    end

    subgraph Gateway["🚪 API Gateway"]
        Auth["Authentication<br/>& Authorization"]
        Router["Request Router"]
    end

    subgraph AI["🤖 AI Layer"]
        Chat["Chat Interface"]
        RAG["RAG Engine<br/>(Context-Aware AI)"]
        LLM["LLM Provider<br/>(Ollama/OpenAI)"]
    end

    subgraph Services["⚙️ Microservices"]
        InvService["📦 Inventory<br/>Service"]
        WorkService["👥 Workforce<br/>Service"]
        AlertService["🔔 Alert<br/>Service"]
        ReportService["📊 Report<br/>Service"]
    end

    subgraph Data["💾 Data Layer"]
        PostgreSQL[(PostgreSQL<br/>Main DB)]
        Redis[(Redis<br/>Cache)]
        VectorDB[(Vector DB<br/>Embeddings)]
    end

    subgraph MCP["🔌 MCP Servers (Future)"]
        MCPInventory["Inventory MCP"]
        MCPWorkforce["Workforce MCP"]
        MCPCalendar["Calendar MCP"]
    end

    Mobile --> Gateway
    Web --> Gateway
    Gateway --> Auth
    Auth --> Router
    Router --> AI
    Router --> Services
    AI --> RAG
    RAG --> LLM
    RAG --> VectorDB
    Chat --> RAG
    InvService --> PostgreSQL
    WorkService --> PostgreSQL
    AlertService --> Redis
    Services --> Data
    AI -.-> MCP
```

## User Flow - Main Interactions

```mermaid
flowchart LR
    subgraph User["👤 User Actions"]
        A["Open App"]
        B["Chat or Browse"]
    end

    subgraph ChatFlow["💬 Chat Flow"]
        C1["Ask Question"]
        C2["AI Understands Intent"]
        C3["AI Executes Action"]
        C4["Shows Result"]
    end

    subgraph VisualFlow["👁️ Visual Flow"]
        V1["Open Dashboard"]
        V2["See Stock/Schedule"]
        V3["Click to Edit"]
        V4["Changes Saved"]
    end

    A --> B
    B -->|"Chat"| C1 --> C2 --> C3 --> C4
    B -->|"Browse"| V1 --> V2 --> V3 --> V4
```

## Inventory Data Model

```mermaid
erDiagram
    PRODUCT {
        uuid id PK
        string name
        string category
        int box_size "items per box"
        float min_stock_level
        int expiry_days
    }
    
    STOCK {
        uuid id PK
        uuid product_id FK
        int quantity_boxes
        int quantity_items
        date expiry_date
        string location
        timestamp updated_at
    }
    
    SUPPLIER {
        uuid id PK
        string name
        string phone
        string email
    }
    
    PRODUCT_SUPPLIER {
        uuid product_id FK
        uuid supplier_id FK
        float price_per_box
        int min_order_qty
    }
    
    STOCK_MOVEMENT {
        uuid id PK
        uuid product_id FK
        string movement_type "IN/OUT/ADJUSTMENT"
        int boxes_change
        int items_change
        uuid user_id FK
        timestamp created_at
        string notes
    }

    PRODUCT ||--o{ STOCK : has
    PRODUCT ||--o{ PRODUCT_SUPPLIER : "supplied by"
    SUPPLIER ||--o{ PRODUCT_SUPPLIER : supplies
    PRODUCT ||--o{ STOCK_MOVEMENT : tracks
```

## Workforce Data Model

```mermaid
erDiagram
    EMPLOYEE {
        uuid id PK
        string name
        string phone
        string email
        boolean is_active
    }
    
    SKILL {
        uuid id PK
        string name "driver/pizza/falafel/manager/cashier"
    }
    
    EMPLOYEE_SKILL {
        uuid employee_id FK
        uuid skill_id FK
        int proficiency_level "1-5"
    }
    
    AVAILABILITY {
        uuid id PK
        uuid employee_id FK
        int day_of_week "0-6"
        time start_time
        time end_time
        date valid_from
        date valid_until
    }
    
    UNAVAILABILITY {
        uuid id PK
        uuid employee_id FK
        date date
        string reason
    }
    
    SHIFT {
        uuid id PK
        date date
        time start_time
        time end_time
        uuid required_skill FK
        int min_employees
    }
    
    SHIFT_ASSIGNMENT {
        uuid id PK
        uuid shift_id FK
        uuid employee_id FK
        string status "scheduled/confirmed/completed/absent"
    }
    
    HOLIDAY {
        uuid id PK
        date date
        string name
        boolean is_busy "more staff needed?"
    }

    EMPLOYEE ||--o{ EMPLOYEE_SKILL : has
    SKILL ||--o{ EMPLOYEE_SKILL : "assigned to"
    EMPLOYEE ||--o{ AVAILABILITY : defines
    EMPLOYEE ||--o{ UNAVAILABILITY : reports
    SHIFT ||--o{ SHIFT_ASSIGNMENT : contains
    EMPLOYEE ||--o{ SHIFT_ASSIGNMENT : "assigned to"
    SKILL ||--o{ SHIFT : requires
```

## User & Access Control

```mermaid
erDiagram
    USER {
        uuid id PK
        string email
        string password_hash
        string role "owner/manager/employee"
        uuid employee_id FK "nullable"
        boolean is_active
    }
    
    PERMISSION {
        uuid id PK
        string name "inventory.read/inventory.write/schedule.manage"
    }
    
    ROLE_PERMISSION {
        string role
        uuid permission_id FK
    }
    
    USER ||--o| EMPLOYEE : "linked to"
    PERMISSION ||--o{ ROLE_PERMISSION : "granted via"
```

## Chat/AI Flow

```mermaid
sequenceDiagram
    participant U as 👤 User
    participant C as 💬 Chat UI
    participant AI as 🤖 AI Engine
    participant RAG as 📚 RAG
    participant API as ⚙️ Services
    participant DB as 💾 Database

    U->>C: "What should I order this week?"
    C->>AI: Process message
    AI->>RAG: Get relevant context
    RAG->>DB: Query current stock levels
    RAG->>DB: Query historical usage
    RAG-->>AI: Context: low stock items, usage patterns
    AI->>AI: Generate recommendation
    AI-->>C: "Based on stock levels, order:<br/>- 5 boxes Cola (only 2 left)<br/>- 3 boxes Flour (weekend coming)"
    C-->>U: Display recommendation
    U->>C: "Create the order"
    C->>AI: Action intent detected
    AI->>API: Create purchase order
    API->>DB: Save order
    API-->>AI: Order #1234 created
    AI-->>C: "Order #1234 created. Send to supplier?"
    C-->>U: Show confirmation
```

## Microservices Boundaries

```mermaid
flowchart TB
    subgraph "Inventory Microservice"
        I1["Products API"]
        I2["Stock API"]
        I3["Movements API"]
        I4["Orders API"]
    end
    
    subgraph "Workforce Microservice"
        W1["Employees API"]
        W2["Skills API"]
        W3["Availability API"]
        W4["Schedule API"]
    end
    
    subgraph "AI Microservice"
        A1["Chat API"]
        A2["Embeddings API"]
        A3["RAG Query API"]
    end
    
    subgraph "Notifications Microservice"
        N1["Push Notifications"]
        N2["Email"]
        N3["WhatsApp (future)"]
    end
    
    subgraph "Auth Microservice"
        AU1["Login/Register"]
        AU2["Token Management"]
        AU3["Permissions"]
    end
```

## MCP Integration (Future)

```mermaid
flowchart LR
    subgraph "AI Assistant"
        LLM["LLM<br/>(Claude/GPT)"]
    end
    
    subgraph "MCP Servers"
        MCP1["📦 Inventory MCP<br/>- Check stock<br/>- Add stock<br/>- Create order"]
        MCP2["👥 Workforce MCP<br/>- Get schedule<br/>- Assign shift<br/>- Check availability"]
        MCP3["📅 Calendar MCP<br/>- Get holidays<br/>- Mark busy days"]
    end
    
    subgraph "Backend Services"
        API["REST API"]
        DB[(Database)]
    end
    
    LLM <-->|"MCP Protocol"| MCP1
    LLM <-->|"MCP Protocol"| MCP2
    LLM <-->|"MCP Protocol"| MCP3
    MCP1 --> API
    MCP2 --> API
    MCP3 --> API
    API --> DB
```

---

## Technology Stack

| Layer | Technology | Why |
|-------|------------|-----|
| **Backend** | Go | Fast, simple, great for APIs |
| **Database** | PostgreSQL | Reliable, feature-rich |
| **Cache** | Redis | Fast sessions, rate limiting |
| **Vector DB** | pgvector | Embeddings in PostgreSQL |
| **AI/LLM** | Ollama → OpenAI | Local dev, cloud prod |
| **Frontend** | React/Next.js or Flutter | Mobile-first |
| **Auth** | JWT + OAuth | Secure, standard |
| **Deploy** | Kubernetes | Scale, reliability |
| **MCP** | Go MCP SDK | AI tool integration |

---

## MVP Scope (Phase 1)

### Option A: Start with Inventory
```
Week 1-2: Products + Stock CRUD
Week 3-4: Stock movements + history
Week 5-6: AI chat for inventory queries
Week 7-8: Mobile UI + alerts
```

### Option B: Start with Workforce
```
Week 1-2: Employees + Skills CRUD
Week 3-4: Availability + Shifts
Week 5-6: AI-assisted schedule generation
Week 7-8: Mobile UI + notifications
```

### Recommendation: Start with Inventory
- Simpler domain to learn with
- More immediate value (know what to order)
- AI use cases are clearer
- Can add workforce later
