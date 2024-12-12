package rolestore

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/dromara/carbon/v2"
	"github.com/gouniverse/base/database"
	"github.com/gouniverse/sb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func (store *store) EntityRoleCount(ctx context.Context, options EntityRoleQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q, _, err := store.entityRoleSelectQuery(options)

	sqlStr, params, errSql := q.Prepared(true).
		Limit(1).
		Select(goqu.COUNT(goqu.Star()).As("count")).
		ToSQL()

	if errSql != nil {
		return -1, nil
	}

	store.logSql("select", sqlStr, params...)

	mapped, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, params...)
	if err != nil {
		return -1, err
	}

	if len(mapped) < 1 {
		return -1, nil
	}

	countStr := mapped[0]["count"]

	i, err := strconv.ParseInt(countStr, 10, 64)

	if err != nil {
		return -1, err

	}

	return i, nil
}

func (store *store) EntityRoleCreate(ctx context.Context, entityRole EntityRoleInterface) error {
	if entityRole == nil {
		return errors.New("rolestore > EntityRoleCreate. entityRole is nil")
	}

	if entityRole.RoleID() == "" {
		return errors.New("rolestore > EntityRoleCreate. entityRole roleID is empty")
	}

	if entityRole.EntityID() == "" {
		return errors.New("rolestore > EntityRoleCreate. entityRole entityID is empty")
	}

	if entityRole.EntityType() == "" {
		return errors.New("rolestore > EntityRoleCreate. entityRole entityType is empty")
	}

	entityRoleExists, err := store.EntityRoleFindByEntityAndRole(
		ctx,
		entityRole.EntityType(),
		entityRole.EntityID(),
		entityRole.RoleID(),
	)

	if err != nil {
		return err
	}

	if entityRoleExists != nil {
		return errors.New("rolestore > EntityRoleCreate. entityRole with the same entityType-entityID-roleID combination already exists")
	}

	entityRole.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	entityRole.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := entityRole.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.entityRoleTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("insert", sqlStr, params...)

	if store.db == nil {
		return errors.New("entityRolestore: database is nil")
	}

	_, err = database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	if err != nil {
		return err
	}

	entityRole.MarkAsNotDirty()

	return nil
}

func (store *store) EntityRoleDelete(ctx context.Context, entityRole EntityRoleInterface) error {
	if entityRole == nil {
		return errors.New("entityRole is nil")
	}

	return store.EntityRoleDeleteByID(ctx, entityRole.ID())
}

func (store *store) EntityRoleDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("entityRole id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.entityRoleTableName).
		Prepared(true).
		Where(goqu.C(COLUMN_ID).Eq(id)).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("delete", sqlStr, params...)

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	return err
}

func (store *store) EntityRoleFindByEntityAndRole(
	ctx context.Context,
	entityType string,
	entityID string,
	roleID string,
) (entityRole EntityRoleInterface, err error) {
	if entityType == "" {
		return nil, errors.New("EntityRoleFindByEntityAndRole entityType is empty")
	}

	if entityID == "" {
		return nil, errors.New("EntityRoleFindByEntityAndRole entityID is empty")
	}

	if roleID == "" {
		return nil, errors.New("EntityRoleFindByEntityAndRole roleID is empty")
	}

	query := NewEntityRoleQuery().
		SetEntityType(entityType).
		SetEntityID(entityID).
		SetRoleID(roleID).
		SetLimit(1)

	list, err := store.EntityRoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) EntityRoleFindByID(ctx context.Context, id string) (entityRole EntityRoleInterface, err error) {
	if id == "" {
		return nil, errors.New("entityRole id is empty")
	}

	query := NewEntityRoleQuery().SetID(id).SetLimit(1)

	list, err := store.EntityRoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) EntityRoleList(ctx context.Context, query EntityRoleQueryInterface) ([]EntityRoleInterface, error) {
	if query == nil {
		return []EntityRoleInterface{}, errors.New("at entityRole list > entityRole query is nil")
	}

	q, columns, err := store.entityRoleSelectQuery(query)

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	if errSql != nil {
		return []EntityRoleInterface{}, nil
	}

	store.logSql("select", sqlStr, sqlParams...)

	if store.db == nil {
		return []EntityRoleInterface{}, errors.New("entityRolestore: database is nil")
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []EntityRoleInterface{}, err
	}

	list := []EntityRoleInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewEntityRoleFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) EntityRoleSoftDelete(ctx context.Context, entityRole EntityRoleInterface) error {
	if entityRole == nil {
		return errors.New("at entityRole soft delete > entityRole is nil")
	}

	entityRole.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.EntityRoleUpdate(ctx, entityRole)
}

func (store *store) EntityRoleSoftDeleteByID(ctx context.Context, id string) error {
	entityRole, err := store.EntityRoleFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.EntityRoleSoftDelete(ctx, entityRole)
}

func (store *store) EntityRoleUpdate(ctx context.Context, entityRole EntityRoleInterface) error {
	if entityRole == nil {
		return errors.New("at entityRole update > entityRole is nil")
	}

	entityRole.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := entityRole.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.entityRoleTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(entityRole.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("update", sqlStr, params...)

	if store.db == nil {
		return errors.New("entityRolestore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	entityRole.MarkAsNotDirty()

	return err
}

func (store *store) entityRoleSelectQuery(options EntityRoleQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, nil, errors.New("entityRole options is nil")
	}

	if err := options.Validate(); err != nil {
		return nil, nil, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.entityRoleTableName)

	if options.HasEntityID() {
		q = q.Where(goqu.C(COLUMN_ENTITY_ID).Eq(options.EntityID()))
	}

	if options.HasEntityType() {
		q = q.Where(goqu.C(COLUMN_ENTITY_TYPE).Eq(options.EntityType()))
	}

	if options.HasID() {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if options.HasIDIn() {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.HasRoleID() {
		q = q.Where(goqu.C(COLUMN_ROLE_ID).Eq(options.RoleID()))
	}

	if options.HasCreatedAtGte() && options.HasCreatedAtLte() {
		q = q.Where(
			goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()),
			goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()),
		)
	} else if options.HasCreatedAtGte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Gte(options.CreatedAtGte()))
	} else if options.HasCreatedAtLte() {
		q = q.Where(goqu.C(COLUMN_CREATED_AT).Lte(options.CreatedAtLte()))
	}

	if !options.IsCountOnly() {
		if options.HasLimit() {
			q = q.Limit(cast.ToUint(options.Limit()))
		}

		if options.HasOffset() {
			q = q.Offset(cast.ToUint(options.Offset()))
		}
	}

	if options.HasOrderBy() {
		sort := lo.Ternary(options.HasSortDirection(), options.SortDirection(), sb.DESC)
		if strings.EqualFold(sort, sb.ASC) {
			q = q.Order(goqu.I(options.OrderBy()).Asc())
		} else {
			q = q.Order(goqu.I(options.OrderBy()).Desc())
		}
	}

	columns = []any{}

	for _, column := range options.Columns() {
		columns = append(columns, column)
	}

	if options.SoftDeletedIncluded() {
		return q, columns, nil // soft deleted entityRoles requested specifically
	}

	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted), columns, nil
}
