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

func (store *store) RoleCount(ctx context.Context, options RoleQueryInterface) (int64, error) {
	options.SetCountOnly(true)

	q, _, err := store.roleSelectQuery(options)

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

func (store *store) RoleCreate(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("role is nil")
	}

	role.SetCreatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))
	role.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	data := role.Data()

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Insert(store.roleTableName).
		Prepared(true).
		Rows(data).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("insert", sqlStr, params...)

	if store.db == nil {
		return errors.New("rolestore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	if err != nil {
		return err
	}

	role.MarkAsNotDirty()

	return nil
}

func (store *store) RoleDelete(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("role is nil")
	}

	return store.RoleDeleteByID(ctx, role.ID())
}

func (store *store) RoleDeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("role id is empty")
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Delete(store.roleTableName).
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

func (store *store) RoleFindByHandle(ctx context.Context, handle string) (role RoleInterface, err error) {
	if handle == "" {
		return nil, errors.New("role handle is empty")
	}

	query := NewRoleQuery().SetHandle(handle).SetLimit(1)

	list, err := store.RoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) RoleFindByID(ctx context.Context, id string) (role RoleInterface, err error) {
	if id == "" {
		return nil, errors.New("role id is empty")
	}

	query := NewRoleQuery().SetID(id).SetLimit(1)

	list, err := store.RoleList(ctx, query)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return nil, nil
}

func (store *store) RoleList(ctx context.Context, query RoleQueryInterface) ([]RoleInterface, error) {
	if query == nil {
		return []RoleInterface{}, errors.New("at role list > role query is nil")
	}

	q, columns, err := store.roleSelectQuery(query)

	sqlStr, sqlParams, errSql := q.Prepared(true).Select(columns...).ToSQL()

	if errSql != nil {
		return []RoleInterface{}, nil
	}

	store.logSql("select", sqlStr, sqlParams...)

	if store.db == nil {
		return []RoleInterface{}, errors.New("rolestore: database is nil")
	}

	modelMaps, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, sqlParams...)

	if err != nil {
		return []RoleInterface{}, err
	}

	list := []RoleInterface{}

	lo.ForEach(modelMaps, func(modelMap map[string]string, index int) {
		model := NewRoleFromExistingData(modelMap)
		list = append(list, model)
	})

	return list, nil
}

func (store *store) RoleSoftDelete(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("at role soft delete > role is nil")
	}

	role.SetSoftDeletedAt(carbon.Now(carbon.UTC).ToDateTimeString(carbon.UTC))

	return store.RoleUpdate(ctx, role)
}

func (store *store) RoleSoftDeleteByID(ctx context.Context, id string) error {
	role, err := store.RoleFindByID(ctx, id)

	if err != nil {
		return err
	}

	return store.RoleSoftDelete(ctx, role)
}

func (store *store) RoleUpdate(ctx context.Context, role RoleInterface) error {
	if role == nil {
		return errors.New("at role update > role is nil")
	}

	role.SetUpdatedAt(carbon.Now(carbon.UTC).ToDateTimeString())

	dataChanged := role.DataChanged()

	delete(dataChanged, COLUMN_ID) // ID is not updateable

	if len(dataChanged) < 1 {
		return nil
	}

	sqlStr, params, errSql := goqu.Dialect(store.dbDriverName).
		Update(store.roleTableName).
		Prepared(true).
		Set(dataChanged).
		Where(goqu.C(COLUMN_ID).Eq(role.ID())).
		ToSQL()

	if errSql != nil {
		return errSql
	}

	store.logSql("update", sqlStr, params...)

	if store.db == nil {
		return errors.New("rolestore: database is nil")
	}

	_, err := database.Execute(store.toQuerableContext(ctx), sqlStr, params...)

	role.MarkAsNotDirty()

	return err
}

func (store *store) roleSelectQuery(options RoleQueryInterface) (selectDataset *goqu.SelectDataset, columns []any, err error) {
	if options == nil {
		return nil, nil, errors.New("role options is nil")
	}

	if err := options.Validate(); err != nil {
		return nil, nil, err
	}

	q := goqu.Dialect(store.dbDriverName).From(store.roleTableName)

	if options.HasID() {
		q = q.Where(goqu.C(COLUMN_ID).Eq(options.ID()))
	}

	if options.HasIDIn() {
		q = q.Where(goqu.C(COLUMN_ID).In(options.IDIn()))
	}

	if options.HasStatus() {
		q = q.Where(goqu.C(COLUMN_STATUS).Eq(options.Status()))
	}

	if options.HasStatusIn() {
		q = q.Where(goqu.C(COLUMN_STATUS).In(options.StatusIn()))
	}

	if options.HasHandle() {
		q = q.Where(goqu.C(COLUMN_HANDLE).Eq(options.Handle()))
	}

	if options.HasTitleLike() {
		q = q.Where(goqu.C(COLUMN_TITLE).ILike(`%` + options.TitleLike() + `%`))
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
		return q, columns, nil // soft deleted roles requested specifically
	}

	softDeleted := goqu.C(COLUMN_SOFT_DELETED_AT).
		Gt(carbon.Now(carbon.UTC).ToDateTimeString())

	return q.Where(softDeleted), columns, nil
}
