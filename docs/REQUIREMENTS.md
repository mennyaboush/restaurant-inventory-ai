# 📋 Product Requirements Document

## Overview
Restaurant inventory management system for a Pizza & Falafel restaurant.

---

## 🎯 Core Problem Statement

**Current Pain:**
- Don't know what's in stock at any time
- Manual tracking is error-prone
- Running out of items during busy hours
- Waste from expired products
- Time-consuming to prepare orders

**Solution:**
AI-powered inventory system that understands natural language (text/voice), tracks stock in real-time, and proactively alerts about low stock and expiring items.

---

## 📦 Unit System (CRITICAL)

### The Challenge
Products come in different units and sizes:

```
┌─────────────────────────────────────────────────────────────────────────┐
│                         UNIT COMPLEXITY                                 │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  DRINKS:                                                                │
│  ├── Small Cola Box = 24 bottles (330ml)                               │
│  ├── Large Cola Box = 6 bottles (1.5L)                                 │
│  └── User says: "took 3 colas" → Which size? How many?                 │
│                                                                         │
│  VEGETABLES:                                                            │
│  ├── Can order by: KG, Box, or Individual                              │
│  ├── 1 Box of onions ≈ 10 KG ≈ ~50 onions                             │
│  └── User says: "took 10 onions" → 10 pieces, not 10 KG!              │
│                                                                         │
│  AI MUST:                                                               │
│  ├── Understand context                                                │
│  ├── Ask for clarification when ambiguous                              │
│  └── Detect inconsistencies ("You took 100 onions? That seems high")   │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### Data Model for Units

```
PRODUCT
├── name: "Cola"
├── category: "Drinks"
├── variants: [
│     {
│       name: "Small (330ml)",
│       sku: "COLA-SM",
│       units: [
│         { type: "box", contains: 24, contains_unit: "bottle" },
│         { type: "bottle", is_base: true }
│       ]
│     },
│     {
│       name: "Large (1.5L)",
│       sku: "COLA-LG", 
│       units: [
│         { type: "box", contains: 6, contains_unit: "bottle" },
│         { type: "bottle", is_base: true }
│       ]
│     }
│   ]
└── default_expiry_days: 365

