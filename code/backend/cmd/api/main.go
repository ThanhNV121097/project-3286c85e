package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-3286c85e/backend/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
)

type app struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

type errorBody struct {
	Error errorDetail `json:"error"`
}

type errorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type greetingBody struct {
	Greeting string `json:"greeting"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := run(logger); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return errors.New("DATABASE_URL is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return fmt.Errorf("create database pool: %w", err)
	}
	defer db.Close()

	if err := applyMigrations(ctx, db); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}

	a := app{db: db, logger: logger}
	server := http.Server{
		Addr:              ":" + listenPort(),
		Handler:           a.routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info("listening", "addr", server.Addr)
	return server.ListenAndServe()
}

func listenPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	if port := os.Getenv("APP_PORT"); port != "" {
		return port
	}
	return "8080"
}

func (a app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", a.health)
	mux.HandleFunc("GET /v1/greeting", a.greeting)
	return mux
}

func (a app) health(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := a.db.Ping(ctx); err != nil {
		a.logger.Error("health database ping failed", "error", err)
		writeError(w, http.StatusServiceUnavailable, "service_unavailable", "Service unavailable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (a app) greeting(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var greeting string
	if err := a.db.QueryRow(ctx, "SELECT text FROM greetings WHERE id = 1").Scan(&greeting); err != nil {
		a.logger.Error("read greeting failed", "error", err)
		writeError(w, http.StatusInternalServerError, "internal_error", "Internal server error")
		return
	}
	writeJSON(w, http.StatusOK, greetingBody{Greeting: greeting})
}

func applyMigrations(ctx context.Context, db *pgxpool.Pool) error {
	if _, err := db.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
		version text PRIMARY KEY,
		applied_at timestamptz NOT NULL DEFAULT now()
	)`); err != nil {
		return fmt.Errorf("ensure schema_migrations: %w", err)
	}

	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return fmt.Errorf("read embedded migrations: %w", err)
	}

	var names []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".up.sql") {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)

	for _, name := range names {
		version := strings.TrimSuffix(name, ".up.sql")
		var exists bool
		if err := db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)", version).Scan(&exists); err != nil {
			return fmt.Errorf("check migration %s: %w", version, err)
		}
		if exists {
			continue
		}

		sqlBytes, err := migrations.Files.ReadFile(name)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", name, err)
		}

		tx, err := db.Begin(ctx)
		if err != nil {
			return fmt.Errorf("begin migration %s: %w", version, err)
		}
		if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("execute migration %s: %w", version, err)
		}
		if _, err := tx.Exec(ctx, "INSERT INTO schema_migrations (version) VALUES ($1)", version); err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("record migration %s: %w", version, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("commit migration %s: %w", version, err)
		}
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Default().Error("write response failed", "error", err)
	}
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, errorBody{Error: errorDetail{Code: code, Message: message}})
}
