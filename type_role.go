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

type role struct {
	dataobject.DataObject
}

var _ RoleInterface = (*role)(nil)

// == CONSTRUCTORS ============================================================

func NewRole() RoleInterface {
	o := (&role{}).
		SetID(uid.HumanUid()).
		SetStatus(ROLE_STATUS_INACTIVE).
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

func NewRoleFromExistingData(data map[string]string) RoleInterface {
	o := &role{}
	o.Hydrate(data)
	return o
}

// == METHODS =================================================================

func (o *role) IsActive() bool {
	return o.Status() == ROLE_STATUS_ACTIVE
}

func (o *role) IsSoftDeleted() bool {
	return o.SoftDeletedAtCarbon().Compare("<", carbon.Now(carbon.UTC))
}

func (o *role) IsInactive() bool {
	return o.Status() == ROLE_STATUS_INACTIVE
}

// == SETTERS AND GETTERS =====================================================

func (o *role) CreatedAt() string {
	return o.Get(COLUMN_CREATED_AT)
}

func (o *role) CreatedAtCarbon() carbon.Carbon {
	return carbon.Parse(o.CreatedAt(), carbon.UTC)
}

func (o *role) SetCreatedAt(createdAt string) RoleInterface {
	o.Set(COLUMN_CREATED_AT, createdAt)
	return o
}

func (o *role) Handle() string {
	return o.Get(COLUMN_HANDLE)
}

func (o *role) SetHandle(handle string) RoleInterface {
	o.Set(COLUMN_HANDLE, handle)
	return o
}

func (o *role) ID() string {
	return o.Get(COLUMN_ID)
}

func (o *role) SetID(id string) RoleInterface {
	o.Set(COLUMN_ID, id)
	return o
}

func (o *role) Memo() string {
	return o.Get(COLUMN_MEMO)
}

func (o *role) SetMemo(memo string) RoleInterface {
	o.Set(COLUMN_MEMO, memo)
	return o
}

func (o *role) Metas() (map[string]string, error) {
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

func (o *role) Meta(name string) string {
	metas, err := o.Metas()

	if err != nil {
		return ""
	}

	if value, exists := metas[name]; exists {
		return value
	}

	return ""
}

func (o *role) SetMeta(name, value string) error {
	return o.UpsertMetas(map[string]string{name: value})
}

// SetMetas stores metas as json string
// Warning: it overwrites any existing metas
func (o *role) SetMetas(metas map[string]string) error {
	mapString, err := utils.ToJSON(metas)
	if err != nil {
		return err
	}
	o.Set(COLUMN_METAS, mapString)
	return nil
}

func (o *role) UpsertMetas(metas map[string]string) error {
	currentMetas, err := o.Metas()

	if err != nil {
		return err
	}

	for k, v := range metas {
		currentMetas[k] = v
	}

	return o.SetMetas(currentMetas)
}

func (o *role) SoftDeletedAt() string {
	return o.Get(COLUMN_SOFT_DELETED_AT)
}

func (o *role) SoftDeletedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.SoftDeletedAt(), carbon.UTC)
}

func (o *role) SetSoftDeletedAt(deletedAt string) RoleInterface {
	o.Set(COLUMN_SOFT_DELETED_AT, deletedAt)
	return o
}

func (o *role) Status() string {
	return o.Get(COLUMN_STATUS)
}

func (o *role) SetStatus(status string) RoleInterface {
	o.Set(COLUMN_STATUS, status)
	return o
}

func (o *role) Title() string {
	return o.Get(COLUMN_TITLE)
}

func (o *role) SetTitle(title string) RoleInterface {
	o.Set(COLUMN_TITLE, title)
	return o
}

func (o *role) UpdatedAt() string {
	return o.Get(COLUMN_UPDATED_AT)
}

func (o *role) UpdatedAtCarbon() carbon.Carbon {
	return carbon.NewCarbon().Parse(o.Get(COLUMN_UPDATED_AT), carbon.UTC)
}

func (o *role) SetUpdatedAt(updatedAt string) RoleInterface {
	o.Set(COLUMN_UPDATED_AT, updatedAt)
	return o
}
