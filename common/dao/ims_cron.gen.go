// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/donknap/dpanel/common/entity"
)

func newCron(db *gorm.DB, opts ...gen.DOOption) cron {
	_cron := cron{}

	_cron.cronDo.UseDB(db, opts...)
	_cron.cronDo.UseModel(&entity.Cron{})

	tableName := _cron.cronDo.TableName()
	_cron.ALL = field.NewAsterisk(tableName)
	_cron.ID = field.NewInt32(tableName, "id")
	_cron.Title = field.NewString(tableName, "title")
	_cron.Setting = field.NewField(tableName, "setting")

	_cron.fillFieldMap()

	return _cron
}

type cron struct {
	cronDo

	ALL     field.Asterisk
	ID      field.Int32
	Title   field.String
	Setting field.Field

	fieldMap map[string]field.Expr
}

func (c cron) Table(newTableName string) *cron {
	c.cronDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cron) As(alias string) *cron {
	c.cronDo.DO = *(c.cronDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cron) updateTableName(table string) *cron {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.Title = field.NewString(table, "title")
	c.Setting = field.NewField(table, "setting")

	c.fillFieldMap()

	return c
}

func (c *cron) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cron) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["id"] = c.ID
	c.fieldMap["title"] = c.Title
	c.fieldMap["setting"] = c.Setting
}

func (c cron) clone(db *gorm.DB) cron {
	c.cronDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cron) replaceDB(db *gorm.DB) cron {
	c.cronDo.ReplaceDB(db)
	return c
}

type cronDo struct{ gen.DO }

type ICronDo interface {
	gen.SubQuery
	Debug() ICronDo
	WithContext(ctx context.Context) ICronDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICronDo
	WriteDB() ICronDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICronDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICronDo
	Not(conds ...gen.Condition) ICronDo
	Or(conds ...gen.Condition) ICronDo
	Select(conds ...field.Expr) ICronDo
	Where(conds ...gen.Condition) ICronDo
	Order(conds ...field.Expr) ICronDo
	Distinct(cols ...field.Expr) ICronDo
	Omit(cols ...field.Expr) ICronDo
	Join(table schema.Tabler, on ...field.Expr) ICronDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICronDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICronDo
	Group(cols ...field.Expr) ICronDo
	Having(conds ...gen.Condition) ICronDo
	Limit(limit int) ICronDo
	Offset(offset int) ICronDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICronDo
	Unscoped() ICronDo
	Create(values ...*entity.Cron) error
	CreateInBatches(values []*entity.Cron, batchSize int) error
	Save(values ...*entity.Cron) error
	First() (*entity.Cron, error)
	Take() (*entity.Cron, error)
	Last() (*entity.Cron, error)
	Find() ([]*entity.Cron, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Cron, err error)
	FindInBatches(result *[]*entity.Cron, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Cron) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICronDo
	Assign(attrs ...field.AssignExpr) ICronDo
	Joins(fields ...field.RelationField) ICronDo
	Preload(fields ...field.RelationField) ICronDo
	FirstOrInit() (*entity.Cron, error)
	FirstOrCreate() (*entity.Cron, error)
	FindByPage(offset int, limit int) (result []*entity.Cron, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICronDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cronDo) Debug() ICronDo {
	return c.withDO(c.DO.Debug())
}

func (c cronDo) WithContext(ctx context.Context) ICronDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cronDo) ReadDB() ICronDo {
	return c.Clauses(dbresolver.Read)
}

func (c cronDo) WriteDB() ICronDo {
	return c.Clauses(dbresolver.Write)
}

func (c cronDo) Session(config *gorm.Session) ICronDo {
	return c.withDO(c.DO.Session(config))
}

func (c cronDo) Clauses(conds ...clause.Expression) ICronDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cronDo) Returning(value interface{}, columns ...string) ICronDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cronDo) Not(conds ...gen.Condition) ICronDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cronDo) Or(conds ...gen.Condition) ICronDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cronDo) Select(conds ...field.Expr) ICronDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cronDo) Where(conds ...gen.Condition) ICronDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cronDo) Order(conds ...field.Expr) ICronDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cronDo) Distinct(cols ...field.Expr) ICronDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cronDo) Omit(cols ...field.Expr) ICronDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cronDo) Join(table schema.Tabler, on ...field.Expr) ICronDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cronDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICronDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cronDo) RightJoin(table schema.Tabler, on ...field.Expr) ICronDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cronDo) Group(cols ...field.Expr) ICronDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cronDo) Having(conds ...gen.Condition) ICronDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cronDo) Limit(limit int) ICronDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cronDo) Offset(offset int) ICronDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cronDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICronDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cronDo) Unscoped() ICronDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cronDo) Create(values ...*entity.Cron) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cronDo) CreateInBatches(values []*entity.Cron, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cronDo) Save(values ...*entity.Cron) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cronDo) First() (*entity.Cron, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Cron), nil
	}
}

func (c cronDo) Take() (*entity.Cron, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Cron), nil
	}
}

func (c cronDo) Last() (*entity.Cron, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Cron), nil
	}
}

func (c cronDo) Find() ([]*entity.Cron, error) {
	result, err := c.DO.Find()
	return result.([]*entity.Cron), err
}

func (c cronDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Cron, err error) {
	buf := make([]*entity.Cron, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cronDo) FindInBatches(result *[]*entity.Cron, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cronDo) Attrs(attrs ...field.AssignExpr) ICronDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cronDo) Assign(attrs ...field.AssignExpr) ICronDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cronDo) Joins(fields ...field.RelationField) ICronDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cronDo) Preload(fields ...field.RelationField) ICronDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cronDo) FirstOrInit() (*entity.Cron, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Cron), nil
	}
}

func (c cronDo) FirstOrCreate() (*entity.Cron, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Cron), nil
	}
}

func (c cronDo) FindByPage(offset int, limit int) (result []*entity.Cron, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cronDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cronDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cronDo) Delete(models ...*entity.Cron) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cronDo) withDO(do gen.Dao) *cronDo {
	c.DO = *do.(*gen.DO)
	return c
}
