package rolestore

import "errors"

type EntityRoleQueryInterface interface {
	Validate() error

	Columns() []string
	SetColumns(columns []string) EntityRoleQueryInterface

	HasCountOnly() bool
	IsCountOnly() bool
	SetCountOnly(countOnly bool) EntityRoleQueryInterface

	HasCreatedAtGte() bool
	CreatedAtGte() string
	SetCreatedAtGte(createdAtGte string) EntityRoleQueryInterface

	HasCreatedAtLte() bool
	CreatedAtLte() string
	SetCreatedAtLte(createdAtLte string) EntityRoleQueryInterface

	HasEntityID() bool
	EntityID() string
	SetEntityID(entityID string) EntityRoleQueryInterface

	HasEntityType() bool
	EntityType() string
	SetEntityType(entityType string) EntityRoleQueryInterface

	HasID() bool
	ID() string
	SetID(id string) EntityRoleQueryInterface

	HasIDIn() bool
	IDIn() []string
	SetIDIn(idIn []string) EntityRoleQueryInterface

	HasLimit() bool
	Limit() int
	SetLimit(limit int) EntityRoleQueryInterface

	HasOffset() bool
	Offset() int
	SetOffset(offset int) EntityRoleQueryInterface

	HasOrderBy() bool
	OrderBy() string
	SetOrderBy(orderBy string) EntityRoleQueryInterface

	HasRoleID() bool
	RoleID() string
	SetRoleID(roleID string) EntityRoleQueryInterface

	HasSortDirection() bool
	SortDirection() string
	SetSortDirection(sortDirection string) EntityRoleQueryInterface

	HasSoftDeletedIncluded() bool
	SoftDeletedIncluded() bool
	SetSoftDeletedIncluded(softDeletedIncluded bool) EntityRoleQueryInterface

	hasProperty(name string) bool
}

func NewEntityRoleQuery() EntityRoleQueryInterface {
	return &roleEntityQueryImplementation{
		properties: make(map[string]any),
	}
}

type roleEntityQueryImplementation struct {
	properties map[string]any
}

func (c *roleEntityQueryImplementation) Validate() error {
	if c.HasCreatedAtGte() && c.CreatedAtGte() == "" {
		return errors.New("role query. created_at_gte cannot be empty")
	}

	if c.HasCreatedAtLte() && c.CreatedAtLte() == "" {
		return errors.New("role query. created_at_lte cannot be empty")
	}

	if c.HasEntityID() && c.EntityID() == "" {
		return errors.New("role query. entity_id cannot be empty")
	}

	if c.HasEntityType() && c.EntityType() == "" {
		return errors.New("role query. entity_type cannot be empty")
	}

	if c.HasID() && c.ID() == "" {
		return errors.New("role query. id cannot be empty")
	}

	if c.HasIDIn() && len(c.IDIn()) == 0 {
		return errors.New("role query. id_in cannot be empty")
	}

	if c.HasOrderBy() && c.OrderBy() == "" {
		return errors.New("role query. order_by cannot be empty")
	}

	if c.HasSortDirection() && c.SortDirection() == "" {
		return errors.New("role query. sort_direction cannot be empty")
	}

	if c.HasLimit() && c.Limit() <= 0 {
		return errors.New("role query. limit must be greater than 0")
	}

	if c.HasOffset() && c.Offset() < 0 {
		return errors.New("role query. offset must be greater than or equal to 0")
	}

	return nil
}

func (c *roleEntityQueryImplementation) Columns() []string {
	if !c.hasProperty("columns") {
		return []string{}
	}

	return c.properties["columns"].([]string)
}

func (c *roleEntityQueryImplementation) SetColumns(columns []string) EntityRoleQueryInterface {
	c.properties["columns"] = columns

	return c
}

func (c *roleEntityQueryImplementation) HasCountOnly() bool {
	return c.hasProperty("count_only")
}

func (c *roleEntityQueryImplementation) IsCountOnly() bool {
	if !c.HasCountOnly() {
		return false
	}

	return c.properties["count_only"].(bool)
}

func (c *roleEntityQueryImplementation) SetCountOnly(countOnly bool) EntityRoleQueryInterface {
	c.properties["count_only"] = countOnly

	return c
}

func (c *roleEntityQueryImplementation) HasCreatedAtGte() bool {
	return c.hasProperty("created_at_gte")
}

func (c *roleEntityQueryImplementation) CreatedAtGte() string {
	if !c.HasCreatedAtGte() {
		return ""
	}

	return c.properties["created_at_gte"].(string)
}

func (c *roleEntityQueryImplementation) SetCreatedAtGte(createdAtGte string) EntityRoleQueryInterface {
	c.properties["created_at_gte"] = createdAtGte

	return c
}

func (c *roleEntityQueryImplementation) HasCreatedAtLte() bool {
	return c.hasProperty("created_at_lte")
}

func (c *roleEntityQueryImplementation) CreatedAtLte() string {
	if !c.HasCreatedAtLte() {
		return ""
	}

	return c.properties["created_at_lte"].(string)
}

func (c *roleEntityQueryImplementation) SetCreatedAtLte(createdAtLte string) EntityRoleQueryInterface {
	c.properties["created_at_lte"] = createdAtLte

	return c
}