PRODUCT (Vegetable Example)
├── name: "Onion"
├── category: "Vegetables"
├── variants: [
│     {
│       name: "Yellow Onion",
│       sku: "ONION-YEL",
│       units: [
│         { type: "box", contains: 10, contains_unit: "kg" },
│         { type: "kg", contains: 5, contains_unit: "piece" },
│         { type: "piece", is_base: true }
│       ]
│     }
│   ]
└── default_expiry_days: 14
```

### AI Understanding Examples

| User Says | AI Interprets | AI Response |
|-----------|---------------|-------------|
| "Took 3 colas" | Ambiguous | "Small (330ml) or Large (1.5L)?" |
| "Took 3 small colas" | 3 bottles of 330ml | "Got it, -3 small cola bottles" |
| "Took a box of small colas" | 24 bottles | "Got it, -1 box (24 bottles)" |
| "Took 10 onions" | 10 pieces | "Got it, -10 onion pieces" |
| "Received 5kg onions" | 5 kg | "Got it, +5kg onions" |
| "Took 100 onions" | Suspicious | "100 onions seems high. Did you mean 10?" |

---

## 📁 Product Categories

| Category | Examples | Unit Types | Expiry |
|----------|----------|------------|--------|
| Drinks | Cola, Fanta, Water, Juice | Box, Bottle | ~1 year |
| Disposable | Cups, napkins, forks, straws | Box, Pack, Piece | N/A |
| Packaging | Pizza boxes, containers, bags | Box, Pack, Piece | N/A |
| Vegetables | Tomatoes, onions, lettuce | Box, KG, Piece | 3-14 days |
| Canned | Tomato sauce, corn, olives | Box, Can | ~2 years |
| Basic | Flour, oil, yeast, sugar | Bag/Box, KG | 6-12 months |
| Spices | Za'atar, cumin, paprika | Box, Package, Gram | ~1 year |
| Cleaning | Soap, sanitizer, gloves | Box, Bottle, Piece | N/A |
| Dairy | Cheese, milk, labaneh | Box, KG, Unit | 1-4 weeks |
| Meat/Protein | Chicken, falafel mix | Box, KG | 3-7 days |
| Frozen | Fries, frozen items | Box, KG | 3-6 months |
| Bread | Pita, laffa, dough | Bag, Piece | 1-3 days |
| Sauces | Tahini, hummus, amba | Bucket, Bottle, KG | 2-4 weeks |

---

## ⏰ Expiration Handling

### Rules
1. Each product has a `default_expiry_days`
2. When stock arrives, expiry = today + default_expiry_days
3. User CAN override with actual expiry date
4. If user doesn't know → use default

### Alerts
- **7 days before:** "Milk expires in 7 days - use it first!"
- **3 days before:** "URGENT: Tomatoes expire in 3 days!"
- **Expired:** "5kg tomatoes expired yesterday - mark as waste?"

### Waste Tracking
```
When item expires/thrown:
├── Record as "waste"
├── Track reason: expired / damaged / other
├── Use for analytics: "You waste ₪500/month on vegetables"
└── Suggest: "Order smaller quantities of lettuce"
```

---

## 📊 Stock Movements

### Movement Types

| Type | Trigger | Example |
|------|---------|---------|
| **IN** | Delivery received | "+5 boxes cola from Supplier X" |
| **OUT** | Used in kitchen | "-10 onions" |
| **WASTE** | Expired/damaged | "-2kg tomatoes (expired)" |
| **ADJUSTMENT** | Sync/count | "Count shows 50, system shows 48, adjust +2" |
| **RETURN** | Returned to supplier | "-2 boxes (damaged on arrival)" |

### Who Can Do What

| Role | IN | OUT | WASTE | ADJUST |
|------|----|----|-------|--------|
| Owner | ✅ | ✅ | ✅ | ✅ |
| Manager | ✅ | ✅ | ✅ | ✅ |
| Employee | ❌ | ✅ | ✅ | ❌ |

---

## 🔔 Alerts System

### Alert Types

| Alert | Trigger | Notify |
|-------|---------|--------|
| Low Stock | quantity < min_level | Owner (WhatsApp) |
| Expiring Soon | expiry < 7 days | Owner (WhatsApp) |
| Expired | expiry passed | Owner (Push + WhatsApp) |
| Sync Request | Weekly or suspicious activity | Owner |
| Large Movement | Unusually large quantity | Owner |

### Jewish Calendar Integration
```
Special dates (need MORE stock):
├── Rosh Chodesh (1st of Hebrew month): +20% for 1-2 days
├── Shabbat: Closed (no alerts Friday evening - Saturday)
├── Holidays: Variable (Pesach, Sukkot, etc.)
└── Fast days: Less stock needed
```

---

## 👥 User Roles

### Role Definitions

```
┌─────────────────────────────────────────────────────────────────────────┐
│  OWNER (Your Father)                                                    │
├─────────────────────────────────────────────────────────────────────────┤
│  ✅ Full access to everything                                           │
│  ✅ User management (add/remove managers)                               │
│  ✅ View analytics and reports                                          │
│  ✅ Configure alerts and thresholds                                     │
│  ✅ Approve orders                                                      │
│  ✅ Sync/adjust inventory                                               │
└─────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────┐
│  MANAGER (Shift Manager)                                                │
├─────────────────────────────────────────────────────────────────────────┤
│  ✅ Receive deliveries (stock IN)                                       │
│  ✅ Record usage (stock OUT)                                            │
│  ✅ Record waste                                                        │
│  ✅ View stock levels                                                   │
│  ✅ Request sync                                                        │
│  ❌ Cannot manage users                                                 │
│  ❌ Cannot approve large orders                                         │
└─────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────┐
│  EMPLOYEE                                                               │
├─────────────────────────────────────────────────────────────────────────┤
│  ✅ Record usage (stock OUT) - what they took                          │
│  ✅ View stock levels (read only)                                       │
│  ✅ Report waste                                                        │
│  ❌ Cannot receive deliveries                                           │
│  ❌ Cannot adjust inventory                                             │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 🛒 Suppliers & Ordering

### Data Model
```
SUPPLIER
├── name: "Coca Cola Israel"
├── phone: "03-xxx-xxxx"
├── contact_method: "phone" | "app" | "whatsapp"
├── min_order_value: 500  (₪)
├── delivery_days: ["sunday", "tuesday", "thursday"]
└── notes: "Call before 10am"

PRODUCT_SUPPLIER
├── product_id
├── supplier_id
├── price_per_unit: 45.00
├── unit_type: "box"
├── min_quantity: 5
├── last_price_update: "2026-01-15"
└── is_preferred: true
```

### Order Flow (MVP)
```
1. System detects low stock
2. AI suggests: "You need to order: 5 boxes Cola, 10kg Tomatoes"
3. Owner reviews and approves
4. System generates order list (NOT auto-send)
5. Owner calls/orders manually
6. When delivery arrives → Manager records IN
```

