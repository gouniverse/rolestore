package rolestore

import (
	"context"
	"database/sql"

	"github.com/dromara/carbon/v2"
)

type StoreInterface interface {
	// AutoMigrate auto migrates the database schema
	AutoMigrate() error

	// EnableDebug enables or disables the debug mode
	EnableDebug(debug bool)

	// DB returns the underlying database connection
	DB() *sql.DB

	// == Role Methods =======================================================//

	// RoleCount returns the number of roles based on the given query options
	RoleCount(ctx context.Context, options RoleQueryInterface) (int64, error)

	// RoleCreate creates a new role
	RoleCreate(ctx context.Context, role RoleInterface) error

	// RoleDelete deletes a role
	RoleDelete(ctx context.Context, role RoleInterface) error

	// RoleDeleteByID deletes a role by its ID
	RoleDeleteByID(ctx context.Context, id string) error

	// RoleFindByHandle returns a role by its handle
	RoleFindByHandle(ctx context.Context, handle string) (RoleInterface, error)

	// RoleFindByID returns a role by its ID
	RoleFindByID(ctx context.Context, id string) (RoleInterface, error)

	// RoleList returns a list of roles based on the given query options
	RoleList(ctx context.Context, query RoleQueryInterface) ([]RoleInterface, error)

	// RoleSoftDelete soft deletes a role
	RoleSoftDelete(ctx context.Context, role RoleInterface) error

	// RoleSoftDeleteByID soft deletes a role by its ID
	RoleSoftDeleteByID(ctx context.Context, id string) error

	// RoleUpdate updates a role
	RoleUpdate(ctx context.Context, role RoleInterface) error

	// == EntityRole Methods =================================================//

	// EntityRoleCount returns the number of role entities mappings based on the given query options
	EntityRoleCount(ctx context.Context, options EntityRoleQueryInterface) (int64, error)

	// EntityRoleCreate creates a new role entity mapping
	EntityRoleCreate(ctx context.Context, entityRole EntityRoleInterface) error

	// EntityRoleDelete deletes a role entity mapping
	EntityRoleDelete(ctx context.Context, entityRole EntityRoleInterface) error

	// EntityRoleDeleteByID deletes a role entity mapping by its ID
	EntityRoleDeleteByID(ctx context.Context, id string) error

	// EntityRoleFindByEntityAndRole returns a role entity mapping by its entity type, entity ID and role ID
	EntityRoleFindByEntityAndRole(ctx context.Context, entityType string, entityID string, roleID string) (EntityRoleInterface, error)

	// EntityRoleFindByID returns a role entity mapping by its ID
	EntityRoleFindByID(ctx context.Context, id string) (EntityRoleInterface, error)

	// EntityRoleList returns a list of role entity mappings based on the given query options
	EntityRoleList(ctx context.Context, query EntityRoleQueryInterface) ([]EntityRoleInterface, error)

	// EntityRoleSoftDelete soft deletes a role entity mapping
	EntityRoleSoftDelete(ctx context.Context, entityRole EntityRoleInterface) error

	// EntityRoleSoftDeleteByID soft deletes a role entity mapping by its ID
	EntityRoleSoftDeleteByID(ctx context.Context, id string) error

	// EntityRoleUpdate updates a role entity mapping
	EntityRoleUpdate(ctx context.Context, entityRole EntityRoleInterface) error
}

type RoleInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// methods

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool

	// setters and getters

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) RoleInterface

	Handle() string
	SetHandle(handle string) RoleInterface

	ID() string
	SetID(id string) RoleInterface

	Memo() string
	SetMemo(memo string) RoleInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	Status() string
	SetStatus(status string) RoleInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) RoleInterface

	Title() string
	SetTitle(title string) RoleInterface

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) RoleInterface
}

type EntityRoleInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()

	// methods

	IsSoftDeleted() bool

	// setters and getters

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) EntityRoleInterface

	EntityType() string
	SetEntityType(entityType string) EntityRoleInterface

	EntityID() string
	SetEntityID(entityID string) EntityRoleInterface

	ID() string
	SetID(id string) EntityRoleInterface

	Memo() string
	SetMemo(memo string) EntityRoleInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error

	RoleID() string
	SetRoleID(roleID string) EntityRoleInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(softDeletedAt string) EntityRoleInterface

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) EntityRoleInterface
}

type UserInterface interface {
	// from dataobject

	Data() map[string]string
	DataChanged() map[string]string
	MarkAsNotDirty()
	Get(columnName string) string
	Set(columnName string, value string)

	// methods

	IsActive() bool
	IsInactive() bool
	IsSoftDeleted() bool
	IsUnverified() bool

	IsAdministrator() bool
	IsManager() bool
	IsSuperuser() bool

	IsRegistrationCompleted() bool

	// setters and getters

	BusinessName() string
	SetBusinessName(businessName string) UserInterface

	Country() string
	SetCountry(country string) UserInterface

	CreatedAt() string
	CreatedAtCarbon() carbon.Carbon
	SetCreatedAt(createdAt string) UserInterface

	Email() string
	SetEmail(email string) UserInterface

	ID() string
	SetID(id string) UserInterface

	FirstName() string
	SetFirstName(firstName string) UserInterface

	LastName() string
	SetLastName(lastName string) UserInterface

	Memo() string
	SetMemo(memo string) UserInterface

	Meta(name string) string
	SetMeta(name string, value string) error
	Metas() (map[string]string, error)
	SetMetas(metas map[string]string) error
	UpsertMetas(metas map[string]string) error

	MiddleNames() string
	SetMiddleNames(middleNames string) UserInterface

	Password() string
	SetPassword(password string) UserInterface

	Phone() string
	SetPhone(phone string) UserInterface

	ProfileImageUrl() string
	SetProfileImageUrl(profileImageUrl string) UserInterface

	Role() string
	SetRole(role string) UserInterface

	SoftDeletedAt() string
	SoftDeletedAtCarbon() carbon.Carbon
	SetSoftDeletedAt(deletedAt string) UserInterface

	Timezone() string
	SetTimezone(timezone string) UserInterface

	Status() string
	SetStatus(status string) UserInterface

	PasswordCompare(password string) bool

	UpdatedAt() string
	UpdatedAtCarbon() carbon.Carbon
	SetUpdatedAt(updatedAt string) UserInterface
}
