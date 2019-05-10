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

// Fixture is an object representing the database table.
type Fixture struct {
	ID                 int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CompetitionID      int       `boil:"competition_id" json:"competition_id" toml:"competition_id" yaml:"competition_id"`
	CompetitionStageID int       `boil:"competition_stage_id" json:"competition_stage_id" toml:"competition_stage_id" yaml:"competition_stage_id"`
	TeamAID            int       `boil:"team_a_id" json:"team_a_id" toml:"team_a_id" yaml:"team_a_id"`
	TeamBID            int       `boil:"team_b_id" json:"team_b_id" toml:"team_b_id" yaml:"team_b_id"`
	TeamAScore         int       `boil:"team_a_score" json:"team_a_score" toml:"team_a_score" yaml:"team_a_score"`
	TeamBScore         int       `boil:"team_b_score" json:"team_b_score" toml:"team_b_score" yaml:"team_b_score"`
	StartingAt         null.Time `boil:"starting_at" json:"starting_at,omitempty" toml:"starting_at" yaml:"starting_at,omitempty"`
	EndingAt           null.Time `boil:"ending_at" json:"ending_at,omitempty" toml:"ending_at" yaml:"ending_at,omitempty"`
	Ended              null.Bool `boil:"ended" json:"ended,omitempty" toml:"ended" yaml:"ended,omitempty"`
	ExternalEntityID   null.Int  `boil:"external_entity_id" json:"external_entity_id,omitempty" toml:"external_entity_id" yaml:"external_entity_id,omitempty"`
	CreatedAt          null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt          null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *fixtureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fixtureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FixtureColumns = struct {
	ID                 string
	CompetitionID      string
	CompetitionStageID string
	TeamAID            string
	TeamBID            string
	TeamAScore         string
	TeamBScore         string
	StartingAt         string
	EndingAt           string
	Ended              string
	ExternalEntityID   string
	CreatedAt          string
	UpdatedAt          string
}{
	ID:                 "id",
	CompetitionID:      "competition_id",
	CompetitionStageID: "competition_stage_id",
	TeamAID:            "team_a_id",
	TeamBID:            "team_b_id",
	TeamAScore:         "team_a_score",
	TeamBScore:         "team_b_score",
	StartingAt:         "starting_at",
	EndingAt:           "ending_at",
	Ended:              "ended",
	ExternalEntityID:   "external_entity_id",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// Generated where

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var FixtureWhere = struct {
	ID                 whereHelperint
	CompetitionID      whereHelperint
	CompetitionStageID whereHelperint
	TeamAID            whereHelperint
	TeamBID            whereHelperint
	TeamAScore         whereHelperint
	TeamBScore         whereHelperint
	StartingAt         whereHelpernull_Time
	EndingAt           whereHelpernull_Time
	Ended              whereHelpernull_Bool
	ExternalEntityID   whereHelpernull_Int
	CreatedAt          whereHelpernull_Time
	UpdatedAt          whereHelpernull_Time
}{
	ID:                 whereHelperint{field: `id`},
	CompetitionID:      whereHelperint{field: `competition_id`},
	CompetitionStageID: whereHelperint{field: `competition_stage_id`},
	TeamAID:            whereHelperint{field: `team_a_id`},
	TeamBID:            whereHelperint{field: `team_b_id`},
	TeamAScore:         whereHelperint{field: `team_a_score`},
	TeamBScore:         whereHelperint{field: `team_b_score`},
	StartingAt:         whereHelpernull_Time{field: `starting_at`},
	EndingAt:           whereHelpernull_Time{field: `ending_at`},
	Ended:              whereHelpernull_Bool{field: `ended`},
	ExternalEntityID:   whereHelpernull_Int{field: `external_entity_id`},
	CreatedAt:          whereHelpernull_Time{field: `created_at`},
	UpdatedAt:          whereHelpernull_Time{field: `updated_at`},
}

// FixtureRels is where relationship names are stored.
var FixtureRels = struct {
}{}

// fixtureR is where relationships are stored.
type fixtureR struct {
}

// NewStruct creates a new relationship struct
func (*fixtureR) NewStruct() *fixtureR {
	return &fixtureR{}
}

// fixtureL is where Load methods for each relationship are stored.
type fixtureL struct{}

var (
	fixtureColumns               = []string{"id", "competition_id", "competition_stage_id", "team_a_id", "team_b_id", "team_a_score", "team_b_score", "starting_at", "ending_at", "ended", "external_entity_id", "created_at", "updated_at"}
	fixtureColumnsWithoutDefault = []string{"competition_id", "competition_stage_id", "team_a_id", "team_b_id", "team_a_score", "team_b_score", "starting_at", "ending_at", "ended", "external_entity_id", "created_at", "updated_at"}
	fixtureColumnsWithDefault    = []string{"id"}
	fixturePrimaryKeyColumns     = []string{"id"}
)

type (
	// FixtureSlice is an alias for a slice of pointers to Fixture.
	// This should generally be used opposed to []Fixture.
	FixtureSlice []*Fixture
	// FixtureHook is the signature for custom Fixture hook methods
	FixtureHook func(context.Context, boil.ContextExecutor, *Fixture) error

	fixtureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fixtureType                 = reflect.TypeOf(&Fixture{})
	fixtureMapping              = queries.MakeStructMapping(fixtureType)
	fixturePrimaryKeyMapping, _ = queries.BindMapping(fixtureType, fixtureMapping, fixturePrimaryKeyColumns)
	fixtureInsertCacheMut       sync.RWMutex
	fixtureInsertCache          = make(map[string]insertCache)
	fixtureUpdateCacheMut       sync.RWMutex
	fixtureUpdateCache          = make(map[string]updateCache)
	fixtureUpsertCacheMut       sync.RWMutex
	fixtureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fixtureBeforeInsertHooks []FixtureHook
var fixtureBeforeUpdateHooks []FixtureHook
var fixtureBeforeDeleteHooks []FixtureHook
var fixtureBeforeUpsertHooks []FixtureHook

var fixtureAfterInsertHooks []FixtureHook
var fixtureAfterSelectHooks []FixtureHook
var fixtureAfterUpdateHooks []FixtureHook
var fixtureAfterDeleteHooks []FixtureHook
var fixtureAfterUpsertHooks []FixtureHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Fixture) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Fixture) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Fixture) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Fixture) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Fixture) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Fixture) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Fixture) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Fixture) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Fixture) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fixtureAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFixtureHook registers your hook function for all future operations.
func AddFixtureHook(hookPoint boil.HookPoint, fixtureHook FixtureHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		fixtureBeforeInsertHooks = append(fixtureBeforeInsertHooks, fixtureHook)
	case boil.BeforeUpdateHook:
		fixtureBeforeUpdateHooks = append(fixtureBeforeUpdateHooks, fixtureHook)
	case boil.BeforeDeleteHook:
		fixtureBeforeDeleteHooks = append(fixtureBeforeDeleteHooks, fixtureHook)
	case boil.BeforeUpsertHook:
		fixtureBeforeUpsertHooks = append(fixtureBeforeUpsertHooks, fixtureHook)
	case boil.AfterInsertHook:
		fixtureAfterInsertHooks = append(fixtureAfterInsertHooks, fixtureHook)
	case boil.AfterSelectHook:
		fixtureAfterSelectHooks = append(fixtureAfterSelectHooks, fixtureHook)
	case boil.AfterUpdateHook:
		fixtureAfterUpdateHooks = append(fixtureAfterUpdateHooks, fixtureHook)
	case boil.AfterDeleteHook:
		fixtureAfterDeleteHooks = append(fixtureAfterDeleteHooks, fixtureHook)
	case boil.AfterUpsertHook:
		fixtureAfterUpsertHooks = append(fixtureAfterUpsertHooks, fixtureHook)
	}
}