### Order Flow (Future)
```
1-4. Same as above
5. System sends order via API/WhatsApp
6. Supplier confirms
7. Track delivery status
8. Auto-record when confirmed received
```

---

## 🤖 AI Agent Requirements

### Input Methods
- **Text:** Chat in app (Hebrew + English)
- **Voice:** Speech-to-text → process (Hebrew primarily)

### Understanding Context
```
User: "קיבלנו היום 5 ארגזים קולה"
AI understands:
├── Action: Stock IN
├── Product: Cola (which variant? ASK)
├── Quantity: 5 boxes
├── When: Today
└── Response: "Small cola or large cola?"

User: "קטנה"
AI: "Got it. Added 5 boxes of small cola (120 bottles)"
```

### Proactive Behaviors
1. **Clarify ambiguity:** "Which size?" "That seems like a lot, confirm?"
2. **Suggest corrections:** "You have 10 colas, you said took 15. Update stock?"
3. **Learn patterns:** Track requests system doesn't understand
4. **Sync prompts:** "It's been a week, want to do inventory count?"

### Example Conversations

```
👤: "מה המצב במלאי?"
🤖: "Stock Status:
     🔴 Low: Tomatoes (2kg left, need 10kg)
     🟡 Expiring: Milk (3 days left)
     🟢 Good: 45 other items
     
     Should I prepare an order list?"

👤: "כן"
🤖: "Suggested order:
     - Tomatoes: 15kg (Supplier: Fresh Veggies, ₪8/kg)
     - Onions: 1 box (Supplier: Fresh Veggies, ₪45)
     - Cola small: 3 boxes (Supplier: Coca Cola, ₪55/box)
     
     Total: ₪330
     Approve?"
```

### Learning System
```
Track unsupported requests:
├── User asks: "How much did we spend this month?"
├── System: "I can't do that yet, but I've noted it"
├── Log: { request: "spending report", count: 5, priority: "high" }
└── Developer sees: "Users want spending reports" → Add feature
```

---

## 🌍 Localization

### MVP
- **Hebrew (עברית):** Primary language, MUST support
- **Interface:** RTL (right-to-left) layout
- **Currency:** ₪ (ILS)
- **Dates:** Support Hebrew calendar for holidays

### Nice to Have
- **English:** Secondary language option
- **Arabic:** Future consideration

---

## 📱 User Interface

### Approach: Chat-First with Visual Backup

```
┌─────────────────────────────────────────────────────────────────────────┐
│                         APP STRUCTURE                                   │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│   PRIMARY: Chat Interface (80% of interactions)                        │
│   ├── Natural language input                                           │
│   ├── Voice input option                                               │
│   ├── Quick action buttons                                             │
│   └── AI responses with action cards                                   │
│                                                                         │
│   SECONDARY: Visual Screens (when needed)                              │
│   ├── Dashboard: Stock overview, alerts                                │
│   ├── Products: Browse/search all items                                │
│   ├── History: Movement log                                            │
│   └── Settings: Users, suppliers, thresholds                           │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 🔮 Future Features (Not MVP)

| Feature | Priority | Notes |
|---------|----------|-------|
| POS Integration | High | Auto-deduct when sold |
| Auto-ordering | Medium | Send orders to suppliers |
| Recipe costing | Medium | Cost per dish |
| Workforce management | High | Phase 2 |
| Multiple locations | Low | If business grows |
| Barcode scanning | Medium | Speed up input |
| Analytics dashboard | Medium | Trends, insights |
| WhatsApp bot | High | Interact via WhatsApp |

---

## ✅ MVP Scope Summary

### Must Have (v1.0)
- [ ] Products with variants and complex units
- [ ] Stock tracking (boxes, kg, pieces)
- [ ] Stock movements (IN, OUT, WASTE, ADJUST)
- [ ] Basic AI chat (Hebrew text)
- [ ] Low stock alerts (in-app)
- [ ] Default expiration tracking
- [ ] Simple order list generation
- [ ] Owner + Manager roles
- [ ] Mobile-friendly web app

### Should Have (v1.1)
- [ ] WhatsApp notifications
- [ ] Voice input
- [ ] Supplier management
- [ ] Jewish calendar holidays
- [ ] Sync/count feature
- [ ] Basic analytics

### Nice to Have (v1.2+)
- [ ] POS integration
- [ ] Auto-ordering
- [ ] Employee role
- [ ] English language
- [ ] Learning system for unsupported requests
