package rolestore

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/utils"
	_ "modernc.org/sqlite"
)

func initDB(filepath string) (*sql.DB, error) {
	if filepath != ":memory:" && utils.FileExists(filepath) {
		err := os.Remove(filepath) // remove database

		if err != nil {
			return nil, err
		}
	}

	dsn := filepath + "?parseTime=true"
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initStore(filepath string) (StoreInterface, error) {
	db, err := initDB(filepath)

	if err != nil {
		return nil, err
	}

	store, err := NewStore(NewStoreOptions{
		DB:                  db,
		RoleTableName:       "roles_role_table",
		EntityRoleTableName: "roles_entity_role_table",
		AutomigrateEnabled:  true,
		DebugEnabled:        true,
		SqlLogger:           slog.New(slog.NewTextHandler(os.Stdout, nil)),
	})

	if err != nil {
		return nil, err
	}

	if store == nil {
		return nil, errors.New("unexpected nil store")
	}

	return store, nil
}

func TestStoreWithTx(t *testing.T) {
	store, err := initStore("test_store_with_tx.db")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if store == nil {
		t.Fatal("unexpected nil store")
	}

	db := store.DB()

	if db == nil {
		t.Fatal("unexpected nil db")
	}

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	tx, err := db.Begin()

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if tx == nil {
		t.Fatal("unexpected nil tx")
	}

	txCtx := database.Context(context.Background(), tx)

	// create role
	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = store.RoleCreate(txCtx, role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// update role
	role.SetTitle("ROLE_TITLE_2")
	err = store.RoleUpdate(txCtx, role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check role
	roleFound, errFind := store.RoleFindByID(database.Context(context.Background(), db), role.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be nil, as transaction not committed")
	}

	if err := tx.Commit(); err != nil {
		t.Fatal("unexpected error:", err)
	}

	// check role
	roleFound, errFind = store.RoleFindByID(database.Context(context.Background(), db), role.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound == nil {
		t.Fatal("Role MUST be not nil, as transaction committed")
	}

	if roleFound.Title() != "ROLE_TITLE_2" {
		t.Fatal("Role MUST be ROLE_TITLE_2, as transaction committed")
	}
}