func (c *roleEntityQueryImplementation) HasEntityType() bool {
	return c.hasProperty("entity_type")
}

func (c *roleEntityQueryImplementation) EntityType() string {
	if !c.HasEntityType() {
		return ""
	}

	return c.properties["entity_type"].(string)
}

func (c *roleEntityQueryImplementation) SetEntityType(entityType string) EntityRoleQueryInterface {
	c.properties["entity_type"] = entityType

	return c
}

func (c *roleEntityQueryImplementation) HasEntityID() bool {
	return c.hasProperty("entity_id")
}

func (c *roleEntityQueryImplementation) EntityID() string {
	if !c.HasEntityID() {
		return ""
	}

	return c.properties["entity_id"].(string)
}

func (c *roleEntityQueryImplementation) SetEntityID(entityID string) EntityRoleQueryInterface {
	c.properties["entity_id"] = entityID

	return c
}

func (c *roleEntityQueryImplementation) HasID() bool {
	return c.hasProperty("id")
}

func (c *roleEntityQueryImplementation) ID() string {
	if !c.HasID() {
		return ""
	}

	return c.properties["id"].(string)
}

func (c *roleEntityQueryImplementation) SetID(id string) EntityRoleQueryInterface {
	c.properties["id"] = id

	return c
}

func (c *roleEntityQueryImplementation) HasIDIn() bool {
	return c.hasProperty("id_in")
}

func (c *roleEntityQueryImplementation) IDIn() []string {
	if !c.HasIDIn() {
		return []string{}
	}

	return c.properties["id_in"].([]string)
}

func (c *roleEntityQueryImplementation) SetIDIn(idIn []string) EntityRoleQueryInterface {
	c.properties["id_in"] = idIn

	return c
}

func (c *roleEntityQueryImplementation) HasLimit() bool {
	return c.hasProperty("limit")
}

func (c *roleEntityQueryImplementation) Limit() int {
	if !c.HasLimit() {
		return 0
	}

	return c.properties["limit"].(int)
}

func (c *roleEntityQueryImplementation) SetLimit(limit int) EntityRoleQueryInterface {
	c.properties["limit"] = limit

	return c
}

func (c *roleEntityQueryImplementation) HasOffset() bool {
	return c.hasProperty("offset")
}

func (c *roleEntityQueryImplementation) Offset() int {
	if !c.HasOffset() {
		return 0
	}

	return c.properties["offset"].(int)
}

func (c *roleEntityQueryImplementation) SetOffset(offset int) EntityRoleQueryInterface {
	c.properties["offset"] = offset

	return c
}

func (c *roleEntityQueryImplementation) HasOrderBy() bool {
	return c.hasProperty("order_by")
}

func (c *roleEntityQueryImplementation) OrderBy() string {
	if !c.HasOrderBy() {
		return ""
	}

	return c.properties["order_by"].(string)
}

func (c *roleEntityQueryImplementation) SetOrderBy(orderBy string) EntityRoleQueryInterface {
	c.properties["order_by"] = orderBy

	return c
}

func (c *roleEntityQueryImplementation) HasRoleID() bool {
	return c.hasProperty("role_id")
}

func (c *roleEntityQueryImplementation) RoleID() string {
	if !c.HasRoleID() {
		return ""
	}

	return c.properties["role_id"].(string)
}

func (c *roleEntityQueryImplementation) SetRoleID(roleID string) EntityRoleQueryInterface {
	c.properties["role_id"] = roleID

	return c
}

func (c *roleEntityQueryImplementation) HasSortDirection() bool {
	return c.hasProperty("sort_direction")
}

func (c *roleEntityQueryImplementation) SortDirection() string {
	if !c.HasSortDirection() {
		return ""
	}

	return c.properties["sort_direction"].(string)
}

func (c *roleEntityQueryImplementation) SetSortDirection(sortDirection string) EntityRoleQueryInterface {
	c.properties["sort_direction"] = sortDirection

	return c
}

func (c *roleEntityQueryImplementation) HasSoftDeletedIncluded() bool {
	return c.hasProperty("soft_deleted_included")
}

func (c *roleEntityQueryImplementation) SoftDeletedIncluded() bool {
	if !c.HasSoftDeletedIncluded() {
		return false
	}

	return c.properties["soft_deleted_included"].(bool)
}

func (c *roleEntityQueryImplementation) SetSoftDeletedIncluded(softDeletedIncluded bool) EntityRoleQueryInterface {
	c.properties["soft_deleted_included"] = softDeletedIncluded

	return c
}

func (c *roleEntityQueryImplementation) HasTitleLike() bool {
	return c.hasProperty("title_like")
}

func (c *roleEntityQueryImplementation) TitleLike() string {
	if !c.HasTitleLike() {
		return ""
	}

	return c.properties["title_like"].(string)
}

func (c *roleEntityQueryImplementation) SetTitleLike(titleLike string) EntityRoleQueryInterface {
	c.properties["title_like"] = titleLike

	return c
}

func (c *roleEntityQueryImplementation) hasProperty(name string) bool {
	_, ok := c.properties[name]
	return ok
}
