package rolestore

import (
	"context"
	"database/sql"

	"github.com/dromara/carbon/v2"
)

type StoreInterface interface {
	AutoMigrate() error
	EnableDebug(debug bool)
	DB() *sql.DB

	RoleCount(ctx context.Context, options RoleQueryInterface) (int64, error)
	RoleCreate(ctx context.Context, role RoleInterface) error
	RoleDelete(ctx context.Context, role RoleInterface) error
	RoleDeleteByID(ctx context.Context, id string) error
	RoleFindByHandle(ctx context.Context, handle string) (RoleInterface, error)
	RoleFindByID(ctx context.Context, id string) (RoleInterface, error)
	RoleList(ctx context.Context, query RoleQueryInterface) ([]RoleInterface, error)
	RoleSoftDelete(ctx context.Context, role RoleInterface) error
	RoleSoftDeleteByID(ctx context.Context, id string) error
	RoleUpdate(ctx context.Context, role RoleInterface) error
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
