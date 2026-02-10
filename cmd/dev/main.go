package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

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

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func waitForDB(dsn string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		db, err := sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			db.Close()
		}
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("timed out waiting for db")
}

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

func runMigrations(db *sql.DB, migrationsDir string) error {
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		names = append(names, e.Name())
	}
	sort.Strings(names)
	for _, name := range names {
		path := filepath.Join(migrationsDir, name)
		upSQL, err := extractUpSQL(path)
		if err != nil {
			// skip files that don't match
			continue
		}
		if strings.TrimSpace(upSQL) == "" {
			continue
		}
		if _, err := db.Exec(upSQL); err != nil {
			return fmt.Errorf("migration %s failed: %w", name, err)
		}
		fmt.Printf("applied migration %s\n", name)
	}
	return nil
}

func main() {
	var (
		doUp       = flag.Bool("up", true, "start docker-compose and postgres")
		doMigrate  = flag.Bool("migrate", true, "run migrations")
		doTest     = flag.Bool("test", true, "run integration tests")
		doDown     = flag.Bool("down", false, "stop docker-compose at the end")
		timeoutSec = flag.Int("timeout", 60, "seconds to wait for Postgres")
	)
	flag.Parse()

	if _, err := os.Stat(".env"); err == nil {
		if err := loadEnv(".env"); err != nil {
			log.Fatalf("failed loading .env: %v", err)
		}
	}

	if *doUp {
		fmt.Println("Starting docker-compose...")
		if err := runCmd("docker-compose", "up", "-d"); err != nil {
			log.Fatalf("docker-compose up failed: %v", err)
		}
	}

	dsn := connStringFromEnv()
	if err := waitForDB(dsn, time.Duration(*timeoutSec)*time.Second); err != nil {
		log.Fatalf("db not ready: %v", err)
	}
	fmt.Println("Postgres is ready")

	if *doMigrate {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("open db: %v", err)
		}
		defer db.Close()
		if err := runMigrations(db, "migrations"); err != nil {
			log.Fatalf("migrations failed: %v", err)
		}
	}

	if *doTest {
		fmt.Println("Running integration tests (go test -tags=integration ./... )...")
		if err := runCmd("go", "test", "-tags=integration", "./...", "-v"); err != nil {
			log.Fatalf("tests failed: %v", err)
		}
	}

	if *doDown {
		fmt.Println("Stopping docker-compose...")
		if err := runCmd("docker-compose", "down"); err != nil {
			log.Fatalf("docker-compose down failed: %v", err)
		}
	}

	fmt.Println("dev command finished")
}
