package rolestore

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/gouniverse/sb"
)

// NewStoreOptions define the options for creating a new block store
type NewStoreOptions struct {
	// RoleTableName is the name of the role table
	RoleTableName string

	// EntityRoleTableName is the name of the entity to role relation table
	EntityRoleTableName string

	// DB is the underlying database connection
	DB *sql.DB

	// DbDriverName is the database driver name/type
	DbDriverName string

	// AutomigrateEnabled indicates whether to automatically migrate the database
	AutomigrateEnabled bool

	// DebugEnabled enables or disables the debug mode
	DebugEnabled bool

	// SqlLogger is the sql statement logger when debug mode is enabled, defaults to the default logger
	SqlLogger *slog.Logger
}

// NewStore creates a new block store
func NewStore(opts NewStoreOptions) (StoreInterface, error) {
	if opts.RoleTableName == "" {
		return nil, errors.New("role store: RoleTableName is required")
	}

	if opts.EntityRoleTableName == "" {
		return nil, errors.New("role store: EntityRoleTableName is required")
	}

	if opts.DB == nil {
		return nil, errors.New("shop store: DB is required")
	}

	if opts.DbDriverName == "" {
		opts.DbDriverName = sb.DatabaseDriverName(opts.DB)
	}

	if opts.SqlLogger == nil {
		opts.SqlLogger = slog.Default()
	}

	store := &store{
		roleTableName:       opts.RoleTableName,
		entityRoleTableName: opts.EntityRoleTableName,
		automigrateEnabled:  opts.AutomigrateEnabled,
		db:                  opts.DB,
		dbDriverName:        opts.DbDriverName,
		debugEnabled:        opts.DebugEnabled,
		sqlLogger:           opts.SqlLogger,
	}

	if store.automigrateEnabled {
		err := store.AutoMigrate()

		if err != nil {
			return nil, err
		}
	}

	return store, nil
}
