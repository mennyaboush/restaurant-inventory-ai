# 🎯 Final Design Decisions

## Quick Reference for Development

---

## 1. Product Structure

**Decision: Separate products per size/variant**

```
✅ CORRECT:
├── Product: "Cola Small 330ml"
│   └── Box contains: 24 bottles
├── Product: "Cola Large 1.5L"  
│   └── Box contains: 6 bottles
└── Product: "Cola Zero Small 330ml"
    └── Box contains: 24 bottles

❌ NOT:
└── Product: "Cola"
    └── Variants: [Small, Large, Zero...]
```

**Why:** Simpler data model, clearer for users, easier for AI to understand.

---

## 2. AI Clarification Rule

**Decision: ALWAYS ask before changing stock**

```
┌─────────────────────────────────────────────────────────────────────────┐
│                      GOLDEN RULE                                        │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│   If AI is not 100% certain → ASK before changing stock                │
│                                                                         │
│   User: "לקחתי קולה"                                                    │
│   AI: "איזה קולה? קטנה 330 או גדולה 1.5?"                               │
│   User: "קטנה"                                                          │
│   AI: "כמה בקבוקים?"                                                    │
│   User: "3"                                                             │
│   AI: "עדכנתי: -3 בקבוקי קולה קטנה" ✅                                   │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Sync Process

**Decision: Visual list with edit option**

```
┌─────────────────────────────────────────────────────────────────────────┐
│                      SYNC SCREEN                                        │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│   📦 סנכרון מלאי                                        [שמור שינויים]  │
│                                                                         │
│   🔍 חיפוש...                                                           │
│                                                                         │
│   משקאות                                                                │
│   ┌─────────────────────────────────────────────────┐                  │
│   │ קולה קטנה 330    │ 4 ארגזים + 12 בק' │ [ערוך]  │                  │
│   │ קולה גדולה 1.5   │ 2 ארגזים + 0 בק'  │ [ערוך]  │                  │
│   │ פאנטה קטנה       │ 3 ארגזים + 6 בק'  │ [ערוך]  │                  │
│   └─────────────────────────────────────────────────┘                  │
│                                                                         │
│   ירקות                                                                 │
│   ┌─────────────────────────────────────────────────┐                  │
│   │ עגבניות          │ 8 ק"ג             │ [ערוך]  │                  │
│   │ בצל              │ 5 ק"ג             │ [ערוך]  │                  │
│   └─────────────────────────────────────────────────┘                  │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘

When user clicks [ערוך]:
┌─────────────────────────────────────────────┐
│  קולה קטנה 330                              │
│  ─────────────────                          │
│  ארגזים: [4] ▼▲                             │
│  בקבוקים בודדים: [12] ▼▲                    │
│                                             │
│  [ביטול]              [שמור]                │
└─────────────────────────────────────────────┘
```

---

## 4. Initial Setup

**Decision: Start with Sync to populate inventory**

```
First time flow:
1. Owner logs in
2. System: "Welcome! Let's set up your inventory"
3. Owner adds products (or imports from template)
4. Owner does first Sync to enter current quantities
5. System ready to use!
```

---

## 5. "Item Finished" Flow

**Decision: Clarify → Confirm → Set to Zero**

```
User: "נגמר הקולה"

AI: "איזה קולה נגמרה?"
   [קולה קטנה 330]  [קולה גדולה 1.5]  [אחר]

User: [קולה קטנה 330]

AI: "לפי המערכת היו 2 ארגזים + 5 בקבוקים (53 בקבוקים).
     לעדכן למלאי 0?"
   [כן, נגמר]  [לא, יש עוד]

User: [כן, נגמר]

AI: "עדכנתי: קולה קטנה 330 - מלאי 0 ✅
     רוצה להוסיף להזמנה הבאה?"
```

---

## 6. Data Model Summary

```
┌─────────────────────────────────────────────────────────────────────────┐
│                      SIMPLIFIED DATA MODEL                              │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│   PRODUCT                                                               │
│   ├── id (UUID)                                                        │
│   ├── name: "קולה קטנה 330"                                            │
│   ├── name_en: "Cola Small 330ml" (optional)                           │
│   ├── category_id → CATEGORY                                           │
│   ├── unit_type: "bottle" | "kg" | "piece" | "unit"                   │
│   ├── box_size: 24 (how many units in a box, null if no box)          │
│   ├── default_expiry_days: 365                                         │
│   ├── min_stock_level: 48 (in base units)                             │
│   ├── is_active: true                                                  │
│   └── created_at, updated_at                                           │
│                                                                         │
│   STOCK                                                                 │
│   ├── id (UUID)                                                        │
│   ├── product_id → PRODUCT                                             │
│   ├── quantity_boxes: 2                                                │
│   ├── quantity_units: 5 (loose items from opened box)                  │
│   ├── expiry_date: nullable                                            │
│   └── updated_at                                                        │
│                                                                         │
│   STOCK_MOVEMENT                                                        │
│   ├── id (UUID)                                                        │
│   ├── product_id → PRODUCT                                             │
│   ├── type: "IN" | "OUT" | "WASTE" | "ADJUSTMENT"                      │
│   ├── boxes_change: +5 or -2                                           │
│   ├── units_change: +10 or -3                                          │
│   ├── reason: nullable (for waste/adjustment)                          │
│   ├── user_id → USER                                                   │
│   └── created_at                                                        │
│                                                                         │
│   CATEGORY                                                              │
│   ├── id (UUID)                                                        │
│   ├── name: "משקאות"                                                   │
│   ├── name_en: "Drinks"                                                │
│   └── sort_order: 1                                                    │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 7. MVP Feature List (Prioritized)

### Phase 1: Foundation (Weeks 1-3)
- [x] Go project structure
- [ ] Product model & CRUD API
- [ ] Category model & CRUD API  
- [ ] Stock model & basic tracking
- [ ] Stock movements (IN/OUT)
- [ ] PostgreSQL database setup
- [ ] Basic authentication (Owner only)

### Phase 2: Core Features (Weeks 4-6)
- [ ] Stock movement history
- [ ] Low stock calculation
- [ ] Expiration tracking
- [ ] Sync screen (visual list)
- [ ] Basic web UI (mobile-friendly)
- [ ] Hebrew RTL support

### Phase 3: AI Chat (Weeks 7-9)
- [ ] Chat interface
- [ ] Intent recognition (add/remove/check stock)
- [ ] Product name matching
- [ ] Clarification flow
- [ ] Basic Hebrew NLP

### Phase 4: Alerts & Polish (Weeks 10-12)
- [ ] In-app notifications
- [ ] WhatsApp integration (alerts)
- [ ] Manager role
- [ ] Order list generation
- [ ] Jewish calendar holidays

---

## 8. Technology Decisions

| Component | Choice | Reason |
|-----------|--------|--------|
| Backend | Go | Learning goal + performance |
| Database | PostgreSQL | Reliable, full-featured |
| Cache | Redis | Sessions, rate limiting |
| AI/LLM | Ollama (dev) → OpenAI (prod) | Cost + quality balance |
| Frontend | Next.js or SvelteKit | Modern, fast, SSR |
| Mobile | PWA (Progressive Web App) | Works on any device |
| Auth | JWT | Stateless, standard |
| Deploy | Docker → Kubernetes | Learning goal |
| Hebrew NLP | OpenAI/Claude API | Best Hebrew support |

---

## Ready to Start! 🚀

All major decisions are documented. Time to write code!
