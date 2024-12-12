package rolestore

import (
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/dataobject"
	"github.com/gouniverse/maputils"
	"github.com/gouniverse/sb"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
)

// == CLASS ===================================================================

type entityRole struct {
	dataobject.DataObject
}

var _ EntityRoleInterface = (*entityRole)(nil)

// == CONSTRUCTORS ============================================================

func NewEntityRole() EntityRoleInterface {
	o := (&entityRole{}).
		SetID(uid.HumanUid()).
		SetMemo("").
		SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC)).
		SetSoftDeletedAt(sb.MAX_DATETIME)

	err := o.SetMetas(map[string]string{})

	if err != nil {
		return o
	}

	return o
}

func NewEntityRoleFromExistingData(data map[string]string) EntityRoleInterface {
	o := &entityRole{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func (o *entityRole) IsSoftDeleted() bool {
	return o.SoftDeletedAtCarbon().Compare("<", carbon.Now(carbon.UTC))
}

// == SETTERS AND GETTERS =====================================================

func (o *entityRole) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *entityRole) CreatedAtCarbon() carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *entityRole) SetCreatedAt(createdAt string) EntityRoleInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *entityRole) EntityType() string {
	return o.Get(COLUMN_ENTITY_TYPE)
}

func (o *entityRole) SetEntityType(entityType string) EntityRoleInterface {
	o.Set(COLUMN_ENTITY_TYPE, entityType)
	return o
}

func (o *entityRole) EntityID() string {
	return o.Get(COLUMN_ENTITY_ID)
}

func (o *entityRole) SetEntityID(entityID string) EntityRoleInterface {
	o.Set(COLUMN_ENTITY_ID, entityID)
	return o
}

func (o *entityRole) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *entityRole) SetID(id string) EntityRoleInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *entityRole) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *entityRole) SetMemo(memo string) EntityRoleInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *entityRole) Metas() (map[string]string, error) {
	metasStr := o.Get(COLUMN_METAS)

	if metasStr == "" {
		metasStr = "{}"
	}

	metasJson, errJson := utils.FromJSON(metasStr, map[string]string{})
	if errJson != nil {
		return map[string]string{}, errJson
	}

	return maputils.MapStringAnyToMapStringString(metasJson.(map[string]any)), nil
}

func (o *entityRole) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *entityRole) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *entityRole) SetMetas(metas map[string]string) error {
	mapString, err := utils.ToJSON(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, mapString)
	return nil
}

func (o *entityRole) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *entityRole) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *entityRole) SoftDeletedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *entityRole) SetSoftDeletedAt(deletedAt string) EntityRoleInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *entityRole) RoleID() string {
	return o.Get(COLUMN_ROLE_ID)
}

func (o *entityRole) SetRoleID(roleID string) EntityRoleInterface {
	o.Set(COLUMN_ROLE_ID, roleID)
	return o
}

func (o *entityRole) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *entityRole) UpdatedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *entityRole) SetUpdatedAt(updatedAt string) EntityRoleInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
