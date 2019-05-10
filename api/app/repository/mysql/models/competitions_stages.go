// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// CompetitionsStage is an object representing the database table.
type CompetitionsStage struct {
	ID            int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CompetitionID int       `boil:"competition_id" json:"competition_id" toml:"competition_id" yaml:"competition_id"`
	StageID       int       `boil:"stage_id" json:"stage_id" toml:"stage_id" yaml:"stage_id"`
	CreatedAt     null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *competitionsStageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L competitionsStageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompetitionsStageColumns = struct {
	ID            string
	CompetitionID string
	StageID       string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	CompetitionID: "competition_id",
	StageID:       "stage_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// Generated where

var CompetitionsStageWhere = struct {
	ID            whereHelperint
	CompetitionID whereHelperint
	StageID       whereHelperint
	CreatedAt     whereHelpernull_Time
	UpdatedAt     whereHelpernull_Time
}{
	ID:            whereHelperint{field: `id`},
	CompetitionID: whereHelperint{field: `competition_id`},
	StageID:       whereHelperint{field: `stage_id`},
	CreatedAt:     whereHelpernull_Time{field: `created_at`},
	UpdatedAt:     whereHelpernull_Time{field: `updated_at`},
}

// CompetitionsStageRels is where relationship names are stored.
var CompetitionsStageRels = struct {
}{}

// competitionsStageR is where relationships are stored.
type competitionsStageR struct {
}

// NewStruct creates a new relationship struct
func (*competitionsStageR) NewStruct() *competitionsStageR {
	return &competitionsStageR{}
}

// competitionsStageL is where Load methods for each relationship are stored.
type competitionsStageL struct{}

var (
	competitionsStageColumns               = []string{"id", "competition_id", "stage_id", "created_at", "updated_at"}
	competitionsStageColumnsWithoutDefault = []string{"competition_id", "stage_id", "created_at", "updated_at"}
	competitionsStageColumnsWithDefault    = []string{"id"}
	competitionsStagePrimaryKeyColumns     = []string{"id"}
)

type (
	// CompetitionsStageSlice is an alias for a slice of pointers to CompetitionsStage.
	// This should generally be used opposed to []CompetitionsStage.
	CompetitionsStageSlice []*CompetitionsStage
	// CompetitionsStageHook is the signature for custom CompetitionsStage hook methods
	CompetitionsStageHook func(context.Context, boil.ContextExecutor, *CompetitionsStage) error

	competitionsStageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	competitionsStageType                 = reflect.TypeOf(&CompetitionsStage{})
	competitionsStageMapping              = queries.MakeStructMapping(competitionsStageType)
	competitionsStagePrimaryKeyMapping, _ = queries.BindMapping(competitionsStageType, competitionsStageMapping, competitionsStagePrimaryKeyColumns)
	competitionsStageInsertCacheMut       sync.RWMutex
	competitionsStageInsertCache          = make(map[string]insertCache)
	competitionsStageUpdateCacheMut       sync.RWMutex
	competitionsStageUpdateCache          = make(map[string]updateCache)
	competitionsStageUpsertCacheMut       sync.RWMutex
	competitionsStageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var competitionsStageBeforeInsertHooks []CompetitionsStageHook
var competitionsStageBeforeUpdateHooks []CompetitionsStageHook
var competitionsStageBeforeDeleteHooks []CompetitionsStageHook
var competitionsStageBeforeUpsertHooks []CompetitionsStageHook

var competitionsStageAfterInsertHooks []CompetitionsStageHook
var competitionsStageAfterSelectHooks []CompetitionsStageHook
var competitionsStageAfterUpdateHooks []CompetitionsStageHook
var competitionsStageAfterDeleteHooks []CompetitionsStageHook
var competitionsStageAfterUpsertHooks []CompetitionsStageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CompetitionsStage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CompetitionsStage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CompetitionsStage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CompetitionsStage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CompetitionsStage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CompetitionsStage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CompetitionsStage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CompetitionsStage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CompetitionsStage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range competitionsStageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompetitionsStageHook registers your hook function for all future operations.
func AddCompetitionsStageHook(hookPoint boil.HookPoint, competitionsStageHook CompetitionsStageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		competitionsStageBeforeInsertHooks = append(competitionsStageBeforeInsertHooks, competitionsStageHook)
	case boil.BeforeUpdateHook:
		competitionsStageBeforeUpdateHooks = append(competitionsStageBeforeUpdateHooks, competitionsStageHook)
	case boil.BeforeDeleteHook:
		competitionsStageBeforeDeleteHooks = append(competitionsStageBeforeDeleteHooks, competitionsStageHook)
	case boil.BeforeUpsertHook:
		competitionsStageBeforeUpsertHooks = append(competitionsStageBeforeUpsertHooks, competitionsStageHook)
	case boil.AfterInsertHook:
		competitionsStageAfterInsertHooks = append(competitionsStageAfterInsertHooks, competitionsStageHook)
	case boil.AfterSelectHook:
		competitionsStageAfterSelectHooks = append(competitionsStageAfterSelectHooks, competitionsStageHook)
	case boil.AfterUpdateHook:
		competitionsStageAfterUpdateHooks = append(competitionsStageAfterUpdateHooks, competitionsStageHook)
	case boil.AfterDeleteHook:
		competitionsStageAfterDeleteHooks = append(competitionsStageAfterDeleteHooks, competitionsStageHook)
	case boil.AfterUpsertHook:
		competitionsStageAfterUpsertHooks = append(competitionsStageAfterUpsertHooks, competitionsStageHook)
	}
}

// One returns a single competitionsStage record from the query.
func (q competitionsStageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CompetitionsStage, error) {
	o := &CompetitionsStage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for competitions_stages")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CompetitionsStage records from the query.
func (q competitionsStageQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompetitionsStageSlice, error) {
	var o []*CompetitionsStage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CompetitionsStage slice")
	}

	if len(competitionsStageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CompetitionsStage records in the query.
func (q competitionsStageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count competitions_stages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q competitionsStageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if competitions_stages exists")
	}

	return count > 0, nil
}

// CompetitionsStages retrieves all the records using an executor.
func CompetitionsStages(mods ...qm.QueryMod) competitionsStageQuery {
	mods = append(mods, qm.From("`competitions_stages`"))
	return competitionsStageQuery{NewQuery(mods...)}
}

// FindCompetitionsStage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompetitionsStage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CompetitionsStage, error) {
	competitionsStageObj := &CompetitionsStage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `competitions_stages` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, competitionsStageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from competitions_stages")
	}

	return competitionsStageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CompetitionsStage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no competitions_stages provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(competitionsStageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	competitionsStageInsertCacheMut.RLock()
	cache, cached := competitionsStageInsertCache[key]
	competitionsStageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			competitionsStageColumns,
			competitionsStageColumnsWithDefault,
			competitionsStageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(competitionsStageType, competitionsStageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(competitionsStageType, competitionsStageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `competitions_stages` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `competitions_stages` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `competitions_stages` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, competitionsStagePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into competitions_stages")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == competitionsStageMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for competitions_stages")
	}

CacheNoHooks:
	if !cached {
		competitionsStageInsertCacheMut.Lock()
		competitionsStageInsertCache[key] = cache
		competitionsStageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CompetitionsStage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CompetitionsStage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	competitionsStageUpdateCacheMut.RLock()
	cache, cached := competitionsStageUpdateCache[key]
	competitionsStageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			competitionsStageColumns,
			competitionsStagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update competitions_stages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `competitions_stages` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, competitionsStagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(competitionsStageType, competitionsStageMapping, append(wl, competitionsStagePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update competitions_stages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for competitions_stages")
	}

	if !cached {
		competitionsStageUpdateCacheMut.Lock()
		competitionsStageUpdateCache[key] = cache
		competitionsStageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q competitionsStageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for competitions_stages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for competitions_stages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompetitionsStageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitionsStagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `competitions_stages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, competitionsStagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in competitionsStage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all competitionsStage")
	}
	return rowsAff, nil
}

var mySQLCompetitionsStageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CompetitionsStage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no competitions_stages provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(competitionsStageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCompetitionsStageUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	competitionsStageUpsertCacheMut.RLock()
	cache, cached := competitionsStageUpsertCache[key]
	competitionsStageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			competitionsStageColumns,
			competitionsStageColumnsWithDefault,
			competitionsStageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			competitionsStageColumns,
			competitionsStagePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert competitions_stages, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "competitions_stages", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `competitions_stages` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(competitionsStageType, competitionsStageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(competitionsStageType, competitionsStageMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for competitions_stages")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == competitionsStageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(competitionsStageType, competitionsStageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for competitions_stages")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for competitions_stages")
	}

CacheNoHooks:
	if !cached {
		competitionsStageUpsertCacheMut.Lock()
		competitionsStageUpsertCache[key] = cache
		competitionsStageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CompetitionsStage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CompetitionsStage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CompetitionsStage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), competitionsStagePrimaryKeyMapping)
	sql := "DELETE FROM `competitions_stages` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from competitions_stages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for competitions_stages")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q competitionsStageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no competitionsStageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from competitions_stages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for competitions_stages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompetitionsStageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CompetitionsStage slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(competitionsStageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitionsStagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `competitions_stages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, competitionsStagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from competitionsStage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for competitions_stages")
	}

	if len(competitionsStageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CompetitionsStage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompetitionsStage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompetitionsStageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompetitionsStageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitionsStagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `competitions_stages`.* FROM `competitions_stages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, competitionsStagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompetitionsStageSlice")
	}

	*o = slice

	return nil
}

// CompetitionsStageExists checks if the CompetitionsStage row exists.
func CompetitionsStageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `competitions_stages` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if competitions_stages exists")
	}

	return exists, nil
}
