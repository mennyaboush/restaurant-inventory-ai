# 📋 Product Requirements Document (MVP v1.0)

## Overview
Simple, production-ready inventory management system for a Pizza & Falafel restaurant.

**Real User:** Restaurant owner who needs to know what's in stock and when to reorder.

---

## 🎯 Core Problem Statement

**Current Pain:**
- Don't know what's in stock at any time
- Manual tracking with pen and paper is error-prone
- Running out of items during busy hours
- No record of when items were last restocked
- Don't know when to reorder

**MVP Solution:**
Simple web interface to track inventory, log stock movements, and see what's low. Focus on basic tracking with option to add AI features later.

---

## 📦 Product & Stock Model (KISS Principle)

### Design Decision: One Product Per Size/Variant

Instead of complex variant systems, we create separate products for each size:

```
✅ SIMPLE APPROACH:
├── Product: "Coca Cola 330ml Can" (box_size: 24)
├── Product: "Coca Cola 1.5L Plastic" (box_size: 6)
├── Product: "Pepsi 330ml Can" (box_size: 24)
└── Product: "Onions" (box_size: null, sold by kg)

Why?
- Simpler to understand and manage
- Each product has ONE box size
- Clear for users and developers
- Easier to track and report on
```

### Stock Tracking: Boxes + Units

```
PRODUCT
├── id: "PROD-001"
├── name: "קוקה קולה 330 פחית"
├── brand: "Coca Cola"
├── size: 330 (ml)
├── container_type: "can"
├── box_size: 24 (items per box)
├── category: "drinks"
└── min_stock: 48 (alert threshold in units)

STOCK
├── product_id: "PROD-001"
├── quantity_boxes: 3 (full unopened boxes)
├── quantity_units: 12 (loose items from opened box)
└── Total = (3 × 24) + 12 = 84 bottles

STOCK MOVEMENT
├── product_id: "PROD-001"
├── type: "OUT"
├── boxes: -1
├── units: -5
├── performed_by: "Yosef"
├── reported_by: "Manager"
└── reason: "Sold to customer"
```

### Why This Works

| Scenario | How It's Handled |
|----------|------------------|
| "Took 3 colas" | User selects from product list: "Coca Cola 330ml Can" |
| "Received delivery" | User enters: +5 boxes for "Coca Cola 330ml Can" |
| "Opened a box" | System converts: boxes: -1, units: +24 |
| Items sold by weight | Set box_size to null, track only units (kg) |

---

## 📁 Product Categories (MVP)

| Category | Examples | Tracking Unit |
|----------|----------|---------------|
| Drinks | Cola, Fanta, Water, Juice | Boxes + Bottles |
| Vegetables | Tomatoes, onions, lettuce | KG or Pieces |
| Basic Ingredients | Flour, oil, yeast, sugar | Bags/Boxes |
| Dairy | Cheese, milk | KG or Units |
| Packaging | Pizza boxes, bags | Boxes + Pieces |

**Note:** Start with these 5 categories. Can add more later based on actual needs.

---

## 📊 Stock Movements (Audit Trail)

Every change to stock must be logged:

### Movement Types

| Type | When Used | Example |
|------|-----------|---------|
| **IN** | Stock received | "Received 5 boxes cola from supplier" |
| **OUT** | Stock used | "Used 10 tomatoes for pizzas" |
| **WASTE** | Thrown away | "2kg tomatoes expired" |
| **ADJUSTMENT** | Fix count | "Physical count shows 50, system has 48, adjust +2" |

### Who Tracking (IMPORTANT)

```
Every movement tracks TWO people:
├── performed_by: WHO actually did the action
└── reported_by: WHO logged it in the system

Examples:
1. Manager logs: "Yosef took 2 boxes of cola"
   - performed_by: Yosef
   - reported_by: Manager

2. Owner logs they received delivery
   - performed_by: Owner
   - reported_by: Owner (same person)
```

**Why?** Accountability and audit trail. Know who handled what and when.

---

## 🔔 Alerts (Simple Start)

### MVP Alerts (In-App Only)

Show badge or indicator when:
- **Low Stock:** Any product where total units < min_stock
- **Need Sync:** Prompt weekly to do inventory count

### Future Enhancements
- Push notifications (when app is installed on mobile)
- WhatsApp notifications (v2.0)
- Email reports (v2.0)

---

## 👥 User Roles (Start Simple)

### MVP: Owner Only

Start with single user (restaurant owner) who can:
- ✅ Add/edit products
- ✅ Record stock movements (IN/OUT/WASTE/ADJUST
)
- ✅ View stock levels and history
- ✅ Do inventory sync

### Phase 2: Add Manager Role

When needed, add manager accounts:
- ✅ Record movements (IN/OUT)
- ✅ View stock
- ❌ Cannot delete products or adjust inventory

### Phase 3: Add Employee Role

Simple employee tracking:
- ✅ Record what they took (OUT movements)
- ❌ Limited access to other features

---

## � User Interface (MVP)

### Mobile-First Web App

Simple, clean interface optimized for tablet/phone:

**Main Screen: Stock Overview**
```
┌─────────────────────────────────┐
│  📦 מלאי מסעדה                   │
├─────────────────────────────────┤
│  🔴 Low Stock (3)               │
│  ├─ קולה קטנה: 12 left          │
│  ├─ עגבניות: 2kg left           │
│  └─ קמח: 1 bag left             │
│                                 │
│  📋 All Products                │
│  🔄 Record Movement             │
│  ⚙️  Settings                    │
└─────────────────────────────────┘
```

