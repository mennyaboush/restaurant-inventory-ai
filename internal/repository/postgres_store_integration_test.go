//go:build integration
// +build integration

package repository

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/lib/pq"
	repostest "github.com/mennyaboush/restaurant-inventory-ai/internal/repository/test"
)

func prepareDBForTest(t *testing.T) *sql.DB {
	// Load env similar to existing test helper
	if ef := os.Getenv("TEST_ENV_FILE"); ef != "" {
		if _, err := os.Stat(ef); err == nil {
			if err := loadEnv(ef); err != nil {
				t.Fatalf("failed loading %s: %v", ef, err)
			}
		} else {
			t.Fatalf("TEST_ENV_FILE=%s not found: %v\nPlease create the file or run scripts/test-integration.sh to generate a .env.ci from .env.ci.example", ef, err)
		}
	} else {
		envPaths := []string{".env", filepath.Join("..", ".env"), filepath.Join("..", "..", ".env")}
		for _, p := range envPaths {
			if _, err := os.Stat(p); err == nil {
				if err := loadEnv(p); err != nil {
					t.Fatalf("failed loading %s: %v", p, err)
				}
				break
			}
		}
		required := []string{"POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB"}
		missing := []string{}
		for _, k := range required {
			if os.Getenv(k) == "" {
				missing = append(missing, k)
			}
		}
		if len(missing) > 0 {
			t.Fatalf("missing required environment variables: %v\nSet them in the environment or create a .env.ci from .env.ci.example and run scripts/test-integration.sh", missing)
		}
	}

	dsn := connStringFromEnv()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.Ping(); err != nil {
		t.Skipf("cannot reach Postgres at %s: %v", dsn, err)
	}

	// Ensure migrations applied for dev convenience
	// Apply migrations if missing (split statements to avoid multi-statement issues)
	migrationPath := func(name string) string {
		// search upward for migrations/<name>
		wd, _ := os.Getwd()
		dir := wd
		for {
			cand := filepath.Join(dir, "migrations", name)
			if _, err := os.Stat(cand); err == nil {
				return cand
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
			dir = parent
		}
		return filepath.Join("migrations", name) // fallback
	}

	if _, err := db.Exec("SELECT 1 FROM products LIMIT 1"); err != nil {
		up, rerr := extractUpSQL(migrationPath("001_create_products_table.sql"))
		if rerr != nil {
			t.Fatalf("cannot read products migration: %v", rerr)
		}
		stmts := strings.Split(up, ";")
		for _, s := range stmts {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if _, e := db.Exec(s); e != nil {
				t.Fatalf("failed to apply products migration stmt: %v", e)
			}
		}
	}
	if _, err := db.Exec("SELECT 1 FROM stocks LIMIT 1"); err != nil {
		up, rerr := extractUpSQL(migrationPath("002_create_stock_tables.sql"))
		if rerr != nil {
			t.Fatalf("cannot read stocks migration: %v", rerr)
		}
		stmts := strings.Split(up, ";")
		for _, s := range stmts {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if _, e := db.Exec(s); e != nil {
				t.Fatalf("failed to apply stocks migration stmt: %v", e)
			}
		}
	}

	// Verify tables exist; helpful debug if migrations didn't run
	var reg sql.NullString
	if err := db.QueryRow("SELECT to_regclass('public.stocks')").Scan(&reg); err != nil {
		t.Fatalf("failed checking stocks regclass: %v", err)
	}
	if !reg.Valid {
		// list tables for debugging
		rows, _ := db.Query("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname='public'")
		defer rows.Close()
		var tb []string
		for rows.Next() {
			var name string
			_ = rows.Scan(&name)
			tb = append(tb, name)
		}
		t.Fatalf("stocks table missing; public tables: %v", tb)
	}
	return db
}

func TestPostgresStore_CRUD_and_Edges(t *testing.T) {
	db := prepareDBForTest(t)
	defer db.Close()
	repostest.RunStoreIntegrationTests(t, func(db *sql.DB) repostest.Store {
		return NewPostgresStore(db)
	}, db)
}