// One returns a single fixture record from the query.
func (q fixtureQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Fixture, error) {
	o := &Fixture{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for fixtures")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Fixture records from the query.
func (q fixtureQuery) All(ctx context.Context, exec boil.ContextExecutor) (FixtureSlice, error) {
	var o []*Fixture

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Fixture slice")
	}

	if len(fixtureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Fixture records in the query.
func (q fixtureQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count fixtures rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fixtureQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if fixtures exists")
	}

	return count > 0, nil
}

// Fixtures retrieves all the records using an executor.
func Fixtures(mods ...qm.QueryMod) fixtureQuery {
	mods = append(mods, qm.From("`fixtures`"))
	return fixtureQuery{NewQuery(mods...)}
}

// FindFixture retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFixture(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Fixture, error) {
	fixtureObj := &Fixture{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `fixtures` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, fixtureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from fixtures")
	}

	return fixtureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Fixture) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no fixtures provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(fixtureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fixtureInsertCacheMut.RLock()
	cache, cached := fixtureInsertCache[key]
	fixtureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fixtureColumns,
			fixtureColumnsWithDefault,
			fixtureColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fixtureType, fixtureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fixtureType, fixtureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `fixtures` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `fixtures` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `fixtures` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, fixturePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into fixtures")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == fixtureMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for fixtures")
	}

CacheNoHooks:
	if !cached {
		fixtureInsertCacheMut.Lock()
		fixtureInsertCache[key] = cache
		fixtureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Fixture.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Fixture) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fixtureUpdateCacheMut.RLock()
	cache, cached := fixtureUpdateCache[key]
	fixtureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fixtureColumns,
			fixturePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update fixtures, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `fixtures` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, fixturePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fixtureType, fixtureMapping, append(wl, fixturePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update fixtures row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for fixtures")
	}

	if !cached {
		fixtureUpdateCacheMut.Lock()
		fixtureUpdateCache[key] = cache
		fixtureUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fixtureQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for fixtures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for fixtures")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FixtureSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fixturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `fixtures` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fixturePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in fixture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all fixture")
	}
	return rowsAff, nil
}

var mySQLFixtureUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Fixture) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no fixtures provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(fixtureColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFixtureUniqueColumns, o)

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

	fixtureUpsertCacheMut.RLock()
	cache, cached := fixtureUpsertCache[key]
	fixtureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fixtureColumns,
			fixtureColumnsWithDefault,
			fixtureColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			fixtureColumns,
			fixturePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert fixtures, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "fixtures", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `fixtures` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(fixtureType, fixtureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fixtureType, fixtureMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for fixtures")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == fixtureMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(fixtureType, fixtureMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for fixtures")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for fixtures")
	}

CacheNoHooks:
	if !cached {
		fixtureUpsertCacheMut.Lock()
		fixtureUpsertCache[key] = cache
		fixtureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Fixture record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Fixture) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Fixture provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), fixturePrimaryKeyMapping)
	sql := "DELETE FROM `fixtures` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from fixtures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for fixtures")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fixtureQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fixtureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from fixtures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for fixtures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FixtureSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Fixture slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(fixtureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fixturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `fixtures` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fixturePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from fixture slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for fixtures")
	}

	if len(fixtureAfterDeleteHooks) != 0 {
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
func (o *Fixture) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFixture(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FixtureSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FixtureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fixturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `fixtures`.* FROM `fixtures` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fixturePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FixtureSlice")
	}

	*o = slice

	return nil
}

// FixtureExists checks if the Fixture row exists.
func FixtureExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `fixtures` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if fixtures exists")
	}

	return exists, nil
}