//go:build integration
// +build integration

package repository

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

// loadEnv loads simple KEY=VALUE lines from a file into the process env.
func loadEnv(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		os.Setenv(key, val)
	}
	return scanner.Err()
}

func connStringFromEnv() string {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("POSTGRES_USER")
	db := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
}

// extractUpSQL reads a migrate-style SQL file and returns the Up section.
func extractUpSQL(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	content := string(data)
	upMarker := "-- +migrate Up"
	downMarker := "-- +migrate Down"
	i := strings.Index(content, upMarker)
	if i == -1 {
		return "", fmt.Errorf("no Up marker in %s", path)
	}
	j := strings.Index(content, downMarker)
	if j == -1 {
		j = len(content)
	}
	up := content[i+len(upMarker) : j]
	return strings.TrimSpace(up), nil
}

func TestPostgresIntegration(t *testing.T) {
	// Load env file. Prefer an explicit path from TEST_ENV_FILE (useful for CI
	// or explicit test runs). If not set, search upward from the package
	// directory to the repo root to find a project `.env` for convenience.
	if ef := os.Getenv("TEST_ENV_FILE"); ef != "" {
		if _, err := os.Stat(ef); err == nil {
			if err := loadEnv(ef); err != nil {
				t.Fatalf("failed loading %s: %v", ef, err)
			}
		} else {
			t.Fatalf("TEST_ENV_FILE=%s not found: %v", ef, err)
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
	}

	dsn := connStringFromEnv()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Skipf("cannot reach Postgres at %s: %v", dsn, err)
	}

	// Ensure products table exists. If missing, attempt to run migration file (dev convenience).
	_, err = db.Exec("SELECT 1 FROM products LIMIT 1")
	if err != nil {
		// attempt to run migration
		migPath := filepath.Join("migrations", "001_create_products_table.sql")
		upSQL, rerr := extractUpSQL(migPath)
		if rerr != nil {
			t.Fatalf("products table missing and cannot read migration: %v, readErr=%v", err, rerr)
		}
		if _, e := db.Exec(upSQL); e != nil {
			t.Fatalf("failed to run migration: %v", e)
		}
	}

	// Insert a sample product and read it back
	_, err = db.Exec(`INSERT INTO products (id, name, brand, size, container_type, box_size, price, category)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
        ON CONFLICT (brand,size,container_type) DO NOTHING`, "test-1", "Test Product", "TestBrand", 1, "box", 0, 1.23, "testcat")
	if err != nil {
		t.Fatalf("insert product failed: %v", err)
	}

	var name string
	err = db.QueryRow("SELECT name FROM products WHERE id=$1", "test-1").Scan(&name)
	if err != nil {
		t.Fatalf("select product failed: %v", err)
	}
	if name != "Test Product" {
		t.Fatalf("unexpected product name: %s", name)
	}
}