**Record Movement Screen:**
```
┌─────────────────────────────────┐
│  🔄 רישום תנועה                  │
├─────────────────────────────────┤
│  Type: [IN] [OUT] [WASTE]       │
│                                 │
│  Product: [Search...]           │
│                                 │
│  Boxes: [ 0 ] ▲▼                │
│  Units: [ 0 ] ▲▼                │
│                                 │
│  Who: [Select person...]        │
│  Reason: [Optional note...]     │
│                                 │
│  [Cancel]  [Save Movement]      │
└─────────────────────────────────┘
```

### Language: Hebrew (RTL)

- Primary language: עברית
- Right-to-left layout
- English product names allowed
- Currency: ₪ (ILS)

---

## ✅ MVP Scope (v1.0 - Production Ready)

### Phase 1: Core Inventory (Weeks 1-4) ✅ IN PROGRESS

- [x] Data models (Product, Stock, StockMovement)
- [x] PostgreSQL database setup
- [x] REST API for products (CRUD)
- [x] REST API for stock operations
- [ ] Basic authentication (Owner login)
- [ ] Categories management
- [ ] Movement history view

### Phase 2: Web Interface (Weeks 5-6)

- [ ] Mobile-responsive web UI
- [ ] Hebrew RTL support
- [ ] Product list and search
- [ ] Record movement form
- [ ] Stock overview dashboard
- [ ] Low stock indicators

### Phase 3: Polish & Deploy (Weeks 7-8)

- [ ] Input validation and error handling
- [ ] Docker containerization
- [ ] Deploy to production (Kubernetes or simple VPS)
- [ ] User testing with restaurant owner
- [ ] Bug fixes and improvements

### Success Criteria

The MVP is ready when:
1. ✅ Owner can add products
2. ✅ Owner can record stock movements (IN/OUT)
3. ✅ Owner can see current stock levels
4. ✅ Owner can see movement history
5. ⬜ System shows low stock warnings
6. ⬜ Works on mobile browser (tablet/phone)
7. ⬜ Interface is in Hebrew
8. ⬜ Deployed and accessible from anywhere

---

## 🔮 Future Enhancements (v2.0+)

| Feature | Priority | When |
|---------|----------|------|
| **AI Chat Interface** | High | v2.0 - Natural language for movements |
| **Manager/Employee Roles** | Medium | v2.0 - When hiring staff |
| **Supplier Management** | Medium | v2.0 - Track prices and orders |
| **Expiration Tracking** | High | v2.1 - For perishables |
| **WhatsApp Notifications** | High | v2.1 - Alerts on phone |
| **Smart Order Suggestions** | Medium | v3.0 - AI predicts needs |
| **Recipe Costing** | Low | v3.0 - Cost per dish |
| **POS Integration** | Low | v3.0 - Auto-deduct on sales |
| **Mobile App** | Medium | v3.0 - Native iOS/Android |

---

## 📐 Technical Architecture (MVP)

### Simple Stack

```
┌─────────────────────────────────────────┐
│  Frontend: HTML + JavaScript (simple)   │
│  (or React if time allows)              │
└─────────────────────────────────────────┘
              ↓ REST API
┌─────────────────────────────────────────┐
│  Backend: Go                            │
│  - Chi router                           │
│  - JWT auth                             │
│  - Database queries                     │
└─────────────────────────────────────────┘
              ↓ SQL
┌─────────────────────────────────────────┐
│  Database: PostgreSQL                   │
│  - products table                       │
│  - stocks table                         │
│  - stock_movements table                │
│  - users table                          │
└─────────────────────────────────────────┘
```

### No Microservices YET

Start with a **monolith** - one service, one database, deployed together.

**Why?**
- Faster to develop
- Easier to debug
- Sufficient for MVP
- Can split into microservices later if needed

### Deployment Options

1. **Simple:** Single VPS (DigitalOcean droplet - $6/month)
2. **Learning:** Kubernetes cluster (more complex but teaches K8s)
3. **Free Tier:** Railway.app or Render.com (free for small apps)

**Recommendation:** Start with VPS, move to Kubernetes after MVP works.

---

## 🎯 Success Metrics

How we'll know it's working:

### User Adoption
- Owner uses it daily for 2 weeks
- No return to pen-and-paper tracking
- Owner recommends it

### Functionality
- Zero data loss (all movements logged correctly)
- <2 seconds response time on mobile
- Works offline-first (future: PWA)

### Business Value
- Owner knows stock status anytime
- Reduces "ran out" incidents
- Saves time (vs. manual tracking)

---

## 📝 Notes

### Design Philosophy: KISS (Keep It Simple, Stupid)

This MVP prioritizes:
1. **Working over Perfect** - Ship something usable fast
2. **Simple over Complex** - One product per variant, not nested objects
3. **Manual over Automated** - Owner orders manually, system just tracks
4. **Learn over Assume** - Build what's needed, not what might be needed

### AI Strategy

**MVP:** No AI yet. Focus on core inventory tracking.

**v2.0:** Add AI chat layer on top of working system:
- Chat interface for recording movements
- Natural language understanding
- Smart suggestions

**Why later?** AI is complex. Get the foundation right first.
