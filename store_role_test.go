package rolestore

import (
	"context"
	"strings"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
)

func TestStoreRoleCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.RoleCount(context.Background(), NewRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")
	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.RoleCount(context.Background(), NewRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}

	err = store.RoleCreate(context.Background(), NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.RoleCount(context.Background(), NewRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 2 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreRoleCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreRoleDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleDelete(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	roleFound, err := store.RoleFindByID(context.Background(), role.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be nil")
	}

	roleFindWithDeleted, err := store.RoleList(context.Background(), NewRoleQuery().
		SetID(role.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) != 0 {
		t.Fatal("Role MUST be nil")
	}
}

func TestStoreRoleDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleDeleteByID(context.Background(), role.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	roleFound, err := store.RoleFindByID(context.Background(), role.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be nil")
	}

	roleFindWithDeleted, err := store.RoleList(context.Background(), NewRoleQuery().
		SetID(role.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) != 0 {
		t.Fatal("Role MUST NOT be found")
	}
}

func TestStoreRoleFindByHandle(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = role.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleCreate(database.Context(context.Background(), store.DB()), role)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	roleFound, errFind := store.RoleFindByHandle(database.Context(context.Background(), store.DB()), role.Handle())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound == nil {
		t.Fatal("Role MUST NOT be nil")
	}

	if roleFound.ID() != role.ID() {
		t.Fatal("IDs do not match")
	}

	if roleFound.Handle() != role.Handle() {
		t.Fatal("Handles do not match")
	}

	if roleFound.Title() != role.Title() {
		t.Fatal("Titles do not match")
	}

	if roleFound.Status() != role.Status() {
		t.Fatal("Statuses do not match")
	}

	if roleFound.Meta("education_1") != role.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if roleFound.Meta("education_2") != role.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if roleFound.Meta("education_3") != role.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreRoleFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = role.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ctx := database.Context(context.Background(), store.DB())
	err = store.RoleCreate(ctx, role)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	roleFound, errFind := store.RoleFindByID(ctx, role.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound == nil {
		t.Fatal("Role MUST NOT be nil")
	}

	if roleFound.ID() != role.ID() {
		t.Fatal("IDs do not match")
	}

	if roleFound.Handle() != role.Handle() {
		t.Fatal("Handles do not match")
	}

	if roleFound.Title() != role.Title() {
		t.Fatal("Titles do not match")
	}

	if roleFound.Status() != role.Status() {
		t.Fatal("Statuses do not match")
	}

	if roleFound.Meta("education_1") != role.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if roleFound.Meta("education_2") != role.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if roleFound.Meta("education_3") != role.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreRoleList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role1 := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE_1").
		SetTitle("ROLE_TITLE_1")

	role2 := NewRole().
		SetStatus(ROLE_STATUS_INACTIVE).
		SetHandle("ROLE_HANDLE_2").
		SetTitle("ROLE_TITLE_2")

	roles := []RoleInterface{
		role1,
		role2,
	}

	for _, role := range roles {
		err = store.RoleCreate(context.Background(), role)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	listActive, err := store.RoleList(context.Background(), NewRoleQuery().SetStatus(ROLE_STATUS_ACTIVE))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listActive) != 1 {
		t.Fatal("unexpected list length:", len(listActive))
	}

	listEmail, err := store.RoleList(context.Background(), NewRoleQuery().SetHandle("ROLE_HANDLE_2"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(listEmail) != 1 {
		t.Fatal("unexpected list length:", len(listEmail))
	}
}

func TestStoreRoleSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleSoftDelete(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if role.SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("Role MUST be soft deleted")
	}

	roleFound, errFind := store.RoleFindByID(context.Background(), role.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be soft deleted, so MUST be nil")
	}

	roleFindWithDeleted, err := store.RoleList(context.Background(), NewRoleQuery().
		SetSoftDeletedIncluded(true).
		SetID(role.ID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) == 0 {
		t.Fatal("Role MUST be soft deleted")
	}

	if strings.Contains(roleFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("Role MUST be soft deleted", role.SoftDeletedAt())
	}

	if !roleFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Role MUST be soft deleted")
	}
}

func TestStoreRoleSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	role := NewRole().
		SetStatus(ROLE_STATUS_ACTIVE).
		SetHandle("ROLE_HANDLE").
		SetTitle("ROLE_TITLE")

	err = store.RoleCreate(context.Background(), role)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.RoleSoftDeleteByID(context.Background(), role.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if role.SoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("Role MUST NOT be soft deleted, as it was soft deleted by ID")
	}

	roleFound, errFind := store.RoleFindByID(context.Background(), role.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if roleFound != nil {
		t.Fatal("Role MUST be nil")
	}
	query := NewRoleQuery().
		SetSoftDeletedIncluded(true).
		SetID(role.ID()).
		SetLimit(1)

	roleFindWithDeleted, err := store.RoleList(context.Background(), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(roleFindWithDeleted) == 0 {
		t.Fatal("Role MUST be soft deleted")
	}

	if strings.Contains(roleFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("Role MUST be soft deleted", role.SoftDeletedAt())
	}

	if !roleFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("Role MUST be soft deleted")
	}
}
