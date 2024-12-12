package rolestore

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/gouniverse/base/database"
)

// == TYPE ====================================================================

type store struct {
	// roleTableName is the name of the role table
	roleTableName string

	// entityRoleTableName is the name of the role entity relation table
	entityRoleTableName string

	// db is the underlying database connection
	db *sql.DB

	// dbDriverName is the database driver name/type
	dbDriverName string

	// automigrateEnabled enables or disables automigration
	automigrateEnabled bool

	// debugEnabled enables or disables debug mode
	debugEnabled bool

	// sqlLogger is the sql logger used when debug mode is enabled
	sqlLogger *slog.Logger
}

// == INTERFACE ===============================================================

var _ StoreInterface = (*store)(nil) // verify it extends the interface

// PUBLIC METHODS ============================================================

// AutoMigrate auto-migrates the database schema
func (store *store) AutoMigrate() error {
	if store.db == nil {
		return errors.New("rolestore: database is nil")
	}

	sqlStr := store.sqlRoleTableCreate()

	if sqlStr == "" {
		return errors.New("rolestore: role table create sql is empty")
	}

	_, err := store.db.Exec(sqlStr)

	if err != nil {
		return err
	}

	sqlStr = store.sqlEntityRoleTableCreate()

	if sqlStr == "" {
		return errors.New("rolestore: entity role table create sql is empty")
	}

	_, err = store.db.Exec(sqlStr)

	if err != nil {
		return err
	}

	return nil
}

// DB returns the underlying database connection
func (store *store) DB() *sql.DB {
	return store.db
}

// EnableDebug - enables or disables the debug mode
func (st *store) EnableDebug(debug bool) {
	st.debugEnabled = debug
}

// logSql logs sql to the sql logger, if debug mode is enabled
func (store *store) logSql(sqlOperationType string, sql string, params ...interface{}) {
	if !store.debugEnabled {
		return
	}

	if store.sqlLogger != nil {
		store.sqlLogger.Debug("sql: "+sqlOperationType, slog.String("sql", sql), slog.Any("params", params))
	}
}

// toQuerableContext converts the context to a QueryableContext
func (store *store) toQuerableContext(ctx context.Context) database.QueryableContext {
	if database.IsQueryableContext(ctx) {
		return ctx.(database.QueryableContext)
	}

	return database.Context(ctx, store.db)
}
