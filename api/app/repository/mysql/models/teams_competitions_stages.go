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

// TeamsCompetitionsStage is an object representing the database table.
type TeamsCompetitionsStage struct {
	ID                 int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TeamCompetitionID  int       `boil:"team_competition_id" json:"team_competition_id" toml:"team_competition_id" yaml:"team_competition_id"`
	CompetitionStageID int       `boil:"competition_stage_id" json:"competition_stage_id" toml:"competition_stage_id" yaml:"competition_stage_id"`
	Position           int       `boil:"position" json:"position" toml:"position" yaml:"position"`
	Points             int       `boil:"points" json:"points" toml:"points" yaml:"points"`
	GoalsFor           int       `boil:"goals_for" json:"goals_for" toml:"goals_for" yaml:"goals_for"`
	GoalsAgainst       int       `boil:"goals_against" json:"goals_against" toml:"goals_against" yaml:"goals_against"`
	GamesPlayed        int       `boil:"games_played" json:"games_played" toml:"games_played" yaml:"games_played"`
	CreatedAt          null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt          null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *teamsCompetitionsStageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L teamsCompetitionsStageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TeamsCompetitionsStageColumns = struct {
	ID                 string
	TeamCompetitionID  string
	CompetitionStageID string
	Position           string
	Points             string
	GoalsFor           string
	GoalsAgainst       string
	GamesPlayed        string
	CreatedAt          string
	UpdatedAt          string
}{
	ID:                 "id",
	TeamCompetitionID:  "team_competition_id",
	CompetitionStageID: "competition_stage_id",
	Position:           "position",
	Points:             "points",
	GoalsFor:           "goals_for",
	GoalsAgainst:       "goals_against",
	GamesPlayed:        "games_played",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// Generated where

var TeamsCompetitionsStageWhere = struct {
	ID                 whereHelperint
	TeamCompetitionID  whereHelperint
	CompetitionStageID whereHelperint
	Position           whereHelperint
	Points             whereHelperint
	GoalsFor           whereHelperint
	GoalsAgainst       whereHelperint
	GamesPlayed        whereHelperint
	CreatedAt          whereHelpernull_Time
	UpdatedAt          whereHelpernull_Time
}{
	ID:                 whereHelperint{field: `id`},
	TeamCompetitionID:  whereHelperint{field: `team_competition_id`},
	CompetitionStageID: whereHelperint{field: `competition_stage_id`},
	Position:           whereHelperint{field: `position`},
	Points:             whereHelperint{field: `points`},
	GoalsFor:           whereHelperint{field: `goals_for`},
	GoalsAgainst:       whereHelperint{field: `goals_against`},
	GamesPlayed:        whereHelperint{field: `games_played`},
	CreatedAt:          whereHelpernull_Time{field: `created_at`},
	UpdatedAt:          whereHelpernull_Time{field: `updated_at`},
}

// TeamsCompetitionsStageRels is where relationship names are stored.
var TeamsCompetitionsStageRels = struct {
}{}

// teamsCompetitionsStageR is where relationships are stored.
type teamsCompetitionsStageR struct {
}

// NewStruct creates a new relationship struct
func (*teamsCompetitionsStageR) NewStruct() *teamsCompetitionsStageR {
	return &teamsCompetitionsStageR{}
}

// teamsCompetitionsStageL is where Load methods for each relationship are stored.
type teamsCompetitionsStageL struct{}

var (
	teamsCompetitionsStageColumns               = []string{"id", "team_competition_id", "competition_stage_id", "position", "points", "goals_for", "goals_against", "games_played", "created_at", "updated_at"}
	teamsCompetitionsStageColumnsWithoutDefault = []string{"team_competition_id", "competition_stage_id", "position", "points", "goals_for", "goals_against", "games_played", "created_at", "updated_at"}
	teamsCompetitionsStageColumnsWithDefault    = []string{"id"}
	teamsCompetitionsStagePrimaryKeyColumns     = []string{"id"}
)

type (
	// TeamsCompetitionsStageSlice is an alias for a slice of pointers to TeamsCompetitionsStage.
	// This should generally be used opposed to []TeamsCompetitionsStage.
	TeamsCompetitionsStageSlice []*TeamsCompetitionsStage
	// TeamsCompetitionsStageHook is the signature for custom TeamsCompetitionsStage hook methods
	TeamsCompetitionsStageHook func(context.Context, boil.ContextExecutor, *TeamsCompetitionsStage) error

	teamsCompetitionsStageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	teamsCompetitionsStageType                 = reflect.TypeOf(&TeamsCompetitionsStage{})
	teamsCompetitionsStageMapping              = queries.MakeStructMapping(teamsCompetitionsStageType)
	teamsCompetitionsStagePrimaryKeyMapping, _ = queries.BindMapping(teamsCompetitionsStageType, teamsCompetitionsStageMapping, teamsCompetitionsStagePrimaryKeyColumns)
	teamsCompetitionsStageInsertCacheMut       sync.RWMutex
	teamsCompetitionsStageInsertCache          = make(map[string]insertCache)
	teamsCompetitionsStageUpdateCacheMut       sync.RWMutex
	teamsCompetitionsStageUpdateCache          = make(map[string]updateCache)
	teamsCompetitionsStageUpsertCacheMut       sync.RWMutex
	teamsCompetitionsStageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var teamsCompetitionsStageBeforeInsertHooks []TeamsCompetitionsStageHook
var teamsCompetitionsStageBeforeUpdateHooks []TeamsCompetitionsStageHook
var teamsCompetitionsStageBeforeDeleteHooks []TeamsCompetitionsStageHook
var teamsCompetitionsStageBeforeUpsertHooks []TeamsCompetitionsStageHook

var teamsCompetitionsStageAfterInsertHooks []TeamsCompetitionsStageHook
var teamsCompetitionsStageAfterSelectHooks []TeamsCompetitionsStageHook
var teamsCompetitionsStageAfterUpdateHooks []TeamsCompetitionsStageHook
var teamsCompetitionsStageAfterDeleteHooks []TeamsCompetitionsStageHook
var teamsCompetitionsStageAfterUpsertHooks []TeamsCompetitionsStageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TeamsCompetitionsStage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TeamsCompetitionsStage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TeamsCompetitionsStage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TeamsCompetitionsStage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TeamsCompetitionsStage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TeamsCompetitionsStage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TeamsCompetitionsStage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TeamsCompetitionsStage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TeamsCompetitionsStage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamsCompetitionsStageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTeamsCompetitionsStageHook registers your hook function for all future operations.
func AddTeamsCompetitionsStageHook(hookPoint boil.HookPoint, teamsCompetitionsStageHook TeamsCompetitionsStageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		teamsCompetitionsStageBeforeInsertHooks = append(teamsCompetitionsStageBeforeInsertHooks, teamsCompetitionsStageHook)
	case boil.BeforeUpdateHook:
		teamsCompetitionsStageBeforeUpdateHooks = append(teamsCompetitionsStageBeforeUpdateHooks, teamsCompetitionsStageHook)
	case boil.BeforeDeleteHook:
		teamsCompetitionsStageBeforeDeleteHooks = append(teamsCompetitionsStageBeforeDeleteHooks, teamsCompetitionsStageHook)
	case boil.BeforeUpsertHook:
		teamsCompetitionsStageBeforeUpsertHooks = append(teamsCompetitionsStageBeforeUpsertHooks, teamsCompetitionsStageHook)
	case boil.AfterInsertHook:
		teamsCompetitionsStageAfterInsertHooks = append(teamsCompetitionsStageAfterInsertHooks, teamsCompetitionsStageHook)
	case boil.AfterSelectHook:
		teamsCompetitionsStageAfterSelectHooks = append(teamsCompetitionsStageAfterSelectHooks, teamsCompetitionsStageHook)
	case boil.AfterUpdateHook:
		teamsCompetitionsStageAfterUpdateHooks = append(teamsCompetitionsStageAfterUpdateHooks, teamsCompetitionsStageHook)
	case boil.AfterDeleteHook:
		teamsCompetitionsStageAfterDeleteHooks = append(teamsCompetitionsStageAfterDeleteHooks, teamsCompetitionsStageHook)
	case boil.AfterUpsertHook:
		teamsCompetitionsStageAfterUpsertHooks = append(teamsCompetitionsStageAfterUpsertHooks, teamsCompetitionsStageHook)
	}
}

// One returns a single teamsCompetitionsStage record from the query.
func (q teamsCompetitionsStageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TeamsCompetitionsStage, error) {
	o := &TeamsCompetitionsStage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for teams_competitions_stages")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TeamsCompetitionsStage records from the query.
func (q teamsCompetitionsStageQuery) All(ctx context.Context, exec boil.ContextExecutor) (TeamsCompetitionsStageSlice, error) {
	var o []*TeamsCompetitionsStage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TeamsCompetitionsStage slice")
	}

	if len(teamsCompetitionsStageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TeamsCompetitionsStage records in the query.
func (q teamsCompetitionsStageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count teams_competitions_stages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q teamsCompetitionsStageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if teams_competitions_stages exists")
	}

	return count > 0, nil
}

// TeamsCompetitionsStages retrieves all the records using an executor.
func TeamsCompetitionsStages(mods ...qm.QueryMod) teamsCompetitionsStageQuery {
	mods = append(mods, qm.From("`teams_competitions_stages`"))
	return teamsCompetitionsStageQuery{NewQuery(mods...)}
}

// FindTeamsCompetitionsStage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTeamsCompetitionsStage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TeamsCompetitionsStage, error) {
	teamsCompetitionsStageObj := &TeamsCompetitionsStage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `teams_competitions_stages` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, teamsCompetitionsStageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from teams_competitions_stages")
	}

	return teamsCompetitionsStageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TeamsCompetitionsStage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no teams_competitions_stages provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(teamsCompetitionsStageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	teamsCompetitionsStageInsertCacheMut.RLock()
	cache, cached := teamsCompetitionsStageInsertCache[key]
	teamsCompetitionsStageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			teamsCompetitionsStageColumns,
			teamsCompetitionsStageColumnsWithDefault,
			teamsCompetitionsStageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(teamsCompetitionsStageType, teamsCompetitionsStageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(teamsCompetitionsStageType, teamsCompetitionsStageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `teams_competitions_stages` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `teams_competitions_stages` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `teams_competitions_stages` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, teamsCompetitionsStagePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into teams_competitions_stages")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == teamsCompetitionsStageMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for teams_competitions_stages")
	}

CacheNoHooks:
	if !cached {
		teamsCompetitionsStageInsertCacheMut.Lock()
		teamsCompetitionsStageInsertCache[key] = cache
		teamsCompetitionsStageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TeamsCompetitionsStage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TeamsCompetitionsStage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	teamsCompetitionsStageUpdateCacheMut.RLock()
	cache, cached := teamsCompetitionsStageUpdateCache[key]
	teamsCompetitionsStageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			teamsCompetitionsStageColumns,
			teamsCompetitionsStagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update teams_competitions_stages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `teams_competitions_stages` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, teamsCompetitionsStagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(teamsCompetitionsStageType, teamsCompetitionsStageMapping, append(wl, teamsCompetitionsStagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update teams_competitions_stages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for teams_competitions_stages")
	}

	if !cached {
		teamsCompetitionsStageUpdateCacheMut.Lock()
		teamsCompetitionsStageUpdateCache[key] = cache
		teamsCompetitionsStageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q teamsCompetitionsStageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for teams_competitions_stages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for teams_competitions_stages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TeamsCompetitionsStageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamsCompetitionsStagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `teams_competitions_stages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, teamsCompetitionsStagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in teamsCompetitionsStage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all teamsCompetitionsStage")
	}
	return rowsAff, nil
}

var mySQLTeamsCompetitionsStageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TeamsCompetitionsStage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no teams_competitions_stages provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(teamsCompetitionsStageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTeamsCompetitionsStageUniqueColumns, o)

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

	teamsCompetitionsStageUpsertCacheMut.RLock()
	cache, cached := teamsCompetitionsStageUpsertCache[key]
	teamsCompetitionsStageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			teamsCompetitionsStageColumns,
			teamsCompetitionsStageColumnsWithDefault,
			teamsCompetitionsStageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			teamsCompetitionsStageColumns,
			teamsCompetitionsStagePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert teams_competitions_stages, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "teams_competitions_stages", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `teams_competitions_stages` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(teamsCompetitionsStageType, teamsCompetitionsStageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(teamsCompetitionsStageType, teamsCompetitionsStageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for teams_competitions_stages")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == teamsCompetitionsStageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(teamsCompetitionsStageType, teamsCompetitionsStageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for teams_competitions_stages")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for teams_competitions_stages")
	}

CacheNoHooks:
	if !cached {
		teamsCompetitionsStageUpsertCacheMut.Lock()
		teamsCompetitionsStageUpsertCache[key] = cache
		teamsCompetitionsStageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TeamsCompetitionsStage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TeamsCompetitionsStage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TeamsCompetitionsStage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), teamsCompetitionsStagePrimaryKeyMapping)
	sql := "DELETE FROM `teams_competitions_stages` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from teams_competitions_stages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for teams_competitions_stages")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q teamsCompetitionsStageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no teamsCompetitionsStageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from teams_competitions_stages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for teams_competitions_stages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TeamsCompetitionsStageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TeamsCompetitionsStage slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(teamsCompetitionsStageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamsCompetitionsStagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `teams_competitions_stages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, teamsCompetitionsStagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from teamsCompetitionsStage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for teams_competitions_stages")
	}

	if len(teamsCompetitionsStageAfterDeleteHooks) != 0 {
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
func (o *TeamsCompetitionsStage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTeamsCompetitionsStage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TeamsCompetitionsStageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TeamsCompetitionsStageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamsCompetitionsStagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `teams_competitions_stages`.* FROM `teams_competitions_stages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, teamsCompetitionsStagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TeamsCompetitionsStageSlice")
	}

	*o = slice

	return nil
}

// TeamsCompetitionsStageExists checks if the TeamsCompetitionsStage row exists.
func TeamsCompetitionsStageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `teams_competitions_stages` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if teams_competitions_stages exists")
	}

	return exists, nil
}