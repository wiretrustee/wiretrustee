package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/netbirdio/netbird/util"
)

type benchCase struct {
	name    string
	storeFn func(b *testing.B) Store
	size    int
}

var newFs = func(b *testing.B) Store {
	b.Helper()
	store, _ := NewFileStore(context.Background(), b.TempDir(), nil)
	return store
}

var newSqlite = func(b *testing.B) Store {
	b.Helper()
	store, _ := NewSqliteStore(context.Background(), b.TempDir(), nil)
	return store
}

func BenchmarkTest_StoreWrite(b *testing.B) {
	cases := []benchCase{
		{name: "FileStore_Write", storeFn: newFs, size: 100},
		{name: "SqliteStore_Write", storeFn: newSqlite, size: 100},
		{name: "FileStore_Write", storeFn: newFs, size: 500},
		{name: "SqliteStore_Write", storeFn: newSqlite, size: 500},
		{name: "FileStore_Write", storeFn: newFs, size: 1000},
		{name: "SqliteStore_Write", storeFn: newSqlite, size: 1000},
		{name: "FileStore_Write", storeFn: newFs, size: 2000},
		{name: "SqliteStore_Write", storeFn: newSqlite, size: 2000},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s_%d", c.name, c.size)
		store := c.storeFn(b)

		for i := 0; i < c.size; i++ {
			_ = newAccount(store, i)
		}

		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				i := c.size
				for pb.Next() {
					i++
					err := newAccount(store, i)
					require.NoError(b, err)
				}
			})
		})
	}
}

func BenchmarkTest_StoreRead(b *testing.B) {
	cases := []benchCase{
		{name: "FileStore_Read", storeFn: newFs, size: 100},
		{name: "SqliteStore_Read", storeFn: newSqlite, size: 100},
		{name: "FileStore_Read", storeFn: newFs, size: 500},
		{name: "SqliteStore_Read", storeFn: newSqlite, size: 500},
		{name: "FileStore_Read", storeFn: newFs, size: 1000},
		{name: "SqliteStore_Read", storeFn: newSqlite, size: 1000},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%s_%d", c.name, c.size)
		store := c.storeFn(b)

		for i := 0; i < c.size; i++ {
			_ = newAccount(store, i)
		}

		accounts := store.GetAllAccounts(context.Background())
		id := accounts[c.size-1].Id

		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = store.GetAccount(context.Background(), id)
				}
			})
		})
	}
}

func NewSqliteTestStore(t *testing.T, ctx context.Context, testFile string) (Store, func(), error) {
	dataDir := t.TempDir()
	err := util.CopyFileContents(testFile, filepath.Join(dataDir, "store.db"))
	if err != nil {
		t.Fatal(err)
	}

	store, err := NewSqliteStore(ctx, dataDir, nil)
	if err != nil {
		return nil, nil, err
	}

	return store, func() {
		store.Close(ctx)
		os.Remove(filepath.Join(dataDir, "store.db"))
	}, nil
}
