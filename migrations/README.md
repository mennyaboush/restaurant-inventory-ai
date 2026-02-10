# Database Migrations

This directory contains SQL migration files for database schema changes.

## Naming Convention
Files are named with a timestamp prefix for ordering:
- `001_create_inventory_table.sql`
- `002_add_supplier_table.sql`

## Structure
Each migration file contains:
- `-- +migrate Up` section for applying changes
- `-- +migrate Down` section for reverting changes
