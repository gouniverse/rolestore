package rolestore

import (
	"context"
	"strings"
	"testing"

	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
)

func TestStoreEntityRoleCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	count, err := store.EntityRoleCount(context.Background(), NewEntityRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 0 {
		t.Fatal("unexpected count:", count)
	}

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.EntityRoleCount(context.Background(), NewEntityRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 1 {
		t.Fatal("unexpected count:", count)
	}

	entityRole2 := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_02").
		SetRoleID("ROLE_02")

	err = store.EntityRoleCreate(context.Background(), entityRole2)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	count, err = store.EntityRoleCount(context.Background(), NewEntityRoleQuery())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if count != 2 {
		t.Fatal("unexpected count:", count)
	}
}

func TestStoreEntityRoleCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestStoreEntityRoleCreate_Duplicate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err == nil {
		t.Fatal("must return error as duplicated entity to role relationship")
	}
}

func TestStoreEntityRoleDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityRoleDelete(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	entityRoleFound, err := store.EntityRoleFindByID(context.Background(), entityRole.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityRoleFound != nil {
		t.Fatal("EntityRole MUST be nil")
	}

	entityRoleFindWithDeleted, err := store.EntityRoleList(context.Background(), NewEntityRoleQuery().
		SetID(entityRole.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityRoleFindWithDeleted) != 0 {
		t.Fatal("EntityRole MUST be nil")
	}
}

func TestStoreEntityRoleDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityRoleDeleteByID(context.Background(), entityRole.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	entityRoleFound, err := store.EntityRoleFindByID(context.Background(), entityRole.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityRoleFound != nil {
		t.Fatal("EntityRole MUST be nil")
	}

	entityRoleFindWithDeleted, err := store.EntityRoleList(context.Background(), NewEntityRoleQuery().
		SetID(entityRole.ID()).
		SetSoftDeletedIncluded(true))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityRoleFindWithDeleted) != 0 {
		t.Fatal("EntityRole MUST NOT be found")
	}
}

func TestStoreEntityRoleFindByEntityAndRole(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = entityRole.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityRoleCreate(database.Context(context.Background(), store.DB()), entityRole)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	entityRoleFound, errFind := store.EntityRoleFindByEntityAndRole(database.Context(context.Background(), store.DB()), entityRole.EntityType(), entityRole.EntityID(), entityRole.RoleID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityRoleFound == nil {
		t.Fatal("EntityRole MUST NOT be nil")
	}

	if entityRoleFound.ID() != entityRole.ID() {
		t.Fatal("IDs do not match")
	}

	if entityRoleFound.EntityID() != entityRole.EntityID() {
		t.Fatal("EntityIDs do not match")
	}

	if entityRoleFound.EntityType() != entityRole.EntityType() {
		t.Fatal("EntityTypes do not match")
	}

	if entityRoleFound.RoleID() != entityRole.RoleID() {
		t.Fatal("RoleIDs do not match")
	}

	if entityRoleFound.Meta("education_1") != entityRole.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if entityRoleFound.Meta("education_2") != entityRole.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if entityRoleFound.Meta("education_3") != entityRole.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreEntityRoleFindByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = entityRole.SetMetas(map[string]string{
		"education_1": "Education 1",
		"education_2": "Education 2",
		"education_3": "Education 3",
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	ctx := database.Context(context.Background(), store.DB())
	err = store.EntityRoleCreate(ctx, entityRole)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	entityRoleFound, errFind := store.EntityRoleFindByID(ctx, entityRole.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityRoleFound == nil {
		t.Fatal("EntityRole MUST NOT be nil")
	}

	if entityRoleFound.ID() != entityRole.ID() {
		t.Fatal("IDs do not match")
	}

	if entityRoleFound.EntityID() != entityRole.EntityID() {
		t.Fatal("EntityIDs do not match")
	}

	if entityRoleFound.EntityType() != entityRole.EntityType() {
		t.Fatal("EntityTypes do not match")
	}

	if entityRoleFound.RoleID() != entityRole.RoleID() {
		t.Fatal("RoleIDs do not match")
	}

	if entityRoleFound.Meta("education_1") != entityRole.Meta("education_1") {
		t.Fatal("Metas do not match")
	}

	if entityRoleFound.Meta("education_2") != entityRole.Meta("education_2") {
		t.Fatal("Metas do not match")
	}

	if entityRoleFound.Meta("education_3") != entityRole.Meta("education_3") {
		t.Fatal("Metas do not match")
	}
}

func TestStoreEntityRoleList(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole1 := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	entityRole2 := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_02").
		SetRoleID("ROLE_02")

	entityRoles := []EntityRoleInterface{
		entityRole1,
		entityRole2,
	}

	for _, entityRole := range entityRoles {
		err = store.EntityRoleCreate(context.Background(), entityRole)
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	list1, err := store.EntityRoleList(context.Background(), NewEntityRoleQuery().SetRoleID("ROLE_01"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(list1) != 1 {
		t.Fatal("unexpected list length:", len(list1))
	}

	list2, err := store.EntityRoleList(context.Background(), NewEntityRoleQuery().SetEntityType("USER").SetEntityID("USER_02"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(list2) != 1 {
		t.Fatal("unexpected list length:", len(list2))
	}
}

func TestStoreEntityRoleSoftDelete(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityRoleSoftDelete(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityRole.SoftDeletedAt() == sb.MAX_DATETIME {
		t.Fatal("EntityRole MUST be soft deleted")
	}

	entityRoleFound, errFind := store.EntityRoleFindByID(context.Background(), entityRole.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityRoleFound != nil {
		t.Fatal("EntityRole MUST be soft deleted, so MUST be nil")
	}

	entityRoleFindWithDeleted, err := store.EntityRoleList(context.Background(), NewEntityRoleQuery().
		SetSoftDeletedIncluded(true).
		SetID(entityRole.ID()).
		SetLimit(1))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityRoleFindWithDeleted) == 0 {
		t.Fatal("EntityRole MUST be soft deleted")
	}

	if strings.Contains(entityRoleFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("EntityRole MUST be soft deleted", entityRole.SoftDeletedAt())
	}

	if !entityRoleFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("EntityRole MUST be soft deleted")
	}
}

func TestStoreEntityRoleSoftDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	defer func() {
		if err := store.DB().Close(); err != nil {
			t.Fatal(err)
		}
	}()

	entityRole := NewEntityRole().
		SetEntityType("USER").
		SetEntityID("USER_01").
		SetRoleID("ROLE_01")

	err = store.EntityRoleCreate(context.Background(), entityRole)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = store.EntityRoleSoftDeleteByID(context.Background(), entityRole.ID())

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if entityRole.SoftDeletedAt() != sb.MAX_DATETIME {
		t.Fatal("EntityRole MUST NOT be soft deleted, as it was soft deleted by ID")
	}

	entityRoleFound, errFind := store.EntityRoleFindByID(context.Background(), entityRole.ID())

	if errFind != nil {
		t.Fatal("unexpected error:", errFind)
	}

	if entityRoleFound != nil {
		t.Fatal("EntityRole MUST be nil")
	}
	query := NewEntityRoleQuery().
		SetSoftDeletedIncluded(true).
		SetID(entityRole.ID()).
		SetLimit(1)

	entityRoleFindWithDeleted, err := store.EntityRoleList(context.Background(), query)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if len(entityRoleFindWithDeleted) == 0 {
		t.Fatal("EntityRole MUST be soft deleted")
	}

	if strings.Contains(entityRoleFindWithDeleted[0].SoftDeletedAt(), sb.MAX_DATETIME) {
		t.Fatal("EntityRole MUST be soft deleted", entityRole.SoftDeletedAt())
	}

	if !entityRoleFindWithDeleted[0].IsSoftDeleted() {
		t.Fatal("EntityRole MUST be soft deleted")
	}
}
