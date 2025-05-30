package ent

import (
	"log/slog"
	"testing"

	"github.com/dexidp/dex/storage"
	"github.com/dexidp/dex/storage/conformance"
)

func newSQLiteStorage() storage.Storage {
	logger := slog.New(slog.DiscardHandler)

	cfg := SQLite3{File: ":memory:"}
	s, err := cfg.Open(logger)
	if err != nil {
		panic(err)
	}
	return s
}

func TestSQLite3(t *testing.T) {
	conformance.RunTests(t, newSQLiteStorage)
}
